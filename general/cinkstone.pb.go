// Code generated by protoc-gen-go. DO NOT EDIT.
// source: general/cinkstone.proto

package cinkstone

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

//
// Version information
type GetVersionResponse struct {
	Commit               string               `protobuf:"bytes,1,opt,name=commit,proto3" json:"commit,omitempty"`
	Revision             string               `protobuf:"bytes,2,opt,name=revision,proto3" json:"revision,omitempty"`
	Branch               string               `protobuf:"bytes,3,opt,name=branch,proto3" json:"branch,omitempty"`
	Tag                  string               `protobuf:"bytes,4,opt,name=tag,proto3" json:"tag,omitempty"`
	BuildDate            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=build_date,json=buildDate,proto3" json:"build_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetVersionResponse) Reset()         { *m = GetVersionResponse{} }
func (m *GetVersionResponse) String() string { return proto.CompactTextString(m) }
func (*GetVersionResponse) ProtoMessage()    {}
func (*GetVersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_97f89674e9fd3050, []int{0}
}

func (m *GetVersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVersionResponse.Unmarshal(m, b)
}
func (m *GetVersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVersionResponse.Marshal(b, m, deterministic)
}
func (m *GetVersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVersionResponse.Merge(m, src)
}
func (m *GetVersionResponse) XXX_Size() int {
	return xxx_messageInfo_GetVersionResponse.Size(m)
}
func (m *GetVersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetVersionResponse proto.InternalMessageInfo

func (m *GetVersionResponse) GetCommit() string {
	if m != nil {
		return m.Commit
	}
	return ""
}

func (m *GetVersionResponse) GetRevision() string {
	if m != nil {
		return m.Revision
	}
	return ""
}

func (m *GetVersionResponse) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

func (m *GetVersionResponse) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *GetVersionResponse) GetBuildDate() *timestamp.Timestamp {
	if m != nil {
		return m.BuildDate
	}
	return nil
}

func init() {
	proto.RegisterType((*GetVersionResponse)(nil), "cinkstone.GetVersionResponse")
}

func init() { proto.RegisterFile("general/cinkstone.proto", fileDescriptor_97f89674e9fd3050) }

var fileDescriptor_97f89674e9fd3050 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x49, 0xab, 0xc5, 0x8c, 0x22, 0xb2, 0x48, 0x0d, 0x51, 0xb1, 0xf4, 0xd4, 0x53, 0x02,
	0xf5, 0xe4, 0x51, 0x50, 0xbc, 0x47, 0x11, 0x3c, 0xc9, 0x26, 0x1d, 0xe3, 0x62, 0xf6, 0x0f, 0xbb,
	0xd3, 0x80, 0x57, 0x5f, 0xc1, 0xe7, 0xf0, 0x69, 0x7c, 0x05, 0x1f, 0x44, 0xb2, 0x9b, 0xb6, 0xd0,
	0xde, 0xf6, 0xdb, 0xef, 0x37, 0xc3, 0xf7, 0x0d, 0x9c, 0xd5, 0xa8, 0xd0, 0xf2, 0x26, 0xaf, 0x84,
	0xfa, 0x70, 0xa4, 0x15, 0x66, 0xc6, 0x6a, 0xd2, 0x2c, 0x5e, 0x7f, 0xa4, 0x17, 0xb5, 0xd6, 0x75,
	0x83, 0x39, 0x37, 0x22, 0xe7, 0x4a, 0x69, 0xe2, 0x24, 0xb4, 0x72, 0x01, 0x4c, 0xaf, 0x7a, 0xd7,
	0xab, 0x72, 0xf9, 0x96, 0x93, 0x90, 0xe8, 0x88, 0x4b, 0xd3, 0x03, 0xe7, 0xdb, 0x00, 0x4a, 0x43,
	0x9f, 0xc1, 0x9c, 0xfe, 0x44, 0xc0, 0x1e, 0x90, 0x9e, 0xd1, 0x3a, 0xa1, 0x55, 0x81, 0xce, 0x68,
	0xe5, 0x90, 0x8d, 0x61, 0x54, 0x69, 0x29, 0x05, 0x25, 0xd1, 0x24, 0x9a, 0xc5, 0x45, 0xaf, 0x58,
	0x0a, 0x07, 0x16, 0x5b, 0xd1, 0xb1, 0xc9, 0xc0, 0x3b, 0x6b, 0xdd, 0xcd, 0x94, 0x96, 0xab, 0xea,
	0x3d, 0x19, 0x86, 0x99, 0xa0, 0xd8, 0x09, 0x0c, 0x89, 0xd7, 0xc9, 0x9e, 0xff, 0xec, 0x9e, 0xec,
	0x06, 0xa0, 0x5c, 0x8a, 0x66, 0xf1, 0xba, 0xe0, 0x84, 0xc9, 0xfe, 0x24, 0x9a, 0x1d, 0xce, 0xd3,
	0x2c, 0xc4, 0xcc, 0x56, 0x31, 0xb3, 0xa7, 0x55, 0x8f, 0x22, 0xf6, 0xf4, 0x1d, 0x27, 0x9c, 0xd7,
	0x00, 0xb7, 0xc6, 0x3c, 0xa2, 0x6d, 0x45, 0x85, 0xec, 0x05, 0x60, 0x13, 0x9e, 0x8d, 0x77, 0x56,
	0xdc, 0x77, 0x4d, 0xd3, 0xcb, 0x6c, 0x73, 0xdc, 0xdd, 0xae, 0xd3, 0xd3, 0xaf, 0xdf, 0xbf, 0xef,
	0xc1, 0x31, 0x3b, 0xf2, 0x07, 0x6e, 0x83, 0x5b, 0x8e, 0xfc, 0x92, 0xeb, 0xff, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x0f, 0xfe, 0x24, 0xfa, 0xa1, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AppServiceClient is the client API for AppService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AppServiceClient interface {
	/// Get current version information
	GetVersion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetVersionResponse, error)
}

type appServiceClient struct {
	cc *grpc.ClientConn
}

func NewAppServiceClient(cc *grpc.ClientConn) AppServiceClient {
	return &appServiceClient{cc}
}

func (c *appServiceClient) GetVersion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetVersionResponse, error) {
	out := new(GetVersionResponse)
	err := c.cc.Invoke(ctx, "/cinkstone.AppService/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppServiceServer is the server API for AppService service.
type AppServiceServer interface {
	/// Get current version information
	GetVersion(context.Context, *empty.Empty) (*GetVersionResponse, error)
}

func RegisterAppServiceServer(s *grpc.Server, srv AppServiceServer) {
	s.RegisterService(&_AppService_serviceDesc, srv)
}

func _AppService_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cinkstone.AppService/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).GetVersion(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _AppService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cinkstone.AppService",
	HandlerType: (*AppServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _AppService_GetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "general/cinkstone.proto",
}
