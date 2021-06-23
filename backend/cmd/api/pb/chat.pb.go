// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.3
// source: chat.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Chat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	SenderID  string `protobuf:"bytes,2,opt,name=senderID,proto3" json:"senderID,omitempty"`
	InboxHash string `protobuf:"bytes,3,opt,name=inboxHash,proto3" json:"inboxHash,omitempty"`
	Msg       string `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg,omitempty"`
	File      string `protobuf:"bytes,5,opt,name=file,proto3" json:"file,omitempty"`
	Meta      string `protobuf:"bytes,6,opt,name=meta,proto3" json:"meta,omitempty"`
	DeletedBy string `protobuf:"bytes,7,opt,name=deleted_by,json=deletedBy,proto3" json:"deleted_by,omitempty"`
}

func (x *Chat) Reset() {
	*x = Chat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chat) ProtoMessage() {}

func (x *Chat) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chat.ProtoReflect.Descriptor instead.
func (*Chat) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

func (x *Chat) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Chat) GetSenderID() string {
	if x != nil {
		return x.SenderID
	}
	return ""
}

func (x *Chat) GetInboxHash() string {
	if x != nil {
		return x.InboxHash
	}
	return ""
}

func (x *Chat) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Chat) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

func (x *Chat) GetMeta() string {
	if x != nil {
		return x.Meta
	}
	return ""
}

func (x *Chat) GetDeletedBy() string {
	if x != nil {
		return x.DeletedBy
	}
	return ""
}

type ListChatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The parent resource name, for example, "shelves/shelf1"
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The maximum number of items to return.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The next_page_token value returned from a previous List request, if any.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListChatsRequest) Reset() {
	*x = ListChatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListChatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChatsRequest) ProtoMessage() {}

func (x *ListChatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChatsRequest.ProtoReflect.Descriptor instead.
func (*ListChatsRequest) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{1}
}

func (x *ListChatsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListChatsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListChatsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListChats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chats []*Chat `protobuf:"bytes,1,rep,name=Chats,proto3" json:"Chats,omitempty"`
}

func (x *ListChats) Reset() {
	*x = ListChats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListChats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChats) ProtoMessage() {}

func (x *ListChats) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChats.ProtoReflect.Descriptor instead.
func (*ListChats) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{2}
}

func (x *ListChats) GetChats() []*Chat {
	if x != nil {
		return x.Chats
	}
	return nil
}

var File_chat_proto protoreflect.FileDescriptor

var file_chat_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x66, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x01, 0x0a, 0x04, 0x43,
	0x68, 0x61, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x48, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x48, 0x61, 0x73, 0x68, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12,
	0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66,
	0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x42, 0x79, 0x22, 0x66, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68,
	0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x28,
	0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x74, 0x73, 0x12, 0x1b, 0x0a, 0x05, 0x43,
	0x68, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x43, 0x68, 0x61,
	0x74, 0x52, 0x05, 0x43, 0x68, 0x61, 0x74, 0x73, 0x32, 0x71, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x43, 0x68,
	0x61, 0x74, 0x73, 0x12, 0x11, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61,
	0x74, 0x73, 0x30, 0x01, 0x12, 0x1b, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68,
	0x61, 0x74, 0x12, 0x05, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x43, 0x68, 0x61, 0x74, 0x12, 0x05, 0x2e, 0x43,
	0x68, 0x61, 0x74, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x06, 0x5a, 0x04, 0x2e,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_proto_rawDescOnce sync.Once
	file_chat_proto_rawDescData = file_chat_proto_rawDesc
)

func file_chat_proto_rawDescGZIP() []byte {
	file_chat_proto_rawDescOnce.Do(func() {
		file_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_proto_rawDescData)
	})
	return file_chat_proto_rawDescData
}

var file_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_chat_proto_goTypes = []interface{}{
	(*Chat)(nil),             // 0: Chat
	(*ListChatsRequest)(nil), // 1: ListChatsRequest
	(*ListChats)(nil),        // 2: ListChats
	(*Empty)(nil),            // 3: Empty
}
var file_chat_proto_depIdxs = []int32{
	0, // 0: ListChats.Chats:type_name -> Chat
	1, // 1: ChatService.GetChats:input_type -> ListChatsRequest
	0, // 2: ChatService.DeleteChat:input_type -> Chat
	0, // 3: ChatService.AddChat:input_type -> Chat
	2, // 4: ChatService.GetChats:output_type -> ListChats
	3, // 5: ChatService.DeleteChat:output_type -> Empty
	3, // 6: ChatService.AddChat:output_type -> Empty
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_chat_proto_init() }
func file_chat_proto_init() {
	if File_chat_proto != nil {
		return
	}
	file_friend_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chat); i {
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
		file_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListChatsRequest); i {
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
		file_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListChats); i {
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
			RawDescriptor: file_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_proto_goTypes,
		DependencyIndexes: file_chat_proto_depIdxs,
		MessageInfos:      file_chat_proto_msgTypes,
	}.Build()
	File_chat_proto = out.File
	file_chat_proto_rawDesc = nil
	file_chat_proto_goTypes = nil
	file_chat_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatServiceClient interface {
	GetChats(ctx context.Context, in *ListChatsRequest, opts ...grpc.CallOption) (ChatService_GetChatsClient, error)
	DeleteChat(ctx context.Context, in *Chat, opts ...grpc.CallOption) (*Empty, error)
	AddChat(ctx context.Context, in *Chat, opts ...grpc.CallOption) (*Empty, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) GetChats(ctx context.Context, in *ListChatsRequest, opts ...grpc.CallOption) (ChatService_GetChatsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChatService_serviceDesc.Streams[0], "/ChatService/GetChats", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceGetChatsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatService_GetChatsClient interface {
	Recv() (*ListChats, error)
	grpc.ClientStream
}

type chatServiceGetChatsClient struct {
	grpc.ClientStream
}

func (x *chatServiceGetChatsClient) Recv() (*ListChats, error) {
	m := new(ListChats)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) DeleteChat(ctx context.Context, in *Chat, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/ChatService/DeleteChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) AddChat(ctx context.Context, in *Chat, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/ChatService/AddChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
type ChatServiceServer interface {
	GetChats(*ListChatsRequest, ChatService_GetChatsServer) error
	DeleteChat(context.Context, *Chat) (*Empty, error)
	AddChat(context.Context, *Chat) (*Empty, error)
}

// UnimplementedChatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (*UnimplementedChatServiceServer) GetChats(*ListChatsRequest, ChatService_GetChatsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetChats not implemented")
}
func (*UnimplementedChatServiceServer) DeleteChat(context.Context, *Chat) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChat not implemented")
}
func (*UnimplementedChatServiceServer) AddChat(context.Context, *Chat) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddChat not implemented")
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_GetChats_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListChatsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).GetChats(m, &chatServiceGetChatsServer{stream})
}

type ChatService_GetChatsServer interface {
	Send(*ListChats) error
	grpc.ServerStream
}

type chatServiceGetChatsServer struct {
	grpc.ServerStream
}

func (x *chatServiceGetChatsServer) Send(m *ListChats) error {
	return x.ServerStream.SendMsg(m)
}

func _ChatService_DeleteChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Chat)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).DeleteChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatService/DeleteChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).DeleteChat(ctx, req.(*Chat))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_AddChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Chat)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).AddChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatService/AddChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).AddChat(ctx, req.(*Chat))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteChat",
			Handler:    _ChatService_DeleteChat_Handler,
		},
		{
			MethodName: "AddChat",
			Handler:    _ChatService_AddChat_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetChats",
			Handler:       _ChatService_GetChats_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chat.proto",
}
