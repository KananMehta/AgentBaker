// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: pkg/proto/nbcontract/v1/gpuconfig.proto

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

type GPUConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// NvidiaState represents the Nvidia GPU state
	NvidiaState *FeatureState `protobuf:"varint,1,opt,name=nvidia_state,json=nvidiaState,proto3,enum=nbcontract.v1.FeatureState,oneof" json:"nvidia_state,omitempty"`
	// ConfigGpuDriver specifies whether bootstrap process should install and configure the GPU driver when necessary. Configuration includes appropriate set up of components like the fabric manager where applicable.
	ConfigGpuDriver *FeatureState `protobuf:"varint,2,opt,name=config_gpu_driver,json=configGpuDriver,proto3,enum=nbcontract.v1.FeatureState,oneof" json:"config_gpu_driver,omitempty"`
	// GpuDevicePlugin specifies whether special config is needed for MIG GPUs that use GPU dedicated VHDs and enable the device plugin (for all GPU dedicated VHDs)
	GpuDevicePlugin *FeatureState `protobuf:"varint,3,opt,name=gpu_device_plugin,json=gpuDevicePlugin,proto3,enum=nbcontract.v1.FeatureState,oneof" json:"gpu_device_plugin,omitempty"`
	// GpuInstanceProfile represents the GPU instance profile.
	GpuInstanceProfile *string `protobuf:"bytes,4,opt,name=gpu_instance_profile,json=gpuInstanceProfile,proto3,oneof" json:"gpu_instance_profile,omitempty"`
	// GpuImageSha represents AKS-GPU image's SHA
	GpuImageSha *string `protobuf:"bytes,5,opt,name=gpu_image_sha,json=gpuImageSha,proto3,oneof" json:"gpu_image_sha,omitempty"`
}

func (x *GPUConfig) Reset() {
	*x = GPUConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_nbcontract_v1_gpuconfig_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GPUConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GPUConfig) ProtoMessage() {}

func (x *GPUConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_nbcontract_v1_gpuconfig_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GPUConfig.ProtoReflect.Descriptor instead.
func (*GPUConfig) Descriptor() ([]byte, []int) {
	return file_pkg_proto_nbcontract_v1_gpuconfig_proto_rawDescGZIP(), []int{0}
}

func (x *GPUConfig) GetNvidiaState() FeatureState {
	if x != nil && x.NvidiaState != nil {
		return *x.NvidiaState
	}
	return FeatureState_FEATURE_STATE_UNSPECIFIED
}

func (x *GPUConfig) GetConfigGpuDriver() FeatureState {
	if x != nil && x.ConfigGpuDriver != nil {
		return *x.ConfigGpuDriver
	}
	return FeatureState_FEATURE_STATE_UNSPECIFIED
}

func (x *GPUConfig) GetGpuDevicePlugin() FeatureState {
	if x != nil && x.GpuDevicePlugin != nil {
		return *x.GpuDevicePlugin
	}
	return FeatureState_FEATURE_STATE_UNSPECIFIED
}

func (x *GPUConfig) GetGpuInstanceProfile() string {
	if x != nil && x.GpuInstanceProfile != nil {
		return *x.GpuInstanceProfile
	}
	return ""
}

func (x *GPUConfig) GetGpuImageSha() string {
	if x != nil && x.GpuImageSha != nil {
		return *x.GpuImageSha
	}
	return ""
}

var File_pkg_proto_nbcontract_v1_gpuconfig_proto protoreflect.FileDescriptor

var file_pkg_proto_nbcontract_v1_gpuconfig_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x62, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x70, 0x75, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6e, 0x62, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x2a, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x76,
	0x31, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x03, 0x0a, 0x09, 0x47, 0x50, 0x55, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x43, 0x0a, 0x0c, 0x6e, 0x76, 0x69, 0x64, 0x69, 0x61, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x6e, 0x62, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x6e, 0x76, 0x69, 0x64, 0x69, 0x61, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x4c, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x5f, 0x67, 0x70, 0x75, 0x5f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x48,
	0x01, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x47, 0x70, 0x75, 0x44, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x4c, 0x0a, 0x11, 0x67, 0x70, 0x75, 0x5f, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1b, 0x2e, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x48, 0x02, 0x52,
	0x0f, 0x67, 0x70, 0x75, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x88, 0x01, 0x01, 0x12, 0x35, 0x0a, 0x14, 0x67, 0x70, 0x75, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x03, 0x52, 0x12, 0x67, 0x70, 0x75, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0d, 0x67, 0x70,
	0x75, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x68, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x04, 0x52, 0x0b, 0x67, 0x70, 0x75, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x68, 0x61,
	0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x6e, 0x76, 0x69, 0x64, 0x69, 0x61, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f,
	0x67, 0x70, 0x75, 0x5f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x67,
	0x70, 0x75, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x42, 0x17, 0x0a, 0x15, 0x5f, 0x67, 0x70, 0x75, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x67, 0x70,
	0x75, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x68, 0x61, 0x42, 0xba, 0x01, 0x0a, 0x11,
	0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76,
	0x31, 0x42, 0x0e, 0x47, 0x70, 0x75, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x41, 0x7a, 0x75, 0x72, 0x65, 0x2f, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x6b, 0x65, 0x72,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x62, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4e, 0x58, 0x58, 0xaa, 0x02, 0x0d, 0x4e, 0x62,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0d, 0x4e, 0x62,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x19, 0x4e, 0x62,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x4e, 0x62, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_nbcontract_v1_gpuconfig_proto_rawDescOnce sync.Once
	file_pkg_proto_nbcontract_v1_gpuconfig_proto_rawDescData = file_pkg_proto_nbcontract_v1_gpuconfig_proto_rawDesc
)

