// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: pkg/proto/nbcontract/v1/config.proto

package nbcontractv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WorkloadRuntime int32

const (
	WorkloadRuntime_WR_UNSPECIFIED WorkloadRuntime = 0
	WorkloadRuntime_OCI_CONTAINER  WorkloadRuntime = 1
	WorkloadRuntime_WASM_WASI      WorkloadRuntime = 2
)

// Enum value maps for WorkloadRuntime.
var (
	WorkloadRuntime_name = map[int32]string{
		0: "WR_UNSPECIFIED",
		1: "OCI_CONTAINER",
		2: "WASM_WASI",
	}
	WorkloadRuntime_value = map[string]int32{
		"WR_UNSPECIFIED": 0,
		"OCI_CONTAINER":  1,
		"WASM_WASI":      2,
	}
)

func (x WorkloadRuntime) Enum() *WorkloadRuntime {
	p := new(WorkloadRuntime)
	*p = x
	return p
}

func (x WorkloadRuntime) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WorkloadRuntime) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_proto_nbcontract_v1_config_proto_enumTypes[0].Descriptor()
}

func (WorkloadRuntime) Type() protoreflect.EnumType {
	return &file_pkg_proto_nbcontract_v1_config_proto_enumTypes[0]
}

func (x WorkloadRuntime) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorkloadRuntime.Descriptor instead.
func (WorkloadRuntime) EnumDescriptor() ([]byte, []int) {
	return file_pkg_proto_nbcontract_v1_config_proto_rawDescGZIP(), []int{0}
}

