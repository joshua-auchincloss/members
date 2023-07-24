// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: api/v1/registry/pkg/def.proto

package pkg

import (
	common "members/grpc/api/v1/common"
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NewVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectKey string                    `protobuf:"bytes,1,opt,name=project_key,json=projectKey,proto3" json:"project_key,omitempty"`
	Version    string                    `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Protos     []*common.RegisteredProto `protobuf:"bytes,3,rep,name=protos,proto3" json:"protos,omitempty"`
}

func (x *NewVersionRequest) Reset() {
	*x = NewVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_registry_pkg_def_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewVersionRequest) ProtoMessage() {}

func (x *NewVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_registry_pkg_def_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewVersionRequest.ProtoReflect.Descriptor instead.
func (*NewVersionRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_registry_pkg_def_proto_rawDescGZIP(), []int{0}
}

func (x *NewVersionRequest) GetProjectKey() string {
	if x != nil {
		return x.ProjectKey
	}
	return ""
}

func (x *NewVersionRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *NewVersionRequest) GetProtos() []*common.RegisteredProto {
	if x != nil {
		return x.Protos
	}
	return nil
}

var File_api_v1_registry_pkg_def_proto protoreflect.FileDescriptor

var file_api_v1_registry_pkg_def_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x17, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x1a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a,
	0x11, 0x4e, 0x65, 0x77, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x42, 0xc9, 0x01, 0x0a, 0x1b, 0x63, 0x6f,
	0x6d, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x42, 0x08, 0x44, 0x65, 0x66, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x20, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0xa2, 0x02, 0x04, 0x4d, 0x56, 0x52, 0x50, 0xaa, 0x02,
	0x17, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2e, 0x50, 0x6b, 0x67, 0xca, 0x02, 0x17, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5c, 0x50,
	0x6b, 0x67, 0xe2, 0x02, 0x23, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x5c, 0x56, 0x31, 0x5c,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5c, 0x50, 0x6b, 0x67, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1a, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x3a, 0x3a, 0x50, 0x6b, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_registry_pkg_def_proto_rawDescOnce sync.Once
	file_api_v1_registry_pkg_def_proto_rawDescData = file_api_v1_registry_pkg_def_proto_rawDesc
)

func file_api_v1_registry_pkg_def_proto_rawDescGZIP() []byte {
	file_api_v1_registry_pkg_def_proto_rawDescOnce.Do(func() {
		file_api_v1_registry_pkg_def_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_registry_pkg_def_proto_rawDescData)
	})
	return file_api_v1_registry_pkg_def_proto_rawDescData
}

var file_api_v1_registry_pkg_def_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_v1_registry_pkg_def_proto_goTypes = []interface{}{
	(*NewVersionRequest)(nil),      // 0: members.v1.registry.pkg.NewVersionRequest
	(*common.RegisteredProto)(nil), // 1: members.v1.common.RegisteredProto
}
var file_api_v1_registry_pkg_def_proto_depIdxs = []int32{
	1, // 0: members.v1.registry.pkg.NewVersionRequest.protos:type_name -> members.v1.common.RegisteredProto
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_v1_registry_pkg_def_proto_init() }
func file_api_v1_registry_pkg_def_proto_init() {
	if File_api_v1_registry_pkg_def_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_registry_pkg_def_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewVersionRequest); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_registry_pkg_def_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1_registry_pkg_def_proto_goTypes,
		DependencyIndexes: file_api_v1_registry_pkg_def_proto_depIdxs,
		MessageInfos:      file_api_v1_registry_pkg_def_proto_msgTypes,
	}.Build()
	File_api_v1_registry_pkg_def_proto = out.File
	file_api_v1_registry_pkg_def_proto_rawDesc = nil
	file_api_v1_registry_pkg_def_proto_goTypes = nil
	file_api_v1_registry_pkg_def_proto_depIdxs = nil
}
