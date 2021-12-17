// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: departure.proto

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

type Departure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DisplayInformations *Display   `protobuf:"bytes,1,opt,name=display_informations,json=displayInformations,proto3" json:"display_informations,omitempty"`
	StopPoint           *StopPoint `protobuf:"bytes,2,opt,name=stop_point,json=stopPoint,proto3" json:"stop_point,omitempty"`
	Route               *Route     `protobuf:"bytes,3,opt,name=route,proto3" json:"route,omitempty"`
	Links               []*Link    `protobuf:"bytes,4,rep,name=links,proto3" json:"links,omitempty"`
}

func (x *Departure) Reset() {
	*x = Departure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_departure_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Departure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Departure) ProtoMessage() {}

func (x *Departure) ProtoReflect() protoreflect.Message {
	mi := &file_departure_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Departure.ProtoReflect.Descriptor instead.
func (*Departure) Descriptor() ([]byte, []int) {
	return file_departure_proto_rawDescGZIP(), []int{0}
}

func (x *Departure) GetDisplayInformations() *Display {
	if x != nil {
		return x.DisplayInformations
	}
	return nil
}

func (x *Departure) GetStopPoint() *StopPoint {
	if x != nil {
		return x.StopPoint
	}
	return nil
}

func (x *Departure) GetRoute() *Route {
	if x != nil {
		return x.Route
	}
	return nil
}

func (x *Departure) GetLinks() []*Link {
	if x != nil {
		return x.Links
	}
	return nil
}

var File_departure_proto protoreflect.FileDescriptor

var file_departure_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x72, 0x70, 0x63, 0x1a, 0x0d, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0a, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x01, 0x0a, 0x09,
	0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x75, 0x72, 0x65, 0x12, 0x3f, 0x0a, 0x14, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x52, 0x13, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x49, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2d, 0x0a, 0x0a, 0x73, 0x74,
	0x6f, 0x70, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x09,
	0x73, 0x74, 0x6f, 0x70, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x05, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x05, 0x6c,
	0x69, 0x6e, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x42, 0x20, 0x5a, 0x1e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x76, 0x69, 0x74,
	0x69, 0x61, 0x2f, 0x6e, 0x61, 0x76, 0x69, 0x74, 0x69, 0x61, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_departure_proto_rawDescOnce sync.Once
	file_departure_proto_rawDescData = file_departure_proto_rawDesc
)

func file_departure_proto_rawDescGZIP() []byte {
	file_departure_proto_rawDescOnce.Do(func() {
		file_departure_proto_rawDescData = protoimpl.X.CompressGZIP(file_departure_proto_rawDescData)
	})
	return file_departure_proto_rawDescData
}

var file_departure_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_departure_proto_goTypes = []interface{}{
	(*Departure)(nil), // 0: rpc.Departure
	(*Display)(nil),   // 1: rpc.Display
	(*StopPoint)(nil), // 2: rpc.StopPoint
	(*Route)(nil),     // 3: rpc.Route
	(*Link)(nil),      // 4: rpc.Link
}
var file_departure_proto_depIdxs = []int32{
	1, // 0: rpc.Departure.display_informations:type_name -> rpc.Display
	2, // 1: rpc.Departure.stop_point:type_name -> rpc.StopPoint
	3, // 2: rpc.Departure.route:type_name -> rpc.Route
	4, // 3: rpc.Departure.links:type_name -> rpc.Link
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_departure_proto_init() }
func file_departure_proto_init() {
	if File_departure_proto != nil {
		return
	}
	file_display_proto_init()
	file_place_proto_init()
	file_route_proto_init()
	file_link_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_departure_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Departure); i {
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
			RawDescriptor: file_departure_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_departure_proto_goTypes,
		DependencyIndexes: file_departure_proto_depIdxs,
		MessageInfos:      file_departure_proto_msgTypes,
	}.Build()
	File_departure_proto = out.File
	file_departure_proto_rawDesc = nil
	file_departure_proto_goTypes = nil
	file_departure_proto_depIdxs = nil
}