type Configuration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Semantic version of this node bootstrap contract
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// TLS bootstrap config
	TlsBootstrappingConfig *TLSBootstrappingConfig `protobuf:"bytes,2,opt,name=tls_bootstrapping_config,json=tlsBootstrappingConfig,proto3" json:"tls_bootstrapping_config,omitempty"`
	// Kube binary URL config
	KubeBinaryConfig *KubeBinaryConfig `protobuf:"bytes,3,opt,name=kube_binary_config,json=kubeBinaryConfig,proto3" json:"kube_binary_config,omitempty"`
	// Custom cloud config
	CustomCloudConfig *CustomCloudConfig `protobuf:"bytes,4,opt,name=custom_cloud_config,json=customCloudConfig,proto3" json:"custom_cloud_config,omitempty"`
	// Kubernetes API server configuration
	ApiServerConfig *ApiServerConfig `protobuf:"bytes,5,opt,name=api_server_config,json=apiServerConfig,proto3" json:"api_server_config,omitempty"`
	// Various Kubernetes cluster level configuration
	ClusterConfig *ClusterConfig `protobuf:"bytes,6,opt,name=cluster_config,json=clusterConfig,proto3" json:"cluster_config,omitempty"`
	// Authentication configuration
	AuthConfig *AuthConfig `protobuf:"bytes,7,opt,name=auth_config,json=authConfig,proto3" json:"auth_config,omitempty"`
	// The CLI tool runc configuration
	RuncConfig *RuncConfig `protobuf:"bytes,8,opt,name=runc_config,json=runcConfig,proto3" json:"runc_config,omitempty"`
	// Containerd configuration
	ContainerdConfig *ContainerdConfig `protobuf:"bytes,9,opt,name=containerd_config,json=containerdConfig,proto3" json:"containerd_config,omitempty"`
	// Teleport configuration
	TeleportConfig *TeleportConfig `protobuf:"bytes,10,opt,name=teleport_config,json=teleportConfig,proto3" json:"teleport_config,omitempty"`
	// Kubelet configuration
	KubeletConfig *KubeletConfig `protobuf:"bytes,11,opt,name=kubelet_config,json=kubeletConfig,proto3" json:"kubelet_config,omitempty"`
	// Custom search domain configurations
	CustomSearchDomainConfig *CustomSearchDomainConfig `protobuf:"bytes,12,opt,name=custom_search_domain_config,json=customSearchDomainConfig,proto3" json:"custom_search_domain_config,omitempty"`
	// Custom Linux OS configurations including SwapFile, SysCtl configs, etc.
	CustomLinuxOsConfig *CustomLinuxOSConfig `protobuf:"bytes,13,opt,name=custom_linux_os_config,json=customLinuxOsConfig,proto3" json:"custom_linux_os_config,omitempty"`
	// HTTP/HTTPS proxy configuration for the node
	HttpProxyConfig *HTTPProxyConfig `protobuf:"bytes,14,opt,name=http_proxy_config,json=httpProxyConfig,proto3" json:"http_proxy_config,omitempty"`
	// GPU configuration for the node
	GpuConfig *GPUConfig `protobuf:"bytes,15,opt,name=gpu_config,json=gpuConfig,proto3" json:"gpu_config,omitempty"`
	// Kubernetes certificate authority (CA) certificate, required by the node to establish TLS with the API server
	KubernetesCaCert string `protobuf:"bytes,16,opt,name=kubernetes_ca_cert,json=kubernetesCaCert,proto3" json:"kubernetes_ca_cert,omitempty"`
	// Cluster/user config
	KubernetesVersion string `protobuf:"bytes,17,opt,name=kubernetes_version,json=kubernetesVersion,proto3" json:"kubernetes_version,omitempty"` // Q: can this be auto-detected? Or is this part of specifying the desired node version?
	// Kube proxy URL
	KubeProxyUrl string `protobuf:"bytes,18,opt,name=kube_proxy_url,json=kubeProxyUrl,proto3" json:"kube_proxy_url,omitempty"`
	// The VM size of the node
	VmSize string `protobuf:"bytes,19,opt,name=vm_size,json=vmSize,proto3" json:"vm_size,omitempty"`
	// Linux admin username. If not specified, the default value is "azureuser"
	LinuxAdminUsername string `protobuf:"bytes,20,opt,name=linux_admin_username,json=linuxAdminUsername,proto3" json:"linux_admin_username,omitempty"`
	// Specifies whether the node is a VHD node. This is still needed for some customized scenarios.
	// This is labeled as optional (explicit presence) so that we know whether it's set or not.
	// If it's not set, the default value will be nil.
	IsVhd *bool `protobuf:"varint,21,opt,name=is_vhd,json=isVhd,proto3,oneof" json:"is_vhd,omitempty"`
	// Specifies whether SSH is enabled or disabled on the VM node
	EnableSsh bool `protobuf:"varint,22,opt,name=enable_ssh,json=enableSsh,proto3" json:"enable_ssh,omitempty"`
	// Specifies whether unattended upgrade is enabled or disabled on the VM node
	EnableUnattendedUpgrade bool `protobuf:"varint,23,opt,name=enable_unattended_upgrade,json=enableUnattendedUpgrade,proto3" json:"enable_unattended_upgrade,omitempty"`
	// The message of the day that is displayed on the VM node when a user logs in
	MessageOfTheDay string `protobuf:"bytes,24,opt,name=message_of_the_day,json=messageOfTheDay,proto3" json:"message_of_the_day,omitempty"`
	// Specifies whether the hosts config agent is enabled or disabled on the VM node
	EnableHostsConfigAgent bool `protobuf:"varint,25,opt,name=enable_hosts_config_agent,json=enableHostsConfigAgent,proto3" json:"enable_hosts_config_agent,omitempty"`
	// Custom CA certificates to be added to the system trust store
	CustomCaCerts []string `protobuf:"bytes,26,rep,name=custom_ca_certs,json=customCaCerts,proto3" json:"custom_ca_certs,omitempty"`
	// A local file path where cluster provision cse output should be stored
	ProvisionOutput string `protobuf:"bytes,27,opt,name=provision_output,json=provisionOutput,proto3" json:"provision_output,omitempty"`
	// Workload runtime, e.g., either "OCIContainer" or "WasmWasi", currently.
	WorkloadRuntime WorkloadRuntime `protobuf:"varint,28,opt,name=workload_runtime,json=workloadRuntime,proto3,enum=nbcontract.v1.WorkloadRuntime" json:"workload_runtime,omitempty"`
	// Specifies whether IPv6 dual stack is enabled or disabled on the VM node
	Ipv6DualStackEnabled bool `protobuf:"varint,29,opt,name=ipv6_dual_stack_enabled,json=ipv6DualStackEnabled,proto3" json:"ipv6_dual_stack_enabled,omitempty"`
	// Command to use for outbound traffic
	OutboundCommand string `protobuf:"bytes,30,opt,name=outbound_command,json=outboundCommand,proto3" json:"outbound_command,omitempty"`
	// specifies whether to ensure no duplicate promiscuous bridge
	EnsureNoDupePromiscuousBridge bool `protobuf:"varint,31,opt,name=ensure_no_dupe_promiscuous_bridge,json=ensureNoDupePromiscuousBridge,proto3" json:"ensure_no_dupe_promiscuous_bridge,omitempty"`
	// Azure private registry server URI
	AzurePrivateRegistryServer string `protobuf:"bytes,32,opt,name=azure_private_registry_server,json=azurePrivateRegistryServer,proto3" json:"azure_private_registry_server,omitempty"`
	// Private egress proxy address
	PrivateEgressProxyAddress string `protobuf:"bytes,33,opt,name=private_egress_proxy_address,json=privateEgressProxyAddress,proto3" json:"private_egress_proxy_address,omitempty"`
	// Specifies whether artifact streaming is enabled or disabled on the VM node
	EnableArtifactStreaming bool `protobuf:"varint,34,opt,name=enable_artifact_streaming,json=enableArtifactStreaming,proto3" json:"enable_artifact_streaming,omitempty"`
	// Specifies whether the node is a Kata node
	IsKata bool `protobuf:"varint,35,opt,name=is_kata,json=isKata,proto3" json:"is_kata,omitempty"`
	// Specifies whether the node needs cgroupv2.
	// Labeled as optional (explicit presence) so that we know whether it's set or not.
	// If it's not set, the default value will be nil and we will get the value on the VHD during bootstrapping.
	NeedsCgroupv2 *bool `protobuf:"varint,36,opt,name=needs_cgroupv2,json=needsCgroupv2,proto3,oneof" json:"needs_cgroupv2,omitempty"`
}

