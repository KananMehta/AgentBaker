package scenario

import (
	"github.com/Azure/agentbaker/pkg/agent/datamodel"
	"github.com/Azure/agentbakere2e/cluster"
	"github.com/Azure/agentbakere2e/config"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// Returns config for the 'gpu' E2E scenario
func azurelinuxv2gpu() *Scenario {
	return &Scenario{
		Name:        "azurelinuxv2-gpu",
		Description: "Tests that a GPU-enabled node using a AzureLinuxV2 (CgroupV2) VHD can be properly bootstrapped",
		Tags: Tags{
			Name:     "azurelinuxv2-gpu",
			OS:       "azurelinuxv2",
			Platform: "x64",
			GPU:      true,
		},
		Config: Config{
			Cluster:     cluster.ClusterKubenet,
			VHDSelector: config.VHDAzureLinuxV2Gen2,
			BootstrapConfigMutator: func(nbc *datamodel.NodeBootstrappingConfiguration) {
				nbc.ContainerService.Properties.AgentPoolProfiles[0].VMSize = "Standard_NC6s_v3"
				nbc.ContainerService.Properties.AgentPoolProfiles[0].Distro = "aks-azurelinux-v2-gen2"
				nbc.AgentPoolProfile.VMSize = "Standard_NC6s_v3"
				nbc.AgentPoolProfile.Distro = "aks-azurelinux-v2-gen2"
				nbc.ConfigGPUDriverIfNeeded = true
				nbc.EnableGPUDevicePluginIfNeeded = false
				nbc.EnableNvidia = true
			},
			VMConfigMutator: func(vmss *armcompute.VirtualMachineScaleSet) {
				vmss.SKU.Name = to.Ptr("Standard_NC6s_v3")
			},
		},
	}
}
