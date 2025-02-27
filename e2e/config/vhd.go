package config

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

const (
	imageGallery       = "/subscriptions/8ecadfc9-d1a3-4ea4-b844-0d9f87e4d7c8/resourceGroups/aksvhdtestbuildrg/providers/Microsoft.Compute/galleries/PackerSigGalleryEastUS/images/"
	noSelectionTagName = "abe2e-ignore"

	fetchResourceIDTimeout = 5 * time.Minute
)

var (
	VHDUbuntu1804Gen2Containerd      = newSIGImageVersionResourceIDFetcher(imageGallery + "1804gen2containerd")
	VHDUbuntu2204Gen2Arm64Containerd = newSIGImageVersionResourceIDFetcher(imageGallery + "2204gen2arm64containerd")
	VHDUbuntu2204Gen2Containerd      = newSIGImageVersionResourceIDFetcher(imageGallery + "2204gen2containerd")
	VHDAzureLinuxV2Gen2Arm64         = newSIGImageVersionResourceIDFetcher(imageGallery + "AzureLinuxV2gen2arm64")
	VHDAzureLinuxV2Gen2              = newSIGImageVersionResourceIDFetcher(imageGallery + "AzureLinuxV2gen2")
	VHDCBLMarinerV2Gen2Arm64         = newSIGImageVersionResourceIDFetcher(imageGallery + "CBLMarinerV2gen2arm64")
	VHDCBLMarinerV2Gen2              = newSIGImageVersionResourceIDFetcher(imageGallery + "CBLMarinerV2gen2")

	// this is a particular 2204gen2containerd image originally built with private packages,
	// if we ever want to update this then we'd need to run a new VHD build using private package overrides
	VHDUbuntu2204Gen2ContainerdPrivateKubePkg = newStaticSIGImageVersionResourceIDFetcher(imageGallery + "2204Gen2/versions/1.1704411049.2812")
)

var ErrNotFound = fmt.Errorf("not found")

type sigImageDefinition struct {
	subscriptionID string
	resourceGroup  string
	gallery        string
	definition     string
}

type sigImageVersion struct {
	sigImageDefinition
	version string
}

func newSIGImageDefinitionFromResourceID(resourceID *arm.ResourceID) sigImageDefinition {
	return sigImageDefinition{
		subscriptionID: resourceID.SubscriptionID,
		resourceGroup:  resourceID.ResourceGroupName,
		gallery:        resourceID.Parent.Name,
		definition:     resourceID.Name,
	}
}

func newSIGImageVersionFromResourceID(resourceID *arm.ResourceID) sigImageVersion {
	return sigImageVersion{
		sigImageDefinition: newSIGImageDefinitionFromResourceID(resourceID.Parent),
		version:            resourceID.Name,
	}
}

// VHDResourceID represents a resource ID pointing to a VHD in Azure. This could be theoretically
// be the resource ID of a managed image or SIG image version, though for now this will always be a SIG image version.
type VHDResourceID string

func (id VHDResourceID) Short() string {
	sep := "Microsoft.Compute/galleries/"
	str := string(id)
	if strings.Contains(str, sep) && !strings.HasSuffix(str, sep) {
		return strings.Split(str, sep)[1]
	}
	return str
}

// newSIGImageVersionResourceIDFetcher is a factory function
// it returns a function that fetches the latest VHDResourceID for a given image
// the function is memoized and will only evaluate once on the first call
func newSIGImageVersionResourceIDFetcher(imageDefinitionResourceID string) func() (VHDResourceID, error) {
	resourceID := VHDResourceID("")
	var err error
	once := sync.Once{}
	// evaluate the function once and cache the result
	return func() (VHDResourceID, error) {
		once.Do(func() {
			resourceID, err = findLatestSIGImageVersionWithTag(imageDefinitionResourceID, SIGVersionTagName, SIGVersionTagValue)
			if err != nil {
				err = fmt.Errorf("img: %s, tag %s=%s, err %w", imageDefinitionResourceID, SIGVersionTagName, SIGVersionTagValue, err)
				log.Printf("failed to find the latest image %s", err)
			} else {
				log.Printf("Resource ID for %s: %s", imageDefinitionResourceID, resourceID)
			}
		})
		return resourceID, err
	}
}

func newStaticSIGImageVersionResourceIDFetcher(imageVersionResourceID string) func() (VHDResourceID, error) {
	resourceID := VHDResourceID("")
	var err error
	once := sync.Once{}

	return func() (VHDResourceID, error) {
		once.Do(func() {
			resourceID, err = ensureStaticSIGImageVersion(imageVersionResourceID)
			if err != nil {
				err = fmt.Errorf("img: %s, err: %w", imageVersionResourceID, err)
				log.Printf("failed to find static image %s", err)
			} else {
				log.Printf("Resource ID for %s: %s", imageVersionResourceID, resourceID)
			}
		})
		return resourceID, err
	}
}