func (x *Configuration) Reset() {
	*x = Configuration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_nbcontract_v1_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Configuration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Configuration) ProtoMessage() {}

func (x *Configuration) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_nbcontract_v1_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Configuration.ProtoReflect.Descriptor instead.
func (*Configuration) Descriptor() ([]byte, []int) {
	return file_pkg_proto_nbcontract_v1_config_proto_rawDescGZIP(), []int{0}
}

func (x *Configuration) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Configuration) GetTlsBootstrappingConfig() *TLSBootstrappingConfig {
	if x != nil {
		return x.TlsBootstrappingConfig
	}
	return nil
}

func (x *Configuration) GetKubeBinaryConfig() *KubeBinaryConfig {
	if x != nil {
		return x.KubeBinaryConfig
	}
	return nil
}

func (x *Configuration) GetCustomCloudConfig() *CustomCloudConfig {
	if x != nil {
		return x.CustomCloudConfig
	}
	return nil
}

func (x *Configuration) GetApiServerConfig() *ApiServerConfig {
	if x != nil {
		return x.ApiServerConfig
	}
	return nil
}

func (x *Configuration) GetClusterConfig() *ClusterConfig {
	if x != nil {
		return x.ClusterConfig
	}
	return nil
}

func (x *Configuration) GetAuthConfig() *AuthConfig {
	if x != nil {
		return x.AuthConfig
	}
	return nil
}

func (x *Configuration) GetRuncConfig() *RuncConfig {
	if x != nil {
		return x.RuncConfig
	}
	return nil
}

func (x *Configuration) GetContainerdConfig() *ContainerdConfig {
	if x != nil {
		return x.ContainerdConfig
	}
	return nil
}

func (x *Configuration) GetTeleportConfig() *TeleportConfig {
	if x != nil {
		return x.TeleportConfig
	}
	return nil
}

func (x *Configuration) GetKubeletConfig() *KubeletConfig {
	if x != nil {
		return x.KubeletConfig
	}
	return nil
}

func (x *Configuration) GetCustomSearchDomainConfig() *CustomSearchDomainConfig {
	if x != nil {
		return x.CustomSearchDomainConfig
	}
	return nil
}

func (x *Configuration) GetCustomLinuxOsConfig() *CustomLinuxOSConfig {
	if x != nil {
		return x.CustomLinuxOsConfig
	}
	return nil
}