func file_pkg_proto_nbcontract_v1_gpuconfig_proto_rawDescGZIP() []byte {
	file_pkg_proto_nbcontract_v1_gpuconfig_proto_rawDescOnce.Do(func() {
		file_pkg_proto_nbcontract_v1_gpuconfig_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_nbcontract_v1_gpuconfig_proto_rawDescData)
	})
	return file_pkg_proto_nbcontract_v1_gpuconfig_proto_rawDescData
}

var file_pkg_proto_nbcontract_v1_gpuconfig_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_proto_nbcontract_v1_gpuconfig_proto_goTypes = []interface{}{
	(*GPUConfig)(nil), // 0: nbcontract.v1.GPUConfig
	(FeatureState)(0), // 1: nbcontract.v1.FeatureState
}
var file_pkg_proto_nbcontract_v1_gpuconfig_proto_depIdxs = []int32{
	1, // 0: nbcontract.v1.GPUConfig.nvidia_state:type_name -> nbcontract.v1.FeatureState
	1, // 1: nbcontract.v1.GPUConfig.config_gpu_driver:type_name -> nbcontract.v1.FeatureState
	1, // 2: nbcontract.v1.GPUConfig.gpu_device_plugin:type_name -> nbcontract.v1.FeatureState
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_proto_nbcontract_v1_gpuconfig_proto_init() }
func file_pkg_proto_nbcontract_v1_gpuconfig_proto_init() {
	if File_pkg_proto_nbcontract_v1_gpuconfig_proto != nil {
		return
	}
	file_pkg_proto_nbcontract_v1_featurestate_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_nbcontract_v1_gpuconfig_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GPUConfig); i {
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
	file_pkg_proto_nbcontract_v1_gpuconfig_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_proto_nbcontract_v1_gpuconfig_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_nbcontract_v1_gpuconfig_proto_goTypes,
		DependencyIndexes: file_pkg_proto_nbcontract_v1_gpuconfig_proto_depIdxs,
		MessageInfos:      file_pkg_proto_nbcontract_v1_gpuconfig_proto_msgTypes,
	}.Build()
	File_pkg_proto_nbcontract_v1_gpuconfig_proto = out.File
	file_pkg_proto_nbcontract_v1_gpuconfig_proto_rawDesc = nil
	file_pkg_proto_nbcontract_v1_gpuconfig_proto_goTypes = nil
	file_pkg_proto_nbcontract_v1_gpuconfig_proto_depIdxs = nil
}
