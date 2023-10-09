// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/params/v1beta1/query.proto

package proposal

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
	// subspace defines the module to query the parameter for.
	Subspace string `protobuf:"bytes,1,opt,name=subspace,proto3" json:"subspace,omitempty"`
	// key defines the key of the parameter in the subspace.
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b32979c1792ccc4, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

func (m *QueryParamsRequest) GetSubspace() string {
	if m != nil {
		return m.Subspace
	}
	return ""
}

func (m *QueryParamsRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// param defines the queried parameter.
	Param ParamChange `protobuf:"bytes,1,opt,name=param,proto3" json:"param"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b32979c1792ccc4, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParam() ParamChange {
	if m != nil {
		return m.Param
	}
	return ParamChange{}
}

// QuerySubspacesRequest defines a request type for querying for all registered
// subspaces and all keys for a subspace.
//
// Since: cosmos-sdk 0.46
type QuerySubspacesRequest struct {
}

func (m *QuerySubspacesRequest) Reset()         { *m = QuerySubspacesRequest{} }
func (m *QuerySubspacesRequest) String() string { return proto.CompactTextString(m) }
func (*QuerySubspacesRequest) ProtoMessage()    {}
func (*QuerySubspacesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b32979c1792ccc4, []int{2}
}
func (m *QuerySubspacesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySubspacesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySubspacesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySubspacesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySubspacesRequest.Merge(m, src)
}
func (m *QuerySubspacesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QuerySubspacesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySubspacesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySubspacesRequest proto.InternalMessageInfo

// QuerySubspacesResponse defines the response types for querying for all
// registered subspaces and all keys for a subspace.
//
// Since: cosmos-sdk 0.46
type QuerySubspacesResponse struct {
	Subspaces []*Subspace `protobuf:"bytes,1,rep,name=subspaces,proto3" json:"subspaces,omitempty"`
}

func (m *QuerySubspacesResponse) Reset()         { *m = QuerySubspacesResponse{} }
func (m *QuerySubspacesResponse) String() string { return proto.CompactTextString(m) }
func (*QuerySubspacesResponse) ProtoMessage()    {}
func (*QuerySubspacesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b32979c1792ccc4, []int{3}
}
func (m *QuerySubspacesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySubspacesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySubspacesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySubspacesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySubspacesResponse.Merge(m, src)
}
func (m *QuerySubspacesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QuerySubspacesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySubspacesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySubspacesResponse proto.InternalMessageInfo

func (m *QuerySubspacesResponse) GetSubspaces() []*Subspace {
	if m != nil {
		return m.Subspaces
	}
	return nil
}

// Subspace defines a parameter subspace name and all the keys that exist for
// the subspace.
//
// Since: cosmos-sdk 0.46
type Subspace struct {
	Subspace string   `protobuf:"bytes,1,opt,name=subspace,proto3" json:"subspace,omitempty"`
	Keys     []string `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (m *Subspace) Reset()         { *m = Subspace{} }
func (m *Subspace) String() string { return proto.CompactTextString(m) }
func (*Subspace) ProtoMessage()    {}
func (*Subspace) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b32979c1792ccc4, []int{4}
}
func (m *Subspace) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Subspace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Subspace.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Subspace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subspace.Merge(m, src)
}
func (m *Subspace) XXX_Size() int {
	return m.Size()
}
func (m *Subspace) XXX_DiscardUnknown() {
	xxx_messageInfo_Subspace.DiscardUnknown(m)
}

var xxx_messageInfo_Subspace proto.InternalMessageInfo

func (m *Subspace) GetSubspace() string {
	if m != nil {
		return m.Subspace
	}
	return ""
}

func (m *Subspace) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "cosmos.params.v1beta1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "cosmos.params.v1beta1.QueryParamsResponse")
	proto.RegisterType((*QuerySubspacesRequest)(nil), "cosmos.params.v1beta1.QuerySubspacesRequest")
	proto.RegisterType((*QuerySubspacesResponse)(nil), "cosmos.params.v1beta1.QuerySubspacesResponse")
	proto.RegisterType((*Subspace)(nil), "cosmos.params.v1beta1.Subspace")
}