func (x *Configuration) GetHttpProxyConfig() *HTTPProxyConfig {
	if x != nil {
		return x.HttpProxyConfig
	}
	return nil
}

func (x *Configuration) GetGpuConfig() *GPUConfig {
	if x != nil {
		return x.GpuConfig
	}
	return nil
}

func (x *Configuration) GetKubernetesCaCert() string {
	if x != nil {
		return x.KubernetesCaCert
	}
	return ""
}

func (x *Configuration) GetKubernetesVersion() string {
	if x != nil {
		return x.KubernetesVersion
	}
	return ""
}

func (x *Configuration) GetKubeProxyUrl() string {
	if x != nil {
		return x.KubeProxyUrl
	}
	return ""
}

func (x *Configuration) GetVmSize() string {
	if x != nil {
		return x.VmSize
	}
	return ""
}

func (x *Configuration) GetLinuxAdminUsername() string {
	if x != nil {
		return x.LinuxAdminUsername
	}
	return ""
}

func (x *Configuration) GetIsVhd() bool {
	if x != nil && x.IsVhd != nil {
		return *x.IsVhd
	}
	return false
}

func (x *Configuration) GetEnableSsh() bool {
	if x != nil {
		return x.EnableSsh
	}
	return false
}

func (x *Configuration) GetEnableUnattendedUpgrade() bool {
	if x != nil {
		return x.EnableUnattendedUpgrade
	}
	return false
}

func (x *Configuration) GetMessageOfTheDay() string {
	if x != nil {
		return x.MessageOfTheDay
	}
	return ""
}

func (x *Configuration) GetEnableHostsConfigAgent() bool {
	if x != nil {
		return x.EnableHostsConfigAgent
	}
	return false
}

func (x *Configuration) GetCustomCaCerts() []string {
	if x != nil {
		return x.CustomCaCerts
	}
	return nil
}

func (x *Configuration) GetProvisionOutput() string {
	if x != nil {
		return x.ProvisionOutput
	}
	return ""
}

func (x *Configuration) GetWorkloadRuntime() WorkloadRuntime {
	if x != nil {
		return x.WorkloadRuntime
	}
	return WorkloadRuntime_WR_UNSPECIFIED
}

func (x *Configuration) GetIpv6DualStackEnabled() bool {
	if x != nil {
		return x.Ipv6DualStackEnabled
	}
	return false
}

func (x *Configuration) GetOutboundCommand() string {
	if x != nil {
		return x.OutboundCommand
	}
	return ""
}

func (x *Configuration) GetEnsureNoDupePromiscuousBridge() bool {
	if x != nil {
		return x.EnsureNoDupePromiscuousBridge
	}
	return false
}

func (x *Configuration) GetAzurePrivateRegistryServer() string {
	if x != nil {
		return x.AzurePrivateRegistryServer
	}
	return ""
}

func (x *Configuration) GetPrivateEgressProxyAddress() string {
	if x != nil {
		return x.PrivateEgressProxyAddress
	}
	return ""
}

func (x *Configuration) GetEnableArtifactStreaming() bool {
	if x != nil {
		return x.EnableArtifactStreaming
	}
	return false
}

func (x *Configuration) GetIsKata() bool {
	if x != nil {
		return x.IsKata
	}
	return false
}

func (x *Configuration) GetNeedsCgroupv2() bool {
	if x != nil && x.NeedsCgroupv2 != nil {
		return *x.NeedsCgroupv2
	}
	return false
}

var File_pkg_proto_nbcontract_v1_config_proto protoreflect.FileDescriptor

