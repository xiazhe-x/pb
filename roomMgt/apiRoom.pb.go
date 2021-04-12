// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: apiRoom.proto

package roomMgt

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

type ConnParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"` //token
	RoomId uint32 `protobuf:"varint,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	UserId uint32 `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ConnParams) Reset() {
	*x = ConnParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apiRoom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnParams) ProtoMessage() {}

func (x *ConnParams) ProtoReflect() protoreflect.Message {
	mi := &file_apiRoom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnParams.ProtoReflect.Descriptor instead.
func (*ConnParams) Descriptor() ([]byte, []int) {
	return file_apiRoom_proto_rawDescGZIP(), []int{0}
}

func (x *ConnParams) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ConnParams) GetRoomId() uint32 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *ConnParams) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type RoomUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid       uint32 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Nickname  string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	AvatarUrl string `protobuf:"bytes,3,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
}

func (x *RoomUser) Reset() {
	*x = RoomUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apiRoom_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomUser) ProtoMessage() {}

func (x *RoomUser) ProtoReflect() protoreflect.Message {
	mi := &file_apiRoom_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomUser.ProtoReflect.Descriptor instead.
func (*RoomUser) Descriptor() ([]byte, []int) {
	return file_apiRoom_proto_rawDescGZIP(), []int{1}
}

func (x *RoomUser) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *RoomUser) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *RoomUser) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

type ReqMikeMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MikeId uint32 `protobuf:"varint,1,opt,name=mike_id,json=mikeId,proto3" json:"mike_id,omitempty"` // 麦克id
	Genre  uint32 `protobuf:"varint,2,opt,name=genre,proto3" json:"genre,omitempty"`                 //操作类型  1上麦 2下麦 3随机上麦
}

func (x *ReqMikeMsg) Reset() {
	*x = ReqMikeMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apiRoom_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqMikeMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqMikeMsg) ProtoMessage() {}

func (x *ReqMikeMsg) ProtoReflect() protoreflect.Message {
	mi := &file_apiRoom_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqMikeMsg.ProtoReflect.Descriptor instead.
func (*ReqMikeMsg) Descriptor() ([]byte, []int) {
	return file_apiRoom_proto_rawDescGZIP(), []int{2}
}

func (x *ReqMikeMsg) GetMikeId() uint32 {
	if x != nil {
		return x.MikeId
	}
	return 0
}

func (x *ReqMikeMsg) GetGenre() uint32 {
	if x != nil {
		return x.Genre
	}
	return 0
}

var File_apiRoom_proto protoreflect.FileDescriptor

var file_apiRoom_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x70, 0x69, 0x52, 0x6f, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x54, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x57, 0x0a, 0x08, 0x52, 0x6f, 0x6f, 0x6d, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x22, 0x3b,
	0x0a, 0x0a, 0x52, 0x65, 0x71, 0x4d, 0x69, 0x6b, 0x65, 0x4d, 0x73, 0x67, 0x12, 0x17, 0x0a, 0x07,
	0x6d, 0x69, 0x6b, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6d,
	0x69, 0x6b, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x42, 0x20, 0x5a, 0x1e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x69, 0x61, 0x7a, 0x68, 0x65,
	0x2d, 0x78, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x6f, 0x6f, 0x6d, 0x4d, 0x67, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apiRoom_proto_rawDescOnce sync.Once
	file_apiRoom_proto_rawDescData = file_apiRoom_proto_rawDesc
)

func file_apiRoom_proto_rawDescGZIP() []byte {
	file_apiRoom_proto_rawDescOnce.Do(func() {
		file_apiRoom_proto_rawDescData = protoimpl.X.CompressGZIP(file_apiRoom_proto_rawDescData)
	})
	return file_apiRoom_proto_rawDescData
}

var file_apiRoom_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_apiRoom_proto_goTypes = []interface{}{
	(*ConnParams)(nil), // 0: ConnParams
	(*RoomUser)(nil),   // 1: RoomUser
	(*ReqMikeMsg)(nil), // 2: ReqMikeMsg
}
var file_apiRoom_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_apiRoom_proto_init() }
func file_apiRoom_proto_init() {
	if File_apiRoom_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apiRoom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnParams); i {
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
		file_apiRoom_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomUser); i {
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
		file_apiRoom_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqMikeMsg); i {
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
			RawDescriptor: file_apiRoom_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apiRoom_proto_goTypes,
		DependencyIndexes: file_apiRoom_proto_depIdxs,
		MessageInfos:      file_apiRoom_proto_msgTypes,
	}.Build()
	File_apiRoom_proto = out.File
	file_apiRoom_proto_rawDesc = nil
	file_apiRoom_proto_goTypes = nil
	file_apiRoom_proto_depIdxs = nil
}