// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: api/v1/common/def.proto

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

type Service int32

const (
	Service_SERVICE_UNKNOWN  Service = 0
	Service_SERVICE_HEALTH   Service = 1
	Service_SERVICE_REGISTRY Service = 2
	Service_SERVICE_ADMIN    Service = 99
)

// Enum value maps for Service.
var (
	Service_name = map[int32]string{
		0:  "SERVICE_UNKNOWN",
		1:  "SERVICE_HEALTH",
		2:  "SERVICE_REGISTRY",
		99: "SERVICE_ADMIN",
	}
	Service_value = map[string]int32{
		"SERVICE_UNKNOWN":  0,
		"SERVICE_HEALTH":   1,
		"SERVICE_REGISTRY": 2,
		"SERVICE_ADMIN":    99,
	}
)

func (x Service) Enum() *Service {
	p := new(Service)
	*p = x
	return p
}

func (x Service) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Service) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1_common_def_proto_enumTypes[0].Descriptor()
}

func (Service) Type() protoreflect.EnumType {
	return &file_api_v1_common_def_proto_enumTypes[0]
}

func (x Service) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Service.Descriptor instead.
func (Service) EnumDescriptor() ([]byte, []int) {
	return file_api_v1_common_def_proto_rawDescGZIP(), []int{0}
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_common_def_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_common_def_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_api_v1_common_def_proto_rawDescGZIP(), []int{0}
}

type Member struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address    string  `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Dns        string  `protobuf:"bytes,2,opt,name=dns,proto3" json:"dns,omitempty"`
	Service    Service `protobuf:"varint,3,opt,name=service,proto3,enum=members.v1.common.Service" json:"service,omitempty"`
	JoinTime   string  `protobuf:"bytes,4,opt,name=join_time,json=joinTime,proto3" json:"join_time,omitempty"`
	LastHealth string  `protobuf:"bytes,5,opt,name=last_health,json=lastHealth,proto3" json:"last_health,omitempty"`
}

func (x *Member) Reset() {
	*x = Member{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_common_def_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Member) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Member) ProtoMessage() {}

func (x *Member) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_common_def_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Member.ProtoReflect.Descriptor instead.
func (*Member) Descriptor() ([]byte, []int) {
	return file_api_v1_common_def_proto_rawDescGZIP(), []int{1}
}

func (x *Member) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Member) GetDns() string {
	if x != nil {
		return x.Dns
	}
	return ""
}

func (x *Member) GetService() Service {
	if x != nil {
		return x.Service
	}
	return Service_SERVICE_UNKNOWN
}

func (x *Member) GetJoinTime() string {
	if x != nil {
		return x.JoinTime
	}
	return ""
}

func (x *Member) GetLastHealth() string {
	if x != nil {
		return x.LastHealth
	}
	return ""
}

var File_api_v1_common_def_proto protoreflect.FileDescriptor

var file_api_v1_common_def_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x64, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x07, 0x0a, 0x05,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0xa8, 0x01, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x6e,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x6e, 0x73, 0x12, 0x34, 0x0a, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6a, 0x6f, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6a, 0x6f, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x2a, 0x5b, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x53,
	0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x12, 0x0a, 0x0e, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x48, 0x45, 0x41, 0x4c,
	0x54, 0x48, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f,
	0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x52, 0x59, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x45,
	0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x63, 0x42, 0xa3, 0x01,
	0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x08, 0x44, 0x65, 0x66, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x1a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xa2,
	0x02, 0x03, 0x4d, 0x56, 0x43, 0xaa, 0x02, 0x11, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e,
	0x56, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xca, 0x02, 0x11, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xe2, 0x02, 0x1d,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_common_def_proto_rawDescOnce sync.Once
	file_api_v1_common_def_proto_rawDescData = file_api_v1_common_def_proto_rawDesc
)

func file_api_v1_common_def_proto_rawDescGZIP() []byte {
	file_api_v1_common_def_proto_rawDescOnce.Do(func() {
		file_api_v1_common_def_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_common_def_proto_rawDescData)
	})
	return file_api_v1_common_def_proto_rawDescData
}

var file_api_v1_common_def_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_v1_common_def_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_v1_common_def_proto_goTypes = []interface{}{
	(Service)(0),   // 0: members.v1.common.Service
	(*Empty)(nil),  // 1: members.v1.common.Empty
	(*Member)(nil), // 2: members.v1.common.Member
}
var file_api_v1_common_def_proto_depIdxs = []int32{
	0, // 0: members.v1.common.Member.service:type_name -> members.v1.common.Service
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_v1_common_def_proto_init() }
func file_api_v1_common_def_proto_init() {
	if File_api_v1_common_def_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_common_def_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_api_v1_common_def_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Member); i {
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
			RawDescriptor: file_api_v1_common_def_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1_common_def_proto_goTypes,
		DependencyIndexes: file_api_v1_common_def_proto_depIdxs,
		EnumInfos:         file_api_v1_common_def_proto_enumTypes,
		MessageInfos:      file_api_v1_common_def_proto_msgTypes,
	}.Build()
	File_api_v1_common_def_proto = out.File
	file_api_v1_common_def_proto_rawDesc = nil
	file_api_v1_common_def_proto_goTypes = nil
	file_api_v1_common_def_proto_depIdxs = nil
}
