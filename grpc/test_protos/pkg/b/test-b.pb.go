// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: test_protos/pkg/b/test-b.proto

package b

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

type EnumAllowingAlias int32

const (
	EnumAllowingAlias_UNKNOWN EnumAllowingAlias = 0
	EnumAllowingAlias_STARTED EnumAllowingAlias = 1
	EnumAllowingAlias_RUNNING EnumAllowingAlias = 2
)

// Enum value maps for EnumAllowingAlias.
var (
	EnumAllowingAlias_name = map[int32]string{
		0: "UNKNOWN",
		1: "STARTED",
		2: "RUNNING",
	}
	EnumAllowingAlias_value = map[string]int32{
		"UNKNOWN": 0,
		"STARTED": 1,
		"RUNNING": 2,
	}
)

func (x EnumAllowingAlias) Enum() *EnumAllowingAlias {
	p := new(EnumAllowingAlias)
	*p = x
	return p
}

func (x EnumAllowingAlias) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumAllowingAlias) Descriptor() protoreflect.EnumDescriptor {
	return file_test_protos_pkg_b_test_b_proto_enumTypes[0].Descriptor()
}

func (EnumAllowingAlias) Type() protoreflect.EnumType {
	return &file_test_protos_pkg_b_test_b_proto_enumTypes[0]
}

func (x EnumAllowingAlias) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnumAllowingAlias.Descriptor instead.
func (EnumAllowingAlias) EnumDescriptor() ([]byte, []int) {
	return file_test_protos_pkg_b_test_b_proto_rawDescGZIP(), []int{0}
}

