// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: api/v1/common/image.proto

package common

import (
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

type RegisteredProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FileName   string `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	Data       []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	LastUpdate string `protobuf:"bytes,4,opt,name=last_update,json=lastUpdate,proto3" json:"last_update,omitempty"`
}

func (x *RegisteredProto) Reset() {
	*x = RegisteredProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_common_image_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisteredProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisteredProto) ProtoMessage() {}

func (x *RegisteredProto) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_common_image_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisteredProto.ProtoReflect.Descriptor instead.
func (*RegisteredProto) Descriptor() ([]byte, []int) {
	return file_api_v1_common_image_proto_rawDescGZIP(), []int{0}
}

func (x *RegisteredProto) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RegisteredProto) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *RegisteredProto) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *RegisteredProto) GetLastUpdate() string {
	if x != nil {
		return x.LastUpdate
	}
	return ""
}

type ProtoMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProjectKey   string `protobuf:"bytes,2,opt,name=project_key,json=projectKey,proto3" json:"project_key,omitempty"`
	Version      string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Registration string `protobuf:"bytes,4,opt,name=registration,proto3" json:"registration,omitempty"`
}

func (x *ProtoMeta) Reset() {
	*x = ProtoMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_common_image_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoMeta) ProtoMessage() {}

func (x *ProtoMeta) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_common_image_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoMeta.ProtoReflect.Descriptor instead.
func (*ProtoMeta) Descriptor() ([]byte, []int) {
	return file_api_v1_common_image_proto_rawDescGZIP(), []int{1}
}

func (x *ProtoMeta) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProtoMeta) GetProjectKey() string {
	if x != nil {
		return x.ProjectKey
	}
	return ""
}

func (x *ProtoMeta) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ProtoMeta) GetRegistration() string {
	if x != nil {
		return x.Registration
	}
	return ""
}

var File_api_v1_common_image_proto protoreflect.FileDescriptor

var file_api_v1_common_image_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x73,
	0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x22, 0x7a, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4d, 0x65, 0x74, 0x61,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4b, 0x65,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0xa5, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x0a, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0xa2, 0x02, 0x03, 0x4d, 0x56, 0x43, 0xaa, 0x02, 0x11, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xca, 0x02, 0x11,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0xe2, 0x02, 0x1d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x13, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a,
	0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_common_image_proto_rawDescOnce sync.Once
	file_api_v1_common_image_proto_rawDescData = file_api_v1_common_image_proto_rawDesc
)

func file_api_v1_common_image_proto_rawDescGZIP() []byte {
	file_api_v1_common_image_proto_rawDescOnce.Do(func() {
		file_api_v1_common_image_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_common_image_proto_rawDescData)
	})
	return file_api_v1_common_image_proto_rawDescData
}

var file_api_v1_common_image_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_v1_common_image_proto_goTypes = []interface{}{
	(*RegisteredProto)(nil), // 0: members.v1.common.RegisteredProto
	(*ProtoMeta)(nil),       // 1: members.v1.common.ProtoMeta
}
var file_api_v1_common_image_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1_common_image_proto_init() }
func file_api_v1_common_image_proto_init() {
	if File_api_v1_common_image_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_common_image_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisteredProto); i {
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
		file_api_v1_common_image_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoMeta); i {
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
			RawDescriptor: file_api_v1_common_image_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1_common_image_proto_goTypes,
		DependencyIndexes: file_api_v1_common_image_proto_depIdxs,
		MessageInfos:      file_api_v1_common_image_proto_msgTypes,
	}.Build()
	File_api_v1_common_image_proto = out.File
	file_api_v1_common_image_proto_rawDesc = nil
	file_api_v1_common_image_proto_goTypes = nil
	file_api_v1_common_image_proto_depIdxs = nil
}