func ensureStaticSIGImageVersion(imageVersionResourceID string) (VHDResourceID, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), fetchResourceIDTimeout)
	defer cancel()

	rid, err := arm.ParseResourceID(imageVersionResourceID)
	if err != nil {
		return "", fmt.Errorf("parsing image version resouce ID: %w", err)
	}
	version := newSIGImageVersionFromResourceID(rid)

	resp, err := Azure.GalleryImageVersionClient.Get(ctx, version.resourceGroup, version.gallery, version.definition, version.version, nil)
	if err != nil {
		return "", fmt.Errorf("getting live image version info: %w", err)
	}

	liveVersion := &resp.GalleryImageVersion
	if err := ensureProvisioningState(liveVersion); err != nil {
		return "", fmt.Errorf("ensuring image version provisioning state: %w", err)
	}

	if err := ensureReplication(ctx, version.sigImageDefinition, liveVersion); err != nil {
		return "", fmt.Errorf("ensuring image replication: %w", err)
	}

	return VHDResourceID(imageVersionResourceID), nil
}

func findLatestSIGImageVersionWithTag(imageDefinitionResourceID, tagName, tagValue string) (VHDResourceID, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), fetchResourceIDTimeout)
	defer cancel()

	rid, err := arm.ParseResourceID(imageDefinitionResourceID)
	if err != nil {
		return "", fmt.Errorf("parsing image definition resource ID: %w", err)
	}
	definition := newSIGImageDefinitionFromResourceID(rid)

	pager := Azure.GalleryImageVersionClient.NewListByGalleryImagePager(definition.resourceGroup, definition.gallery, definition.definition, nil)
	var latestVersion *armcompute.GalleryImageVersion
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return "", fmt.Errorf("failed to get next page: %w", err)
		}
		versions := page.Value
		for _, version := range versions {
			// skip images tagged with the no-selection tag, indicating they
			// shouldn't be selected dynmically for running abe2e scenarios
			if _, ok := version.Tags[noSelectionTagName]; ok {
				continue
			}
			tag, ok := version.Tags[tagName]
			if !ok || tag == nil || *tag != tagValue {
				continue
			}
			if err := ensureProvisioningState(version); err != nil {
				log.Printf("ensuring image version %s provisioning state: %s, will not consider for selection", *version.ID, err)
				continue
			}
			if latestVersion == nil || version.Properties.PublishingProfile.PublishedDate.After(*latestVersion.Properties.PublishingProfile.PublishedDate) {
				latestVersion = version
			}
		}
	}
	if latestVersion == nil {
		return "", ErrNotFound
	}

	if err := ensureReplication(ctx, definition, latestVersion); err != nil {
		return "", fmt.Errorf("ensuring image replication: %w", err)
	}

	return VHDResourceID(*latestVersion.ID), nil
}

func ensureReplication(ctx context.Context, definition sigImageDefinition, version *armcompute.GalleryImageVersion) error {
	if replicatedToCurrentRegion(version) {
		log.Printf("image version %s is already replicated to region %s", *version.ID, Location)
		return nil
	}
	return replicateToCurrentRegion(ctx, definition, version)
}

func replicatedToCurrentRegion(version *armcompute.GalleryImageVersion) bool {
	for _, targetRegion := range version.Properties.PublishingProfile.TargetRegions {
		if strings.EqualFold(strings.ReplaceAll(*targetRegion.Name, " ", ""), Location) {
			return true
		}
	}
	return false
}

func replicateToCurrentRegion(ctx context.Context, definition sigImageDefinition, version *armcompute.GalleryImageVersion) error {
	log.Printf("will replicate image version %s to region %s...", *version.ID, Location)

	version.Properties.PublishingProfile.TargetRegions = append(version.Properties.PublishingProfile.TargetRegions, &armcompute.TargetRegion{
		Name:                 &Location,
		RegionalReplicaCount: to.Ptr[int32](1),
		StorageAccountType:   to.Ptr(armcompute.StorageAccountTypeStandardLRS),
	})

	resp, err := Azure.GalleryImageVersionClient.BeginCreateOrUpdate(ctx, definition.resourceGroup, definition.gallery, definition.definition, *version.Name, *version, nil)
	if err != nil {
		return fmt.Errorf("begin updating image version target regions: %w", err)
	}
	if _, err := resp.PollUntilDone(ctx, nil); err != nil {
		return fmt.Errorf("updating image version target regions: %w", err)
	}

	return nil
}

func ensureProvisioningState(version *armcompute.GalleryImageVersion) error {
	if *version.Properties.ProvisioningState != armcompute.GalleryImageVersionPropertiesProvisioningStateSucceeded {
		return fmt.Errorf("unexpected provisioning state: %q", *version.Properties.ProvisioningState)
	}
	return nil
}
