// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: api/v1/registry/svc.proto

package registry

import (
	common "members/grpc/api/v1/common"
	pkg "members/grpc/api/v1/registry/pkg"
	reflect "reflect"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_api_v1_registry_svc_proto protoreflect.FileDescriptor

var file_api_v1_registry_svc_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2f, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2e, 0x73, 0x76, 0x63, 0x1a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x67,
	0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x5b, 0x0a, 0x0f, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x2e,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x4e, 0x65, 0x77, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x4d, 0x65, 0x74, 0x61, 0x42, 0xc5, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2e, 0x73, 0x76, 0x63, 0x42, 0x08, 0x53, 0x76, 0x63, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x1c, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0xa2, 0x02, 0x04, 0x4d, 0x56, 0x52, 0x53, 0xaa, 0x02, 0x17, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x53,
	0x76, 0x63, 0xca, 0x02, 0x17, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x5c, 0x56, 0x31, 0x5c,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5c, 0x53, 0x76, 0x63, 0xe2, 0x02, 0x23, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x5c, 0x53, 0x76, 0x63, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x1a, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x3a, 0x3a, 0x56, 0x31,
	0x3a, 0x3a, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x3a, 0x3a, 0x53, 0x76, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_v1_registry_svc_proto_goTypes = []interface{}{
	(*pkg.NewVersionRequest)(nil), // 0: members.v1.registry.pkg.NewVersionRequest
	(*common.ProtoMeta)(nil),      // 1: members.v1.common.ProtoMeta
}
var file_api_v1_registry_svc_proto_depIdxs = []int32{
	0, // 0: members.v1.registry.svc.Registry.RegisterVersion:input_type -> members.v1.registry.pkg.NewVersionRequest
	1, // 1: members.v1.registry.svc.Registry.RegisterVersion:output_type -> members.v1.common.ProtoMeta
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1_registry_svc_proto_init() }
func file_api_v1_registry_svc_proto_init() {
	if File_api_v1_registry_svc_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_registry_svc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_registry_svc_proto_goTypes,
		DependencyIndexes: file_api_v1_registry_svc_proto_depIdxs,
	}.Build()
	File_api_v1_registry_svc_proto = out.File
	file_api_v1_registry_svc_proto_rawDesc = nil
	file_api_v1_registry_svc_proto_goTypes = nil
	file_api_v1_registry_svc_proto_depIdxs = nil
}
