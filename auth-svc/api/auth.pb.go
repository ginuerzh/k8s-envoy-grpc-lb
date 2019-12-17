// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreateTokenRequest struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Platform             string   `protobuf:"bytes,2,opt,name=platform,proto3" json:"platform,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTokenRequest) Reset()         { *m = CreateTokenRequest{} }
func (m *CreateTokenRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTokenRequest) ProtoMessage()    {}
func (*CreateTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *CreateTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTokenRequest.Unmarshal(m, b)
}
func (m *CreateTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTokenRequest.Marshal(b, m, deterministic)
}
func (m *CreateTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTokenRequest.Merge(m, src)
}
func (m *CreateTokenRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTokenRequest.Size(m)
}
func (m *CreateTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTokenRequest proto.InternalMessageInfo

func (m *CreateTokenRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *CreateTokenRequest) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

type CreateTokenResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Expiration           int64    `protobuf:"varint,2,opt,name=expiration,proto3" json:"expiration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTokenResponse) Reset()         { *m = CreateTokenResponse{} }
func (m *CreateTokenResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTokenResponse) ProtoMessage()    {}
func (*CreateTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *CreateTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTokenResponse.Unmarshal(m, b)
}
func (m *CreateTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTokenResponse.Marshal(b, m, deterministic)
}
func (m *CreateTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTokenResponse.Merge(m, src)
}
func (m *CreateTokenResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTokenResponse.Size(m)
}
func (m *CreateTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTokenResponse proto.InternalMessageInfo

func (m *CreateTokenResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CreateTokenResponse) GetExpiration() int64 {
	if m != nil {
		return m.Expiration
	}
	return 0
}

type DecodeTokenRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DecodeTokenRequest) Reset()         { *m = DecodeTokenRequest{} }
func (m *DecodeTokenRequest) String() string { return proto.CompactTextString(m) }
func (*DecodeTokenRequest) ProtoMessage()    {}
func (*DecodeTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{2}
}

func (m *DecodeTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecodeTokenRequest.Unmarshal(m, b)
}
func (m *DecodeTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecodeTokenRequest.Marshal(b, m, deterministic)
}
func (m *DecodeTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecodeTokenRequest.Merge(m, src)
}
func (m *DecodeTokenRequest) XXX_Size() int {
	return xxx_messageInfo_DecodeTokenRequest.Size(m)
}
func (m *DecodeTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DecodeTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DecodeTokenRequest proto.InternalMessageInfo

func (m *DecodeTokenRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type DecodeTokenResponse struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Platform             string   `protobuf:"bytes,2,opt,name=platform,proto3" json:"platform,omitempty"`
	Expiration           int64    `protobuf:"varint,3,opt,name=expiration,proto3" json:"expiration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DecodeTokenResponse) Reset()         { *m = DecodeTokenResponse{} }
func (m *DecodeTokenResponse) String() string { return proto.CompactTextString(m) }
func (*DecodeTokenResponse) ProtoMessage()    {}
func (*DecodeTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{3}
}

func (m *DecodeTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecodeTokenResponse.Unmarshal(m, b)
}
func (m *DecodeTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecodeTokenResponse.Marshal(b, m, deterministic)
}
func (m *DecodeTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecodeTokenResponse.Merge(m, src)
}
func (m *DecodeTokenResponse) XXX_Size() int {
	return xxx_messageInfo_DecodeTokenResponse.Size(m)
}
func (m *DecodeTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DecodeTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DecodeTokenResponse proto.InternalMessageInfo

func (m *DecodeTokenResponse) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *DecodeTokenResponse) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *DecodeTokenResponse) GetExpiration() int64 {
	if m != nil {
		return m.Expiration
	}
	return 0
}

type DeleteTokenRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTokenRequest) Reset()         { *m = DeleteTokenRequest{} }
func (m *DeleteTokenRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteTokenRequest) ProtoMessage()    {}
func (*DeleteTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{4}
}

