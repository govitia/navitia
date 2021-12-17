// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: coord.proto

package rpc

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Coordinates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat float64 `protobuf:"fixed64,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lon float64 `protobuf:"fixed64,2,opt,name=lon,proto3" json:"lon,omitempty"`
}

func (x *Coordinates) Reset() {
	*x = Coordinates{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coord_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coordinates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coordinates) ProtoMessage() {}

func (x *Coordinates) ProtoReflect() protoreflect.Message {
	mi := &file_coord_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coordinates.ProtoReflect.Descriptor instead.
func (*Coordinates) Descriptor() ([]byte, []int) {
	return file_coord_proto_rawDescGZIP(), []int{0}
}

func (x *Coordinates) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *Coordinates) GetLon() float64 {
	if x != nil {
		return x.Lon
	}
	return 0
}

var File_coord_proto protoreflect.FileDescriptor

var file_coord_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72,
	0x70, 0x63, 0x22, 0x31, 0x0a, 0x0b, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65,
	0x73, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03,
	0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x03, 0x6c, 0x6f, 0x6e, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x76, 0x69, 0x74, 0x69, 0x61, 0x2f, 0x6e, 0x61, 0x76, 0x69,
	0x74, 0x69, 0x61, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coord_proto_rawDescOnce sync.Once
	file_coord_proto_rawDescData = file_coord_proto_rawDesc
)

func file_coord_proto_rawDescGZIP() []byte {
	file_coord_proto_rawDescOnce.Do(func() {
		file_coord_proto_rawDescData = protoimpl.X.CompressGZIP(file_coord_proto_rawDescData)
	})
	return file_coord_proto_rawDescData
}

var file_coord_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_coord_proto_goTypes = []interface{}{
	(*Coordinates)(nil), // 0: rpc.Coordinates
}
var file_coord_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_coord_proto_init() }
func file_coord_proto_init() {
	if File_coord_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coord_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coordinates); i {
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
			RawDescriptor: file_coord_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_coord_proto_goTypes,
		DependencyIndexes: file_coord_proto_depIdxs,
		MessageInfos:      file_coord_proto_msgTypes,
	}.Build()
	File_coord_proto = out.File
	file_coord_proto_rawDesc = nil
	file_coord_proto_goTypes = nil
	file_coord_proto_depIdxs = nil
}
