package scenario

import (
	"github.com/Azure/agentbaker/pkg/agent/datamodel"
	"github.com/Azure/agentbakere2e/cluster"
	"github.com/Azure/agentbakere2e/config"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

func azurelinuxv2ARM64() *Scenario {
	return &Scenario{
		Name:        "azurelinuxv2-arm64",
		Description: "Tests that a node using a AzureLinuxV2 (CgroupV2) VHD on ARM64 architecture can be properly bootstrapped",
		Tags: Tags{
			Name:     "azurelinuxv2-arm64",
			OS:       "azurelinuxv2",
			Platform: "arm64",
		},
		Config: Config{
			Cluster:     cluster.ClusterKubenet,
			VHDSelector: config.VHDAzureLinuxV2Gen2Arm64,
			BootstrapConfigMutator: func(nbc *datamodel.NodeBootstrappingConfiguration) {
				nbc.ContainerService.Properties.AgentPoolProfiles[0].VMSize = "Standard_D2pds_V5"
				nbc.ContainerService.Properties.AgentPoolProfiles[0].Distro = "aks-azurelinux-v2-arm64-gen2"
				nbc.ContainerService.Properties.OrchestratorProfile.KubernetesConfig.CustomKubeBinaryURL = "https://acs-mirror.azureedge.net/kubernetes/v1.24.9/binaries/kubernetes-node-linux-arm64.tar.gz"
				nbc.AgentPoolProfile.VMSize = "Standard_D2pds_V5"
				nbc.AgentPoolProfile.Distro = "aks-azurelinux-v2-arm64-gen2"
				nbc.IsARM64 = true
			},
			VMConfigMutator: func(vmss *armcompute.VirtualMachineScaleSet) {
				vmss.SKU.Name = to.Ptr("Standard_D2pds_V5")
			},
		},
	}
}