func init() { proto.RegisterFile("cosmos/params/v1beta1/query.proto", fileDescriptor_2b32979c1792ccc4) }

var fileDescriptor_2b32979c1792ccc4 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcf, 0x8e, 0xd2, 0x40,
	0x18, 0xef, 0x14, 0x21, 0x74, 0xb8, 0xe8, 0x28, 0xda, 0x34, 0x5a, 0xea, 0x44, 0x93, 0x4a, 0xb4,
	0x13, 0xf0, 0x66, 0xa2, 0x07, 0x78, 0x01, 0xad, 0x07, 0x13, 0x6e, 0x03, 0x4e, 0x6a, 0x03, 0x74,
	0x86, 0x4e, 0x31, 0xf6, 0xea, 0xc1, 0xb3, 0xd1, 0x67, 0x30, 0xf1, 0xe8, 0x63, 0x70, 0x24, 0xf1,
	0xe2, 0xc9, 0x6c, 0x60, 0x93, 0x7d, 0x8d, 0x0d, 0x33, 0x2d, 0x9b, 0x65, 0x81, 0x70, 0x69, 0xbe,
	0x7e, 0xf3, 0xfb, 0x37, 0xbf, 0x16, 0x3e, 0x1e, 0x71, 0x39, 0xe5, 0x92, 0x08, 0x9a, 0xd2, 0xa9,
	0x24, 0x9f, 0x3b, 0x43, 0x96, 0xd1, 0x0e, 0x99, 0xcd, 0x59, 0x9a, 0x07, 0x22, 0xe5, 0x19, 0x47,
	0x4d, 0x0d, 0x09, 0x34, 0x24, 0x28, 0x20, 0xce, 0xbd, 0x88, 0x47, 0x5c, 0x21, 0xc8, 0x66, 0xd2,
	0x60, 0xe7, 0x61, 0xc4, 0x79, 0x34, 0x61, 0x84, 0x8a, 0x98, 0xd0, 0x24, 0xe1, 0x19, 0xcd, 0x62,
	0x9e, 0xc8, 0xe2, 0x14, 0xef, 0x77, 0x2b, 0x94, 0x35, 0xe6, 0x0e, 0x9d, 0xc6, 0x09, 0x27, 0xea,
	0xa9, 0x57, 0xb8, 0x07, 0xd1, 0xbb, 0x4d, 0xa0, 0xb7, 0x0a, 0x17, 0xb2, 0xd9, 0x9c, 0xc9, 0x0c,
	0x39, 0xb0, 0x2e, 0xe7, 0x43, 0x29, 0xe8, 0x88, 0xd9, 0xc0, 0x03, 0xbe, 0x15, 0x6e, 0xdf, 0xd1,
	0x6d, 0x58, 0x19, 0xb3, 0xdc, 0x36, 0xd5, 0x7a, 0x33, 0xe2, 0x01, 0xbc, 0x7b, 0x4d, 0x43, 0x0a,
	0x9e, 0x48, 0x86, 0xfa, 0xb0, 0xaa, 0xdc, 0x95, 0x42, 0xa3, 0x8b, 0x83, 0xbd, 0x97, 0x0d, 0x14,
	0xab, 0xff, 0x89, 0x26, 0x11, 0xeb, 0x59, 0x8b, 0xff, 0x2d, 0xe3, 0xf7, 0xc5, 0x9f, 0x36, 0x08,
	0x35, 0x17, 0x3f, 0x80, 0x4d, 0xa5, 0xfd, 0xbe, 0xb0, 0x2f, 0x23, 0xe2, 0x0f, 0xf0, 0xfe, 0xee,
	0x41, 0xe1, 0xfb, 0x1a, 0x5a, 0x65, 0x58, 0x69, 0x03, 0xaf, 0xe2, 0x37, 0xba, 0xad, 0x03, 0xde,
	0x25, 0x39, 0xbc, 0x62, 0xe0, 0x57, 0xb0, 0x5e, 0xae, 0x8f, 0xf6, 0x80, 0xe0, 0xad, 0x31, 0xcb,
	0xa5, 0x6d, 0x7a, 0x15, 0xdf, 0x0a, 0xd5, 0xdc, 0xfd, 0x65, 0xc2, 0xaa, 0x4a, 0x85, 0xbe, 0x01,
	0x58, 0xd3, 0x7d, 0xa0, 0x67, 0x07, 0xcc, 0x6f, 0xf6, 0xee, 0xb4, 0x4f, 0x81, 0xea, 0x6b, 0xe2,
	0xa7, 0x5f, 0xff, 0x9e, 0xff, 0x34, 0x5b, 0xe8, 0x11, 0x39, 0xf6, 0xe5, 0xd1, 0x0f, 0x00, 0xad,
	0x6d, 0x47, 0xe8, 0xf9, 0x31, 0x83, 0xdd, 0x8e, 0x9d, 0x17, 0x27, 0xa2, 0x8b, 0x44, 0xbe, 0x4a,
	0x84, 0x91, 0x77, 0x20, 0xd1, 0xb6, 0xe3, 0xde, 0x9b, 0xc5, 0xca, 0x05, 0xcb, 0x95, 0x0b, 0xce,
	0x56, 0x2e, 0xf8, 0xbe, 0x76, 0x8d, 0xe5, 0xda, 0x35, 0xfe, 0xad, 0x5d, 0x63, 0xf0, 0x44, 0x53,
	0xe5, 0xc7, 0x71, 0x10, 0x73, 0xf2, 0xa5, 0x94, 0xc8, 0x72, 0xc1, 0x24, 0x11, 0x29, 0x17, 0x5c,
	0xd2, 0xc9, 0xb0, 0xa6, 0x7e, 0xde, 0x97, 0x97, 0x01, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x00, 0x44,
	0xa1, 0x63, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Params queries a specific parameter of a module, given its subspace and
	// key.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Subspaces queries for all registered subspaces and all keys for a subspace.
	//
	// Since: cosmos-sdk 0.46
	Subspaces(ctx context.Context, in *QuerySubspacesRequest, opts ...grpc.CallOption) (*QuerySubspacesResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.params.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Subspaces(ctx context.Context, in *QuerySubspacesRequest, opts ...grpc.CallOption) (*QuerySubspacesResponse, error) {
	out := new(QuerySubspacesResponse)
	err := c.cc.Invoke(ctx, "/cosmos.params.v1beta1.Query/Subspaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Params queries a specific parameter of a module, given its subspace and
	// key.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Subspaces queries for all registered subspaces and all keys for a subspace.
	//
	// Since: cosmos-sdk 0.46
	Subspaces(context.Context, *QuerySubspacesRequest) (*QuerySubspacesResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) Subspaces(ctx context.Context, req *QuerySubspacesRequest) (*QuerySubspacesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subspaces not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.params.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Subspaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySubspacesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Subspaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.params.v1beta1.Query/Subspaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Subspaces(ctx, req.(*QuerySubspacesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.params.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Subspaces",
			Handler:    _Query_Subspaces_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/params/v1beta1/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Subspace) > 0 {
		i -= len(m.Subspace)
		copy(dAtA[i:], m.Subspace)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Subspace)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Param.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QuerySubspacesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySubspacesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySubspacesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QuerySubspacesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySubspacesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySubspacesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Subspaces) > 0 {
		for iNdEx := len(m.Subspaces) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Subspaces[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Subspace) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Subspace) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Subspace) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Keys) > 0 {
		for iNdEx := len(m.Keys) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Keys[iNdEx])
			copy(dAtA[i:], m.Keys[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.Keys[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Subspace) > 0 {
		i -= len(m.Subspace)
		copy(dAtA[i:], m.Subspace)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Subspace)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Subspace)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Param.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QuerySubspacesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QuerySubspacesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Subspaces) > 0 {
		for _, e := range m.Subspaces {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *Subspace) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Subspace)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if len(m.Keys) > 0 {
		for _, s := range m.Keys {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subspace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subspace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Param", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Param.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QuerySubspacesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QuerySubspacesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySubspacesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QuerySubspacesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QuerySubspacesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySubspacesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subspaces", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subspaces = append(m.Subspaces, &Subspace{})
			if err := m.Subspaces[len(m.Subspaces)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Subspace) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Subspace: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Subspace: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subspace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subspace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Keys", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Keys = append(m.Keys, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
