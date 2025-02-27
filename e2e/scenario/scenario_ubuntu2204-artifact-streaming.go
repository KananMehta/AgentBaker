package scenario

import (
	"github.com/Azure/agentbaker/pkg/agent/datamodel"
	"github.com/Azure/agentbakere2e/cluster"
	"github.com/Azure/agentbakere2e/config"
)

func ubuntu2204ArtifactStreaming() *Scenario {
	return &Scenario{
		Name:        "ubuntu2204-artifact-streaming",
		Description: "tests that a new ubuntu 2204 node using artifact streaming can be properly bootstrapepd",
		Tags: Tags{
			Name:     "ubuntu2204-artifact-streaming",
			OS:       "ubuntu2204",
			Platform: "x64",
		},
		Config: Config{
			Cluster:     cluster.ClusterKubenet,
			VHDSelector: config.VHDUbuntu2204Gen2Containerd,
			BootstrapConfigMutator: func(nbc *datamodel.NodeBootstrappingConfiguration) {
				nbc.EnableArtifactStreaming = true
				nbc.ContainerService.Properties.AgentPoolProfiles[0].Distro = "aks-ubuntu-containerd-22.04-gen2"
				nbc.AgentPoolProfile.Distro = "aks-ubuntu-containerd-22.04-gen2"
			},
			LiveVMValidators: []*LiveVMValidator{
				NonEmptyDirectoryValidator("/etc/overlaybd"),
			},
		},
	}
}
