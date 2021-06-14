// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1/oauth2/oauth2.proto

package oauth2

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Auth2AuthRequest struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Auth2AuthRequest) Reset()         { *m = Auth2AuthRequest{} }
func (m *Auth2AuthRequest) String() string { return proto.CompactTextString(m) }
func (*Auth2AuthRequest) ProtoMessage()    {}
func (*Auth2AuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a794c649d563384, []int{0}
}

func (m *Auth2AuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Auth2AuthRequest.Unmarshal(m, b)
}
func (m *Auth2AuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Auth2AuthRequest.Marshal(b, m, deterministic)
}
func (m *Auth2AuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Auth2AuthRequest.Merge(m, src)
}
func (m *Auth2AuthRequest) XXX_Size() int {
	return xxx_messageInfo_Auth2AuthRequest.Size(m)
}
func (m *Auth2AuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_Auth2AuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_Auth2AuthRequest proto.InternalMessageInfo

func (m *Auth2AuthRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

type Auth2AuthResponse struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Status               bool     `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Auth2AuthResponse) Reset()         { *m = Auth2AuthResponse{} }
func (m *Auth2AuthResponse) String() string { return proto.CompactTextString(m) }
func (*Auth2AuthResponse) ProtoMessage()    {}
func (*Auth2AuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a794c649d563384, []int{1}
}

func (m *Auth2AuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Auth2AuthResponse.Unmarshal(m, b)
}
func (m *Auth2AuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Auth2AuthResponse.Marshal(b, m, deterministic)
}
func (m *Auth2AuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Auth2AuthResponse.Merge(m, src)
}
func (m *Auth2AuthResponse) XXX_Size() int {
	return xxx_messageInfo_Auth2AuthResponse.Size(m)
}
func (m *Auth2AuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_Auth2AuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_Auth2AuthResponse proto.InternalMessageInfo

func (m *Auth2AuthResponse) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Auth2AuthResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

type Auth2LoginRequest struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Code                 string   `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Auth2LoginRequest) Reset()         { *m = Auth2LoginRequest{} }
func (m *Auth2LoginRequest) String() string { return proto.CompactTextString(m) }
func (*Auth2LoginRequest) ProtoMessage()    {}
func (*Auth2LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a794c649d563384, []int{2}
}

func (m *Auth2LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Auth2LoginRequest.Unmarshal(m, b)
}
func (m *Auth2LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Auth2LoginRequest.Marshal(b, m, deterministic)
}
func (m *Auth2LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Auth2LoginRequest.Merge(m, src)
}
func (m *Auth2LoginRequest) XXX_Size() int {
	return xxx_messageInfo_Auth2LoginRequest.Size(m)
}
func (m *Auth2LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_Auth2LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_Auth2LoginRequest proto.InternalMessageInfo

func (m *Auth2LoginRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *Auth2LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Auth2LoginRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type Auth2LoginResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Account              string   `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Status               bool     `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Auth2LoginResponse) Reset()         { *m = Auth2LoginResponse{} }
func (m *Auth2LoginResponse) String() string { return proto.CompactTextString(m) }
func (*Auth2LoginResponse) ProtoMessage()    {}
func (*Auth2LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a794c649d563384, []int{3}
}

func (m *Auth2LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Auth2LoginResponse.Unmarshal(m, b)
}
func (m *Auth2LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Auth2LoginResponse.Marshal(b, m, deterministic)
}
func (m *Auth2LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Auth2LoginResponse.Merge(m, src)
}
func (m *Auth2LoginResponse) XXX_Size() int {
	return xxx_messageInfo_Auth2LoginResponse.Size(m)
}
func (m *Auth2LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_Auth2LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_Auth2LoginResponse proto.InternalMessageInfo

func (m *Auth2LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Auth2LoginResponse) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *Auth2LoginResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Auth2LoginResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

type Auth2RegistRequest struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Auth2RegistRequest) Reset()         { *m = Auth2RegistRequest{} }
func (m *Auth2RegistRequest) String() string { return proto.CompactTextString(m) }
func (*Auth2RegistRequest) ProtoMessage()    {}
func (*Auth2RegistRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a794c649d563384, []int{4}
}

func (m *Auth2RegistRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Auth2RegistRequest.Unmarshal(m, b)
}
func (m *Auth2RegistRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Auth2RegistRequest.Marshal(b, m, deterministic)
}
func (m *Auth2RegistRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Auth2RegistRequest.Merge(m, src)
}
func (m *Auth2RegistRequest) XXX_Size() int {
	return xxx_messageInfo_Auth2RegistRequest.Size(m)
}
func (m *Auth2RegistRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_Auth2RegistRequest.DiscardUnknown(m)
}

var xxx_messageInfo_Auth2RegistRequest proto.InternalMessageInfo

func (m *Auth2RegistRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *Auth2RegistRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Auth2RegistResponse struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Auth2RegistResponse) Reset()         { *m = Auth2RegistResponse{} }
func (m *Auth2RegistResponse) String() string { return proto.CompactTextString(m) }
func (*Auth2RegistResponse) ProtoMessage()    {}
func (*Auth2RegistResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a794c649d563384, []int{5}
}

func (m *Auth2RegistResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Auth2RegistResponse.Unmarshal(m, b)
}
func (m *Auth2RegistResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Auth2RegistResponse.Marshal(b, m, deterministic)
}
func (m *Auth2RegistResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Auth2RegistResponse.Merge(m, src)
}
func (m *Auth2RegistResponse) XXX_Size() int {
	return xxx_messageInfo_Auth2RegistResponse.Size(m)
}
func (m *Auth2RegistResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_Auth2RegistResponse.DiscardUnknown(m)
}

var xxx_messageInfo_Auth2RegistResponse proto.InternalMessageInfo

func (m *Auth2RegistResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func init() {
	proto.RegisterType((*Auth2AuthRequest)(nil), "athena.auth2.Auth2AuthRequest")
	proto.RegisterType((*Auth2AuthResponse)(nil), "athena.auth2.Auth2AuthResponse")
	proto.RegisterType((*Auth2LoginRequest)(nil), "athena.auth2.Auth2LoginRequest")
	proto.RegisterType((*Auth2LoginResponse)(nil), "athena.auth2.Auth2LoginResponse")
	proto.RegisterType((*Auth2RegistRequest)(nil), "athena.auth2.Auth2RegistRequest")
	proto.RegisterType((*Auth2RegistResponse)(nil), "athena.auth2.Auth2RegistResponse")
}

func init() { proto.RegisterFile("v1/oauth2/oauth2.proto", fileDescriptor_1a794c649d563384) }

var fileDescriptor_1a794c649d563384 = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xb1, 0x4e, 0xeb, 0x30,
	0x14, 0x86, 0x95, 0xdc, 0xb4, 0xb7, 0xb5, 0x7a, 0x75, 0xe1, 0xb4, 0xaa, 0x42, 0x84, 0xda, 0xe2,
	0xa9, 0xaa, 0x20, 0x51, 0xcb, 0xd6, 0x0d, 0x24, 0x16, 0xc4, 0x14, 0x36, 0x24, 0x06, 0x93, 0x5a,
	0x69, 0xa0, 0xd8, 0x21, 0x76, 0xc2, 0xce, 0x2b, 0xf0, 0x60, 0x0c, 0xbc, 0x02, 0x0f, 0x82, 0xe2,
	0xb8, 0xc5, 0x41, 0x2d, 0x0b, 0x8b, 0xe3, 0x3f, 0xfe, 0xf5, 0x7f, 0xf6, 0xf1, 0x31, 0xea, 0x17,
	0xd3, 0x80, 0x93, 0x5c, 0x2e, 0x67, 0xfa, 0xe3, 0xa7, 0x19, 0x97, 0x1c, 0x3a, 0x44, 0x2e, 0x29,
	0x23, 0xbe, 0xfa, 0xe7, 0x1d, 0xc6, 0x9c, 0xc7, 0x2b, 0x1a, 0x90, 0x34, 0x09, 0x08, 0x63, 0x5c,
	0x12, 0x99, 0x70, 0x26, 0x2a, 0x2f, 0x3e, 0x46, 0x7b, 0x67, 0xa5, 0xad, 0x1c, 0x42, 0xfa, 0x94,
	0x53, 0x21, 0xc1, 0x45, 0x7f, 0x05, 0xcd, 0x8a, 0x24, 0xa2, 0xae, 0x35, 0xb2, 0xc6, 0xed, 0x70,
	0x2d, 0xf1, 0x05, 0xda, 0x37, 0xdc, 0x22, 0xe5, 0x4c, 0xd0, 0xdd, 0x76, 0xe8, 0xa3, 0xa6, 0x90,
	0x44, 0xe6, 0xc2, 0xfd, 0x33, 0xb2, 0xc6, 0xad, 0x50, 0x2b, 0x7c, 0xab, 0x63, 0xae, 0x78, 0x9c,
	0x30, 0x83, 0x4a, 0xa2, 0x88, 0xe7, 0x4c, 0xae, 0x63, 0xb4, 0x04, 0x0f, 0xb5, 0x52, 0x22, 0xc4,
	0x33, 0xcf, 0x16, 0xae, 0xad, 0x96, 0x36, 0x1a, 0x00, 0x39, 0x11, 0x5f, 0x50, 0x05, 0x68, 0x87,
	0x6a, 0x8e, 0x53, 0x04, 0x66, 0xbc, 0xde, 0x66, 0x0f, 0x35, 0x24, 0x7f, 0xa0, 0x4c, 0xa7, 0x57,
	0xc2, 0xa4, 0xda, 0x75, 0x2a, 0x20, 0x87, 0x91, 0xc7, 0x4d, 0x72, 0x39, 0x37, 0x0e, 0xe4, 0xd4,
	0x0e, 0x74, 0xa9, 0x89, 0x21, 0x8d, 0x13, 0x21, 0x7f, 0x75, 0x22, 0x7c, 0x82, 0xba, 0xb5, 0x2c,
	0xbd, 0xfd, 0x2f, 0xb4, 0x65, 0xa2, 0x67, 0x6f, 0x36, 0xea, 0x28, 0xff, 0x35, 0xcd, 0x92, 0x22,
	0xa2, 0x10, 0x21, 0xa7, 0xd4, 0x30, 0xf0, 0xcd, 0x36, 0xf0, 0xbf, 0xdf, 0xb2, 0x37, 0xdc, 0xb9,
	0x5e, 0x11, 0xb1, 0xf7, 0xf2, 0xfe, 0xf1, 0x6a, 0xf7, 0x00, 0x54, 0xeb, 0x14, 0xd3, 0xa0, 0x6a,
	0xb5, 0x72, 0x84, 0x7b, 0xd4, 0x50, 0xd5, 0x85, 0x6d, 0x29, 0xe6, 0xb5, 0x7a, 0xa3, 0xdd, 0x06,
	0xcd, 0x19, 0x28, 0x8e, 0x8b, 0xbb, 0x75, 0xce, 0xaa, 0x34, 0xcd, 0xad, 0x09, 0x30, 0xd4, 0xac,
	0x6a, 0x01, 0xdb, 0xb2, 0x6a, 0x25, 0xf7, 0x8e, 0x7e, 0x70, 0x68, 0xdc, 0x50, 0xe1, 0x0e, 0x70,
	0xaf, 0x8e, 0xcb, 0x94, 0x6b, 0x6e, 0x4d, 0xce, 0xff, 0xdf, 0xfc, 0xd3, 0x2b, 0xd5, 0xab, 0xba,
	0x6b, 0xaa, 0xa7, 0x72, 0xfa, 0x19, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x9f, 0xa6, 0xb2, 0x70, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Auth2SerivceClient is the client API for Auth2Serivce service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Auth2SerivceClient interface {
	Auth(ctx context.Context, in *Auth2AuthRequest, opts ...grpc.CallOption) (*Auth2AuthResponse, error)
	Login(ctx context.Context, in *Auth2LoginRequest, opts ...grpc.CallOption) (*Auth2LoginResponse, error)
	Regist(ctx context.Context, in *Auth2RegistRequest, opts ...grpc.CallOption) (*Auth2RegistResponse, error)
}

type auth2SerivceClient struct {
	cc *grpc.ClientConn
}

func NewAuth2SerivceClient(cc *grpc.ClientConn) Auth2SerivceClient {
	return &auth2SerivceClient{cc}
}

func (c *auth2SerivceClient) Auth(ctx context.Context, in *Auth2AuthRequest, opts ...grpc.CallOption) (*Auth2AuthResponse, error) {
	out := new(Auth2AuthResponse)
	err := c.cc.Invoke(ctx, "/athena.auth2.Auth2Serivce/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auth2SerivceClient) Login(ctx context.Context, in *Auth2LoginRequest, opts ...grpc.CallOption) (*Auth2LoginResponse, error) {
	out := new(Auth2LoginResponse)
	err := c.cc.Invoke(ctx, "/athena.auth2.Auth2Serivce/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auth2SerivceClient) Regist(ctx context.Context, in *Auth2RegistRequest, opts ...grpc.CallOption) (*Auth2RegistResponse, error) {
	out := new(Auth2RegistResponse)
	err := c.cc.Invoke(ctx, "/athena.auth2.Auth2Serivce/Regist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Auth2SerivceServer is the server API for Auth2Serivce service.
type Auth2SerivceServer interface {
	Auth(context.Context, *Auth2AuthRequest) (*Auth2AuthResponse, error)
	Login(context.Context, *Auth2LoginRequest) (*Auth2LoginResponse, error)
	Regist(context.Context, *Auth2RegistRequest) (*Auth2RegistResponse, error)
}

// UnimplementedAuth2SerivceServer can be embedded to have forward compatible implementations.
type UnimplementedAuth2SerivceServer struct {
}

func (*UnimplementedAuth2SerivceServer) Auth(ctx context.Context, req *Auth2AuthRequest) (*Auth2AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (*UnimplementedAuth2SerivceServer) Login(ctx context.Context, req *Auth2LoginRequest) (*Auth2LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedAuth2SerivceServer) Regist(ctx context.Context, req *Auth2RegistRequest) (*Auth2RegistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Regist not implemented")
}

func RegisterAuth2SerivceServer(s *grpc.Server, srv Auth2SerivceServer) {
	s.RegisterService(&_Auth2Serivce_serviceDesc, srv)
}

func _Auth2Serivce_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth2AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Auth2SerivceServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/athena.auth2.Auth2Serivce/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Auth2SerivceServer).Auth(ctx, req.(*Auth2AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth2Serivce_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth2LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Auth2SerivceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/athena.auth2.Auth2Serivce/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Auth2SerivceServer).Login(ctx, req.(*Auth2LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth2Serivce_Regist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth2RegistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Auth2SerivceServer).Regist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/athena.auth2.Auth2Serivce/Regist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Auth2SerivceServer).Regist(ctx, req.(*Auth2RegistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth2Serivce_serviceDesc = grpc.ServiceDesc{
	ServiceName: "athena.auth2.Auth2Serivce",
	HandlerType: (*Auth2SerivceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Auth",
			Handler:    _Auth2Serivce_Auth_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Auth2Serivce_Login_Handler,
		},
		{
			MethodName: "Regist",
			Handler:    _Auth2Serivce_Regist_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/oauth2/oauth2.proto",
}