func (m *DeleteTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTokenRequest.Unmarshal(m, b)
}
func (m *DeleteTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTokenRequest.Marshal(b, m, deterministic)
}
func (m *DeleteTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTokenRequest.Merge(m, src)
}
func (m *DeleteTokenRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteTokenRequest.Size(m)
}
func (m *DeleteTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTokenRequest proto.InternalMessageInfo

func (m *DeleteTokenRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type DeleteTokenResponse struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTokenResponse) Reset()         { *m = DeleteTokenResponse{} }
func (m *DeleteTokenResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteTokenResponse) ProtoMessage()    {}
func (*DeleteTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{5}
}

func (m *DeleteTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTokenResponse.Unmarshal(m, b)
}
func (m *DeleteTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTokenResponse.Marshal(b, m, deterministic)
}
func (m *DeleteTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTokenResponse.Merge(m, src)
}
func (m *DeleteTokenResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteTokenResponse.Size(m)
}
func (m *DeleteTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTokenResponse proto.InternalMessageInfo

func (m *DeleteTokenResponse) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateTokenRequest)(nil), "api.CreateTokenRequest")
	proto.RegisterType((*CreateTokenResponse)(nil), "api.CreateTokenResponse")
	proto.RegisterType((*DecodeTokenRequest)(nil), "api.DecodeTokenRequest")
	proto.RegisterType((*DecodeTokenResponse)(nil), "api.DecodeTokenResponse")
	proto.RegisterType((*DeleteTokenRequest)(nil), "api.DeleteTokenRequest")
	proto.RegisterType((*DeleteTokenResponse)(nil), "api.DeleteTokenResponse")
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874) }

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0x87, 0xa9, 0x51, 0xd1, 0xd9, 0x8b, 0xcc, 0x0a, 0x96, 0x3d, 0x88, 0xe4, 0xa2, 0x78, 0xe8,
	0x41, 0x5f, 0x60, 0xfd, 0x73, 0xf3, 0x56, 0x7c, 0x81, 0xb8, 0x3b, 0xb2, 0xc1, 0xb5, 0x89, 0xe9,
	0x04, 0x7c, 0x48, 0x1f, 0x4a, 0x1a, 0x43, 0x48, 0x37, 0x15, 0xf4, 0xd6, 0x99, 0x0e, 0x5f, 0xbe,
	0xf9, 0x25, 0x00, 0xca, 0xf3, 0xa6, 0xb1, 0xce, 0xb0, 0x41, 0xa1, 0xac, 0x96, 0xf7, 0x80, 0x0f,
	0x8e, 0x14, 0xd3, 0xb3, 0x79, 0xa3, 0xae, 0xa5, 0x0f, 0x4f, 0x3d, 0xe3, 0x09, 0x08, 0xaf, 0xd7,
	0x75, 0x75, 0x51, 0x5d, 0x1d, 0xb7, 0xc3, 0x27, 0x2e, 0xe0, 0xc8, 0x6e, 0x15, 0xbf, 0x1a, 0xf7,
	0x5e, 0xef, 0x85, 0x76, 0xaa, 0xe5, 0x13, 0xcc, 0x47, 0x8c, 0xde, 0x9a, 0xae, 0x27, 0x3c, 0x85,
	0x03, 0x1e, 0x1a, 0x11, 0xf3, 0x53, 0xe0, 0x39, 0x00, 0x7d, 0x5a, 0xed, 0x14, 0x6b, 0xd3, 0x05,
	0x94, 0x68, 0xb3, 0x8e, 0xbc, 0x06, 0x7c, 0xa4, 0x95, 0x59, 0x8f, 0x85, 0x26, 0x59, 0x72, 0x05,
	0xf3, 0xd1, 0x6c, 0x3c, 0xf8, 0x5f, 0xf6, 0x3b, 0x42, 0x62, 0x5a, 0x68, 0x4b, 0xfc, 0x17, 0xa1,
	0xcb, 0x41, 0x28, 0x9b, 0xfd, 0x4d, 0xe8, 0xe6, 0xab, 0x82, 0xfd, 0x3b, 0xcf, 0x1b, 0x5c, 0xc2,
	0x2c, 0xcb, 0x0e, 0xcf, 0x1a, 0x65, 0x75, 0x53, 0xde, 0xc8, 0xa2, 0x2e, 0x7f, 0x44, 0xf8, 0x12,
	0x66, 0x59, 0x08, 0x91, 0x50, 0x46, 0x18, 0x09, 0x53, 0x79, 0x05, 0x42, 0xb2, 0x4e, 0x84, 0xdd,
	0x9d, 0x13, 0xa1, 0x58, 0xf0, 0xe5, 0x30, 0xbc, 0xa8, 0xdb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x25, 0xe9, 0xac, 0x12, 0x5f, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthClient interface {
	CreateToken(ctx context.Context, in *CreateTokenRequest, opts ...grpc.CallOption) (*CreateTokenResponse, error)
	DecodeToken(ctx context.Context, in *DecodeTokenRequest, opts ...grpc.CallOption) (*DecodeTokenResponse, error)
	DeleteToken(ctx context.Context, in *DeleteTokenRequest, opts ...grpc.CallOption) (*DeleteTokenResponse, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) CreateToken(ctx context.Context, in *CreateTokenRequest, opts ...grpc.CallOption) (*CreateTokenResponse, error) {
	out := new(CreateTokenResponse)
	err := c.cc.Invoke(ctx, "/api.Auth/CreateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) DecodeToken(ctx context.Context, in *DecodeTokenRequest, opts ...grpc.CallOption) (*DecodeTokenResponse, error) {
	out := new(DecodeTokenResponse)
	err := c.cc.Invoke(ctx, "/api.Auth/DecodeToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) DeleteToken(ctx context.Context, in *DeleteTokenRequest, opts ...grpc.CallOption) (*DeleteTokenResponse, error) {
	out := new(DeleteTokenResponse)
	err := c.cc.Invoke(ctx, "/api.Auth/DeleteToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
type AuthServer interface {
	CreateToken(context.Context, *CreateTokenRequest) (*CreateTokenResponse, error)
	DecodeToken(context.Context, *DecodeTokenRequest) (*DecodeTokenResponse, error)
	DeleteToken(context.Context, *DeleteTokenRequest) (*DeleteTokenResponse, error)
}

// UnimplementedAuthServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (*UnimplementedAuthServer) CreateToken(ctx context.Context, req *CreateTokenRequest) (*CreateTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateToken not implemented")
}
func (*UnimplementedAuthServer) DecodeToken(ctx context.Context, req *DecodeTokenRequest) (*DecodeTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecodeToken not implemented")
}
func (*UnimplementedAuthServer) DeleteToken(ctx context.Context, req *DeleteTokenRequest) (*DeleteTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteToken not implemented")
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_CreateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CreateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Auth/CreateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CreateToken(ctx, req.(*CreateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_DecodeToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecodeTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).DecodeToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Auth/DecodeToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).DecodeToken(ctx, req.(*DecodeTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_DeleteToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).DeleteToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Auth/DeleteToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).DeleteToken(ctx, req.(*DeleteTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateToken",
			Handler:    _Auth_CreateToken_Handler,
		},
		{
			MethodName: "DecodeToken",
			Handler:    _Auth_DecodeToken_Handler,
		},
		{
			MethodName: "DeleteToken",
			Handler:    _Auth_DeleteToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}