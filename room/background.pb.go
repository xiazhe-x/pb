// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: background.proto

package room

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	pb "voice/resource/pb"
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

type RoomBgArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RoomBgArgs) Reset() {
	*x = RoomBgArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_background_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomBgArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomBgArgs) ProtoMessage() {}

func (x *RoomBgArgs) ProtoReflect() protoreflect.Message {
	mi := &file_background_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomBgArgs.ProtoReflect.Descriptor instead.
func (*RoomBgArgs) Descriptor() ([]byte, []int) {
	return file_background_proto_rawDescGZIP(), []int{0}
}

func (x *RoomBgArgs) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type RoomBg struct {
	state         protoimpl.MessageState `xorm:"-"`
	sizeCache     protoimpl.SizeCache `xorm:"-"`
	unknownFields protoimpl.UnknownFields `xorm:"-"`

	Id      int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                          //id
	Picture string `protobuf:"bytes,2,opt,name=picture,proto3" json:"picture,omitempty"`                 //图片地址
	AddTime int64  `protobuf:"varint,3,opt,name=add_time,json=addTime,proto3" json:"add_time,omitempty"` //时间
	Name    string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`                       //名称
}

func (x *RoomBg) Reset() {
	*x = RoomBg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_background_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomBg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomBg) ProtoMessage() {}

func (x *RoomBg) ProtoReflect() protoreflect.Message {
	mi := &file_background_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomBg.ProtoReflect.Descriptor instead.
func (*RoomBg) Descriptor() ([]byte, []int) {
	return file_background_proto_rawDescGZIP(), []int{1}
}

func (x *RoomBg) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RoomBg) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

func (x *RoomBg) GetAddTime() int64 {
	if x != nil {
		return x.AddTime
	}
	return 0
}

func (x *RoomBg) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type BackRoomBgList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  []*RoomBg `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Limit int64     `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Count int64     `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *BackRoomBgList) Reset() {
	*x = BackRoomBgList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_background_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackRoomBgList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackRoomBgList) ProtoMessage() {}

func (x *BackRoomBgList) ProtoReflect() protoreflect.Message {
	mi := &file_background_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackRoomBgList.ProtoReflect.Descriptor instead.
func (*BackRoomBgList) Descriptor() ([]byte, []int) {
	return file_background_proto_rawDescGZIP(), []int{2}
}

func (x *BackRoomBgList) GetData() []*RoomBg {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *BackRoomBgList) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *BackRoomBgList) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_background_proto protoreflect.FileDescriptor

var file_background_proto_rawDesc = []byte{
	0x0a, 0x10, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x1c, 0x0a, 0x0a, 0x52, 0x6f, 0x6f, 0x6d, 0x42, 0x67, 0x41, 0x72, 0x67, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x61,
	0x0a, 0x06, 0x52, 0x6f, 0x6f, 0x6d, 0x42, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x69, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x64, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x64, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x59, 0x0a, 0x0e, 0x42, 0x61, 0x63, 0x6b, 0x52, 0x6f, 0x6f, 0x6d, 0x42, 0x67, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x42, 0x67, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x87, 0x01, 0x0a,
	0x0c, 0x52, 0x6f, 0x6f, 0x6d, 0x42, 0x67, 0x53, 0x70, 0x69, 0x64, 0x65, 0x72, 0x12, 0x2a, 0x0a,
	0x0a, 0x52, 0x6f, 0x6f, 0x6d, 0x42, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0b, 0x2e, 0x52, 0x6f,
	0x6f, 0x6d, 0x42, 0x67, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x0f, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x52,
	0x6f, 0x6f, 0x6d, 0x42, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0a, 0x52, 0x6f, 0x6f,
	0x6d, 0x42, 0x67, 0x45, 0x64, 0x69, 0x74, 0x12, 0x07, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x42, 0x67,
	0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x52, 0x6f, 0x6f, 0x6d, 0x42, 0x67, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x0b, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x42, 0x67, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x07, 0x2e,
	0x52, 0x6f, 0x6f, 0x6d, 0x42, 0x67, 0x42, 0x18, 0x5a, 0x16, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x6f, 0x6f, 0x6d,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_background_proto_rawDescOnce sync.Once
	file_background_proto_rawDescData = file_background_proto_rawDesc
)

func file_background_proto_rawDescGZIP() []byte {
	file_background_proto_rawDescOnce.Do(func() {
		file_background_proto_rawDescData = protoimpl.X.CompressGZIP(file_background_proto_rawDescData)
	})
	return file_background_proto_rawDescData
}

var file_background_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_background_proto_goTypes = []interface{}{
	(*RoomBgArgs)(nil),     // 0: RoomBgArgs
	(*RoomBg)(nil),         // 1: RoomBg
	(*BackRoomBgList)(nil), // 2: BackRoomBgList
	(*pb.Response)(nil),    // 3: common.Response
}
var file_background_proto_depIdxs = []int32{
	1, // 0: BackRoomBgList.data:type_name -> RoomBg
	0, // 1: RoomBgSpider.RoomBgList:input_type -> RoomBgArgs
	1, // 2: RoomBgSpider.RoomBgEdit:input_type -> RoomBg
	0, // 3: RoomBgSpider.RoomBgInfo:input_type -> RoomBgArgs
	2, // 4: RoomBgSpider.RoomBgList:output_type -> BackRoomBgList
	3, // 5: RoomBgSpider.RoomBgEdit:output_type -> common.Response
	1, // 6: RoomBgSpider.RoomBgInfo:output_type -> RoomBg
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_background_proto_init() }
func file_background_proto_init() {
	if File_background_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_background_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomBgArgs); i {
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
		file_background_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomBg); i {
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
		file_background_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackRoomBgList); i {
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
			RawDescriptor: file_background_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_background_proto_goTypes,
		DependencyIndexes: file_background_proto_depIdxs,
		MessageInfos:      file_background_proto_msgTypes,
	}.Build()
	File_background_proto = out.File
	file_background_proto_rawDesc = nil
	file_background_proto_goTypes = nil
	file_background_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RoomBgSpiderClient is the client API for RoomBgSpider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RoomBgSpiderClient interface {
	RoomBgList(ctx context.Context, in *RoomBgArgs, opts ...grpc.CallOption) (*BackRoomBgList, error)
	RoomBgEdit(ctx context.Context, in *RoomBg, opts ...grpc.CallOption) (*pb.Response, error)
	RoomBgInfo(ctx context.Context, in *RoomBgArgs, opts ...grpc.CallOption) (*RoomBg, error)
}

type roomBgSpiderClient struct {
	cc grpc.ClientConnInterface
}

func NewRoomBgSpiderClient(cc grpc.ClientConnInterface) RoomBgSpiderClient {
	return &roomBgSpiderClient{cc}
}

func (c *roomBgSpiderClient) RoomBgList(ctx context.Context, in *RoomBgArgs, opts ...grpc.CallOption) (*BackRoomBgList, error) {
	out := new(BackRoomBgList)
	err := c.cc.Invoke(ctx, "/RoomBgSpider/RoomBgList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomBgSpiderClient) RoomBgEdit(ctx context.Context, in *RoomBg, opts ...grpc.CallOption) (*pb.Response, error) {
	out := new(pb.Response)
	err := c.cc.Invoke(ctx, "/RoomBgSpider/RoomBgEdit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomBgSpiderClient) RoomBgInfo(ctx context.Context, in *RoomBgArgs, opts ...grpc.CallOption) (*RoomBg, error) {
	out := new(RoomBg)
	err := c.cc.Invoke(ctx, "/RoomBgSpider/RoomBgInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoomBgSpiderServer is the server API for RoomBgSpider service.
type RoomBgSpiderServer interface {
	RoomBgList(context.Context, *RoomBgArgs) (*BackRoomBgList, error)
	RoomBgEdit(context.Context, *RoomBg) (*pb.Response, error)
	RoomBgInfo(context.Context, *RoomBgArgs) (*RoomBg, error)
}

// UnimplementedRoomBgSpiderServer can be embedded to have forward compatible implementations.
type UnimplementedRoomBgSpiderServer struct {
}

func (*UnimplementedRoomBgSpiderServer) RoomBgList(context.Context, *RoomBgArgs) (*BackRoomBgList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoomBgList not implemented")
}
func (*UnimplementedRoomBgSpiderServer) RoomBgEdit(context.Context, *RoomBg) (*pb.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoomBgEdit not implemented")
}
func (*UnimplementedRoomBgSpiderServer) RoomBgInfo(context.Context, *RoomBgArgs) (*RoomBg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoomBgInfo not implemented")
}

func RegisterRoomBgSpiderServer(s *grpc.Server, srv RoomBgSpiderServer) {
	s.RegisterService(&_RoomBgSpider_serviceDesc, srv)
}

func _RoomBgSpider_RoomBgList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomBgArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomBgSpiderServer).RoomBgList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RoomBgSpider/RoomBgList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomBgSpiderServer).RoomBgList(ctx, req.(*RoomBgArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomBgSpider_RoomBgEdit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomBg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomBgSpiderServer).RoomBgEdit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RoomBgSpider/RoomBgEdit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomBgSpiderServer).RoomBgEdit(ctx, req.(*RoomBg))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomBgSpider_RoomBgInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomBgArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomBgSpiderServer).RoomBgInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RoomBgSpider/RoomBgInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomBgSpiderServer).RoomBgInfo(ctx, req.(*RoomBgArgs))
	}
	return interceptor(ctx, in, info, handler)
}

var _RoomBgSpider_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RoomBgSpider",
	HandlerType: (*RoomBgSpiderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RoomBgList",
			Handler:    _RoomBgSpider_RoomBgList_Handler,
		},
		{
			MethodName: "RoomBgEdit",
			Handler:    _RoomBgSpider_RoomBgEdit_Handler,
		},
		{
			MethodName: "RoomBgInfo",
			Handler:    _RoomBgSpider_RoomBgInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "background.proto",
}