var file_pkg_proto_nbcontract_v1_config_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x62, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x2d, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x68,
	0x74, 0x74, 0x70, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x70,
	0x75, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x6c, 0x69, 0x6e,
	0x75, 0x78, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2e, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x62, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x62, 0x69,
	0x6e, 0x61, 0x72, 0x79, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x36, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x62, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2b, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x62,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x75, 0x62, 0x65,
	0x6c, 0x65, 0x74, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2c, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x62, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x75, 0x6e, 0x63, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x6c, 0x73, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x10, 0x0a, 0x0d, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x5f, 0x0a, 0x18, 0x74, 0x6c, 0x73, 0x5f, 0x62, 0x6f,
	0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6e, 0x62, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x4c, 0x53, 0x42, 0x6f, 0x6f, 0x74,
	0x73, 0x74, 0x72, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x16, 0x74, 0x6c, 0x73, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4d, 0x0a, 0x12, 0x6b, 0x75, 0x62, 0x65, 0x5f,
	0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x10, 0x6b, 0x75, 0x62, 0x65, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x50, 0x0a, 0x13, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x5f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x11, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4a, 0x0a, 0x11, 0x61, 0x70, 0x69, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x0f, 0x61, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x43, 0x0a, 0x0e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6e,
	0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0d, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3a, 0x0a, 0x0b, 0x61, 0x75, 0x74,
	0x68, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3a, 0x0a, 0x0b, 0x72, 0x75, 0x6e, 0x63, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6e, 0x62, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x63, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0a, 0x72, 0x75, 0x6e, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x4c, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6e,
	0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x10, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x46, 0x0a, 0x0f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x62, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x43, 0x0a, 0x0e, 0x6b, 0x75, 0x62, 0x65, 0x6c,
	0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x4b, 0x75, 0x62, 0x65, 0x6c, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0d, 0x6b,
	0x75, 0x62, 0x65, 0x6c, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x66, 0x0a, 0x1b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x18, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x57, 0x0a, 0x16, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6c,
	0x69, 0x6e, 0x75, 0x78, 0x5f, 0x6f, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4c, 0x69, 0x6e, 0x75, 0x78,
	0x4f, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x13, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x4c, 0x69, 0x6e, 0x75, 0x78, 0x4f, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4a, 0x0a,
	0x11, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6e, 0x62, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x50, 0x72, 0x6f,
	0x78, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0f, 0x68, 0x74, 0x74, 0x70, 0x50, 0x72,
	0x6f, 0x78, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x37, 0x0a, 0x0a, 0x67, 0x70, 0x75,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x50,
	0x55, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x09, 0x67, 0x70, 0x75, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x2c, 0x0a, 0x12, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x5f, 0x63, 0x61, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x61, 0x43, 0x65, 0x72, 0x74,
	0x12, 0x2d, 0x0a, 0x12, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x24, 0x0a, 0x0e, 0x6b, 0x75, 0x62, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6b, 0x75, 0x62, 0x65, 0x50, 0x72, 0x6f,
	0x78, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x76, 0x6d, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x6d, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x30,
	0x0a, 0x14, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6c, 0x69,
	0x6e, 0x75, 0x78, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x06, 0x69, 0x73, 0x5f, 0x76, 0x68, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x00, 0x52, 0x05, 0x69, 0x73, 0x56, 0x68, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x0a,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x73, 0x68, 0x18, 0x16, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x73, 0x68, 0x12, 0x3a, 0x0a, 0x19, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x75, 0x6e, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64,
	0x5f, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x55, 0x6e, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64,
	0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x12, 0x2b, 0x0a, 0x12, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x74, 0x68, 0x65, 0x5f, 0x64, 0x61, 0x79, 0x18, 0x18, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x66, 0x54, 0x68,
	0x65, 0x44, 0x61, 0x79, 0x12, 0x39, 0x0a, 0x19, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x68,
	0x6f, 0x73, 0x74, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x18, 0x19, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x48,
	0x6f, 0x73, 0x74, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12,
	0x26, 0x0a, 0x0f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x63, 0x61, 0x5f, 0x63, 0x65, 0x72,
	0x74, 0x73, 0x18, 0x1a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x43, 0x61, 0x43, 0x65, 0x72, 0x74, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x1b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x12, 0x49, 0x0a, 0x10, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x72,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x6e,
	0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x0f, 0x77, 0x6f,
	0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a,
	0x17, 0x69, 0x70, 0x76, 0x36, 0x5f, 0x64, 0x75, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14,
	0x69, 0x70, 0x76, 0x36, 0x44, 0x75, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x6f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x6f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12,
	0x48, 0x0a, 0x21, 0x65, 0x6e, 0x73, 0x75, 0x72, 0x65, 0x5f, 0x6e, 0x6f, 0x5f, 0x64, 0x75, 0x70,
	0x65, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x63, 0x75, 0x6f, 0x75, 0x73, 0x5f, 0x62, 0x72,
	0x69, 0x64, 0x67, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1d, 0x65, 0x6e, 0x73, 0x75,
	0x72, 0x65, 0x4e, 0x6f, 0x44, 0x75, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x63, 0x75,
	0x6f, 0x75, 0x73, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x12, 0x41, 0x0a, 0x1d, 0x61, 0x7a, 0x75,
	0x72, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x20, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x1a, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x1c,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x21, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x19, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x45, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x3a, 0x0a,
	0x19, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x18, 0x22, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x17, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f,
	0x6b, 0x61, 0x74, 0x61, 0x18, 0x23, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x4b, 0x61,
	0x74, 0x61, 0x12, 0x2a, 0x0a, 0x0e, 0x6e, 0x65, 0x65, 0x64, 0x73, 0x5f, 0x63, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x76, 0x32, 0x18, 0x24, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52, 0x0d, 0x6e, 0x65,
	0x65, 0x64, 0x73, 0x43, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x76, 0x32, 0x88, 0x01, 0x01, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x69, 0x73, 0x5f, 0x76, 0x68, 0x64, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x6e, 0x65,
	0x65, 0x64, 0x73, 0x5f, 0x63, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x76, 0x32, 0x2a, 0x47, 0x0a, 0x0f,
	0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x0e, 0x57, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x4f, 0x43, 0x49, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x41,
	0x49, 0x4e, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x57, 0x41, 0x53, 0x4d, 0x5f, 0x57,
	0x41, 0x53, 0x49, 0x10, 0x02, 0x42, 0xb7, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x62,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x2f, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x42, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x3b,
	0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4e,
	0x58, 0x58, 0xaa, 0x02, 0x0d, 0x4e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x0d, 0x4e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x19, 0x4e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0e, 0x4e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_nbcontract_v1_config_proto_rawDescOnce sync.Once
	file_pkg_proto_nbcontract_v1_config_proto_rawDescData = file_pkg_proto_nbcontract_v1_config_proto_rawDesc
)

