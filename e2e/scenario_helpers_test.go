package e2e

import (
	"context"
	"errors"
	"path/filepath"
	"sync"
	"testing"

	"github.com/Azure/agentbaker/pkg/agent/datamodel"
	"github.com/Azure/agentbakere2e/config"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var once sync.Once

func RunScenario(t *testing.T, s *Scenario) {
	once.Do(func() {
		if err := createE2ELoggingDir(); err != nil {
			panic(err)
		}

		if err := ensureResourceGroup(context.Background(), t); err != nil {
			panic(err)
		}
	})
	t.Parallel()
	ctx := context.Background()
	model, err := s.Cluster(ctx, t)
	require.NoError(t, err)
	maybeSkipScenario(t, s)
	loggingDir, err := createVMLogsDir(t.Name())
	require.NoError(t, err)
	nbc, err := s.PrepareNodeBootstrappingConfiguration(model.NodeBootstrappingConfiguration)
	require.NoError(t, err)

	executeScenario(ctx, t, &scenarioRunOpts{
		clusterConfig: model,
		scenario:      s,
		nbc:           nbc,
		loggingDir:    loggingDir,
	})
}

func maybeSkipScenario(t *testing.T, s *Scenario) {
	s.Tags.OS = s.VHD.OS
	s.Tags.Arch = s.VHD.Arch
	s.Tags.ImageName = s.VHD.Name
	t.Logf("running scenario %q with tags %+v", t.Name(), s.Tags)
	if config.TagsToRun != "" {
		matches, err := s.Tags.MatchesFilters(config.TagsToRun)
		if err != nil {
			t.Fatalf("could not match tags for %q: %s", t.Name(), err)
		}
		if !matches {
			t.Skipf("skipping scenario %q: scenario tags %+v does not match filter %q", t.Name(), s.Tags, config.TagsToRun)
		}
	}

	if config.TagsToSkip != "" {
		matches, err := s.Tags.MatchesAnyFilter(config.TagsToSkip)
		if err != nil {
			t.Fatalf("could not match tags for %q: %s", t.Name(), err)
		}
		if matches {
			t.Skipf("skipping scenario %q: scenario tags %+v matches filter %q", t.Name(), s.Tags, config.TagsToSkip)
		}
	}

	_, err := s.VHD.VHDResourceID(t)
	if err != nil {
		if config.IgnoreScenariosWithMissingVHD && errors.Is(err, config.ErrNotFound) {
			t.Skipf("skipping scenario %q: could not find image", t.Name())
		} else {
			t.Fatalf("could not find image for %q: %s", t.Name(), err)
		}
	}
}

func executeScenario(ctx context.Context, t *testing.T, opts *scenarioRunOpts) {
	rid, _ := opts.scenario.VHD.VHDResourceID(t)
	t.Logf("running scenario %q with image %q in aks cluster %q", t.Name(), rid, *opts.clusterConfig.Model.ID)

	privateKeyBytes, publicKeyBytes, err := getNewRSAKeyPair()
	assert.NoError(t, err)

	vmssName := getVmssName(t)

	vmssSucceeded := true
	vmssModel, err := bootstrapVMSS(ctx, t, vmssName, opts, publicKeyBytes)
	if err != nil {
		vmssSucceeded = false
		if config.SkipTestsWithSKUCapacityIssue {
			var respErr *azcore.ResponseError
			if errors.As(err, &respErr) && respErr.StatusCode == 409 && respErr.ErrorCode == "SkuNotAvailable" {
				t.Skip("skipping scenario SKU not available", t.Name(), err)
			}
		}

		if !isVMExtensionProvisioningError(err) {
			t.Fatalf("creating VMSS %s: %v", vmssName, err)
		}
		t.Logf("vm %s was unable to be provisioned due to a CSE error, will still attempt to extract provisioning logs...\n", vmssName)
		t.Fail()
	}

	if config.KeepVMSS {
		defer func() {
			t.Logf("vmss %q will be retained for debugging purposes, please make sure to manually delete it later", vmssName)
			if vmssModel != nil {
				t.Logf("retained vmss %s resource ID: %q", vmssName, *vmssModel.ID)
			} else {
				t.Logf("WARNING: model of retained vmss %q is nil", vmssName)
			}
			if err := writeToFile(filepath.Join(opts.loggingDir, "sshkey"), string(privateKeyBytes)); err != nil {
				t.Fatalf("failed to write retained vmss %s private ssh key to disk: %s", vmssName, err)
			}
		}()
	} else {
		if vmssModel != nil {
			if err := writeToFile(filepath.Join(opts.loggingDir, "vmssId.txt"), *vmssModel.ID); err != nil {
				t.Fatalf("failed to write vmss %s resource ID to disk: %s", vmssName, err)
			}
		} else {
			t.Logf("WARNING: bootstrapped vmss model was nil for %s", vmssName)
		}
	}

	vmPrivateIP, err := pollGetVMPrivateIP(ctx, t, vmssName, opts)
	require.NoError(t, err)

	// Perform posthoc log extraction when the VMSS creation succeeded or failed due to a CSE error
	defer func() {
		err := pollExtractVMLogs(ctx, t, vmssName, vmPrivateIP, privateKeyBytes, opts)
		require.NoError(t, err)
	}()

	// Only perform node readiness/pod-related checks when VMSS creation succeeded
	if vmssSucceeded {
		t.Logf("vmss %s creation succeeded, proceeding with node readiness and pod checks...", vmssName)
		nodeName, err := validateNodeHealth(ctx, opts.clusterConfig.Kube, vmssName)
		require.NoError(t, err)

		if opts.nbc.AgentPoolProfile.WorkloadRuntime == datamodel.WasmWasi {
			t.Logf("wasm scenario: running wasm validation on %s...", vmssName)
			err = ensureWasmRuntimeClasses(ctx, opts.clusterConfig.Kube)
			require.NoError(t, err)
			err = validateWasm(ctx, t, opts.clusterConfig.Kube, nodeName, string(privateKeyBytes))
			require.NoError(t, err)
		}

		t.Logf("node %s is ready, proceeding with validation commands...", vmssName)

		err = runLiveVMValidators(ctx, t, vmssName, vmPrivateIP, string(privateKeyBytes), opts)
		require.NoError(t, err)

		t.Logf("node %s bootstrapping succeeded!", vmssName)
	}
}
