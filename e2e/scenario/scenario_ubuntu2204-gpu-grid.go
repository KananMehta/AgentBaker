package scenario

import (
	"github.com/Azure/agentbaker/pkg/agent/datamodel"
	"github.com/Azure/agentbakere2e/cluster"
	"github.com/Azure/agentbakere2e/config"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

func ubuntu2204GPUGridDriver() *Scenario {
	return &Scenario{
		Name:        "ubuntu2204-gpu-grid",
		Description: "Tests that a GPU-enabled node using the Ubuntu 2204 VHD with grid driver can be properly bootstrapped",
		Tags: Tags{
			Name:     "ubuntu2204-gpu-grid",
			OS:       "ubuntu2204",
			Platform: "x64",
			GPU:      true,
		},
		Config: Config{
			Cluster:     cluster.ClusterKubenet,
			VHDSelector: config.VHDUbuntu2204Gen2Containerd,
			BootstrapConfigMutator: func(nbc *datamodel.NodeBootstrappingConfiguration) {
				nbc.ContainerService.Properties.AgentPoolProfiles[0].Distro = "aks-ubuntu-containerd-22.04-gen2"
				nbc.AgentPoolProfile.Distro = "aks-ubuntu-containerd-22.04-gen2"
				nbc.AgentPoolProfile.VMSize = "Standard_NV6ads_A10_v5"
				nbc.ConfigGPUDriverIfNeeded = true
				nbc.EnableGPUDevicePluginIfNeeded = false
				nbc.EnableNvidia = true
			},
			VMConfigMutator: func(vmss *armcompute.VirtualMachineScaleSet) {
				vmss.SKU.Name = to.Ptr("Standard_NV6ads_A10_v5")
			},
			LiveVMValidators: []*LiveVMValidator{
				NvidiaSMIInstalledValidator(),
			},
		},
	}
}