type Outer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// some comment added
	InnerMessage []*OuterInner     `protobuf:"bytes,2,rep,name=inner_message,json=innerMessage,proto3" json:"inner_message,omitempty"`
	EnumField    EnumAllowingAlias `protobuf:"varint,3,opt,name=enum_field,json=enumField,proto3,enum=pbtestb.EnumAllowingAlias" json:"enum_field,omitempty"`
	MyMap        map[int32]string  `protobuf:"bytes,4,rep,name=my_map,json=myMap,proto3" json:"my_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Outer) Reset() {
	*x = Outer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_protos_pkg_b_test_b_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Outer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Outer) ProtoMessage() {}

func (x *Outer) ProtoReflect() protoreflect.Message {
	mi := &file_test_protos_pkg_b_test_b_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Outer.ProtoReflect.Descriptor instead.
func (*Outer) Descriptor() ([]byte, []int) {
	return file_test_protos_pkg_b_test_b_proto_rawDescGZIP(), []int{0}
}

func (x *Outer) GetInnerMessage() []*OuterInner {
	if x != nil {
		return x.InnerMessage
	}
	return nil
}

func (x *Outer) GetEnumField() EnumAllowingAlias {
	if x != nil {
		return x.EnumField
	}
	return EnumAllowingAlias_UNKNOWN
}

func (x *Outer) GetMyMap() map[int32]string {
	if x != nil {
		return x.MyMap
	}
	return nil
}

type OuterInner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ival int64 `protobuf:"varint,1,opt,name=ival,proto3" json:"ival,omitempty"`
}

func (x *OuterInner) Reset() {
	*x = OuterInner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_protos_pkg_b_test_b_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OuterInner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OuterInner) ProtoMessage() {}

func (x *OuterInner) ProtoReflect() protoreflect.Message {
	mi := &file_test_protos_pkg_b_test_b_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OuterInner.ProtoReflect.Descriptor instead.
func (*OuterInner) Descriptor() ([]byte, []int) {
	return file_test_protos_pkg_b_test_b_proto_rawDescGZIP(), []int{0, 0}
}

func (x *OuterInner) GetIval() int64 {
	if x != nil {
		return x.Ival
	}
	return 0
}

var File_test_protos_pkg_b_test_b_proto protoreflect.FileDescriptor

var file_test_protos_pkg_b_test_b_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x62, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x70, 0x62, 0x74, 0x65, 0x73, 0x74, 0x62, 0x22, 0x86, 0x02, 0x0a, 0x05, 0x6f, 0x75,
	0x74, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x0d, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x62, 0x74,
	0x65, 0x73, 0x74, 0x62, 0x2e, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x69, 0x6e, 0x6e, 0x65, 0x72,
	0x52, 0x0c, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x39,
	0x0a, 0x0a, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x70, 0x62, 0x74, 0x65, 0x73, 0x74, 0x62, 0x2e, 0x45, 0x6e, 0x75,
	0x6d, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x52, 0x09,
	0x65, 0x6e, 0x75, 0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x30, 0x0a, 0x06, 0x6d, 0x79, 0x5f,
	0x6d, 0x61, 0x70, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x62, 0x74, 0x65,
	0x73, 0x74, 0x62, 0x2e, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x4d, 0x79, 0x4d, 0x61, 0x70, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x6d, 0x79, 0x4d, 0x61, 0x70, 0x1a, 0x1b, 0x0a, 0x05, 0x69,
	0x6e, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x69, 0x76, 0x61, 0x6c, 0x1a, 0x38, 0x0a, 0x0a, 0x4d, 0x79, 0x4d, 0x61,
	0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x2a, 0x3a, 0x0a, 0x11, 0x45, 0x6e, 0x75, 0x6d, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x69,
	0x6e, 0x67, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x32, 0x42,
	0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32,
	0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x74,
	0x65, 0x73, 0x74, 0x62, 0x2e, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x74,
	0x65, 0x73, 0x74, 0x62, 0x2e, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x69, 0x6e, 0x6e, 0x65, 0x72,
	0x22, 0x00, 0x42, 0x75, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x62, 0x74, 0x65, 0x73, 0x74,
	0x62, 0x42, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x42, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x1e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x62, 0xa2,
	0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x50, 0x62, 0x74, 0x65, 0x73, 0x74, 0x62, 0xca,
	0x02, 0x07, 0x50, 0x62, 0x74, 0x65, 0x73, 0x74, 0x62, 0xe2, 0x02, 0x13, 0x50, 0x62, 0x74, 0x65,
	0x73, 0x74, 0x62, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x07, 0x50, 0x62, 0x74, 0x65, 0x73, 0x74, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_test_protos_pkg_b_test_b_proto_rawDescOnce sync.Once
	file_test_protos_pkg_b_test_b_proto_rawDescData = file_test_protos_pkg_b_test_b_proto_rawDesc
)

func file_test_protos_pkg_b_test_b_proto_rawDescGZIP() []byte {
	file_test_protos_pkg_b_test_b_proto_rawDescOnce.Do(func() {
		file_test_protos_pkg_b_test_b_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_protos_pkg_b_test_b_proto_rawDescData)
	})
	return file_test_protos_pkg_b_test_b_proto_rawDescData
}

var file_test_protos_pkg_b_test_b_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_test_protos_pkg_b_test_b_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_test_protos_pkg_b_test_b_proto_goTypes = []interface{}{
	(EnumAllowingAlias)(0), // 0: pbtestb.EnumAllowingAlias
	(*Outer)(nil),          // 1: pbtestb.outer
	(*OuterInner)(nil),     // 2: pbtestb.outer.inner
	nil,                    // 3: pbtestb.outer.MyMapEntry
}
var file_test_protos_pkg_b_test_b_proto_depIdxs = []int32{
	2, // 0: pbtestb.outer.inner_message:type_name -> pbtestb.outer.inner
	0, // 1: pbtestb.outer.enum_field:type_name -> pbtestb.EnumAllowingAlias
	3, // 2: pbtestb.outer.my_map:type_name -> pbtestb.outer.MyMapEntry
	1, // 3: pbtestb.HelloService.SayHello:input_type -> pbtestb.outer
	2, // 4: pbtestb.HelloService.SayHello:output_type -> pbtestb.outer.inner
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_test_protos_pkg_b_test_b_proto_init() }
func file_test_protos_pkg_b_test_b_proto_init() {
	if File_test_protos_pkg_b_test_b_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_protos_pkg_b_test_b_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Outer); i {
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
		file_test_protos_pkg_b_test_b_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OuterInner); i {
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
			RawDescriptor: file_test_protos_pkg_b_test_b_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_test_protos_pkg_b_test_b_proto_goTypes,
		DependencyIndexes: file_test_protos_pkg_b_test_b_proto_depIdxs,
		EnumInfos:         file_test_protos_pkg_b_test_b_proto_enumTypes,
		MessageInfos:      file_test_protos_pkg_b_test_b_proto_msgTypes,
	}.Build()
	File_test_protos_pkg_b_test_b_proto = out.File
	file_test_protos_pkg_b_test_b_proto_rawDesc = nil
	file_test_protos_pkg_b_test_b_proto_goTypes = nil
	file_test_protos_pkg_b_test_b_proto_depIdxs = nil
}
