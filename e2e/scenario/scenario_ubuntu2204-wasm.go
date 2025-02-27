package scenario

import (
	"github.com/Azure/agentbaker/pkg/agent/datamodel"
	"github.com/Azure/agentbakere2e/cluster"
	"github.com/Azure/agentbakere2e/config"
)

func ubuntu2204Wasm() *Scenario {
	return &Scenario{
		Name:        "ubuntu2204-wasm",
		Description: "tests that a new ubuntu 2204 node using krustlet can be properly bootstrapepd",
		Tags: Tags{
			Name:     "ubuntu2204-wasm",
			OS:       "ubuntu2204",
			Platform: "x64",
		},
		Config: Config{
			Cluster:     cluster.ClusterKubenet,
			VHDSelector: config.VHDUbuntu2204Gen2Containerd,
			BootstrapConfigMutator: func(nbc *datamodel.NodeBootstrappingConfiguration) {
				nbc.ContainerService.Properties.AgentPoolProfiles[0].WorkloadRuntime = datamodel.WasmWasi
				nbc.ContainerService.Properties.AgentPoolProfiles[0].Distro = "aks-ubuntu-containerd-22.04-gen2"
				nbc.AgentPoolProfile.WorkloadRuntime = datamodel.WasmWasi
				nbc.AgentPoolProfile.Distro = "aks-ubuntu-containerd-22.04-gen2"
			},
		},
	}
}