func file_pkg_proto_nbcontract_v1_config_proto_rawDescGZIP() []byte {
	file_pkg_proto_nbcontract_v1_config_proto_rawDescOnce.Do(func() {
		file_pkg_proto_nbcontract_v1_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_nbcontract_v1_config_proto_rawDescData)
	})
	return file_pkg_proto_nbcontract_v1_config_proto_rawDescData
}

var file_pkg_proto_nbcontract_v1_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_proto_nbcontract_v1_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_proto_nbcontract_v1_config_proto_goTypes = []interface{}{
	(WorkloadRuntime)(0),             // 0: nbcontract.v1.WorkloadRuntime
	(*Configuration)(nil),            // 1: nbcontract.v1.Configuration
	(*TLSBootstrappingConfig)(nil),   // 2: nbcontract.v1.TLSBootstrappingConfig
	(*KubeBinaryConfig)(nil),         // 3: nbcontract.v1.KubeBinaryConfig
	(*CustomCloudConfig)(nil),        // 4: nbcontract.v1.CustomCloudConfig
	(*ApiServerConfig)(nil),          // 5: nbcontract.v1.ApiServerConfig
	(*ClusterConfig)(nil),            // 6: nbcontract.v1.ClusterConfig
	(*AuthConfig)(nil),               // 7: nbcontract.v1.AuthConfig
	(*RuncConfig)(nil),               // 8: nbcontract.v1.RuncConfig
	(*ContainerdConfig)(nil),         // 9: nbcontract.v1.ContainerdConfig
	(*TeleportConfig)(nil),           // 10: nbcontract.v1.TeleportConfig
	(*KubeletConfig)(nil),            // 11: nbcontract.v1.KubeletConfig
	(*CustomSearchDomainConfig)(nil), // 12: nbcontract.v1.CustomSearchDomainConfig
	(*CustomLinuxOSConfig)(nil),      // 13: nbcontract.v1.CustomLinuxOSConfig
	(*HTTPProxyConfig)(nil),          // 14: nbcontract.v1.HTTPProxyConfig
	(*GPUConfig)(nil),                // 15: nbcontract.v1.GPUConfig
}
var file_pkg_proto_nbcontract_v1_config_proto_depIdxs = []int32{
	2,  // 0: nbcontract.v1.Configuration.tls_bootstrapping_config:type_name -> nbcontract.v1.TLSBootstrappingConfig
	3,  // 1: nbcontract.v1.Configuration.kube_binary_config:type_name -> nbcontract.v1.KubeBinaryConfig
	4,  // 2: nbcontract.v1.Configuration.custom_cloud_config:type_name -> nbcontract.v1.CustomCloudConfig
	5,  // 3: nbcontract.v1.Configuration.api_server_config:type_name -> nbcontract.v1.ApiServerConfig
	6,  // 4: nbcontract.v1.Configuration.cluster_config:type_name -> nbcontract.v1.ClusterConfig
	7,  // 5: nbcontract.v1.Configuration.auth_config:type_name -> nbcontract.v1.AuthConfig
	8,  // 6: nbcontract.v1.Configuration.runc_config:type_name -> nbcontract.v1.RuncConfig
	9,  // 7: nbcontract.v1.Configuration.containerd_config:type_name -> nbcontract.v1.ContainerdConfig
	10, // 8: nbcontract.v1.Configuration.teleport_config:type_name -> nbcontract.v1.TeleportConfig
	11, // 9: nbcontract.v1.Configuration.kubelet_config:type_name -> nbcontract.v1.KubeletConfig
	12, // 10: nbcontract.v1.Configuration.custom_search_domain_config:type_name -> nbcontract.v1.CustomSearchDomainConfig
	13, // 11: nbcontract.v1.Configuration.custom_linux_os_config:type_name -> nbcontract.v1.CustomLinuxOSConfig
	14, // 12: nbcontract.v1.Configuration.http_proxy_config:type_name -> nbcontract.v1.HTTPProxyConfig
	15, // 13: nbcontract.v1.Configuration.gpu_config:type_name -> nbcontract.v1.GPUConfig
	0,  // 14: nbcontract.v1.Configuration.workload_runtime:type_name -> nbcontract.v1.WorkloadRuntime
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_pkg_proto_nbcontract_v1_config_proto_init() }
func file_pkg_proto_nbcontract_v1_config_proto_init() {
	if File_pkg_proto_nbcontract_v1_config_proto != nil {
		return
	}
	file_pkg_proto_nbcontract_v1_httpproxyconfig_proto_init()
	file_pkg_proto_nbcontract_v1_gpuconfig_proto_init()
	file_pkg_proto_nbcontract_v1_customlinuxosconfig_proto_init()
	file_pkg_proto_nbcontract_v1_kubebinaryconfig_proto_init()
	file_pkg_proto_nbcontract_v1_customsearchdomainconfig_proto_init()
	file_pkg_proto_nbcontract_v1_containerdconfig_proto_init()
	file_pkg_proto_nbcontract_v1_authconfig_proto_init()
	file_pkg_proto_nbcontract_v1_kubeletconfig_proto_init()
	file_pkg_proto_nbcontract_v1_teleportconfig_proto_init()
	file_pkg_proto_nbcontract_v1_runcconfig_proto_init()
	file_pkg_proto_nbcontract_v1_customcloudconfig_proto_init()
	file_pkg_proto_nbcontract_v1_apiserverconfig_proto_init()
	file_pkg_proto_nbcontract_v1_tlsbootstrappingconfig_proto_init()
	file_pkg_proto_nbcontract_v1_clusterconfig_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_nbcontract_v1_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Configuration); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_pkg_proto_nbcontract_v1_config_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_proto_nbcontract_v1_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_nbcontract_v1_config_proto_goTypes,
		DependencyIndexes: file_pkg_proto_nbcontract_v1_config_proto_depIdxs,
		EnumInfos:         file_pkg_proto_nbcontract_v1_config_proto_enumTypes,
		MessageInfos:      file_pkg_proto_nbcontract_v1_config_proto_msgTypes,
	}.Build()
	File_pkg_proto_nbcontract_v1_config_proto = out.File
	file_pkg_proto_nbcontract_v1_config_proto_rawDesc = nil
	file_pkg_proto_nbcontract_v1_config_proto_goTypes = nil
	file_pkg_proto_nbcontract_v1_config_proto_depIdxs = nil
}
