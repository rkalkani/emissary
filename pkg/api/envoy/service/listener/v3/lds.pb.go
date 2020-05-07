// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/service/listener/v3/lds.proto

package envoy_service_listener_v3

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/datawire/ambassador/pkg/api/envoy/annotations"
	v3 "github.com/datawire/ambassador/pkg/api/envoy/service/discovery/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	_ "istio.io/gogo-genproto/googleapis/google/api"
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

// [#not-implemented-hide:] Not configuration. Workaround c++ protobuf issue with importing
// services: https://github.com/google/protobuf/issues/4221 and protoxform to upgrade the file.
type LdsDummy struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdsDummy) Reset()         { *m = LdsDummy{} }
func (m *LdsDummy) String() string { return proto.CompactTextString(m) }
func (*LdsDummy) ProtoMessage()    {}
func (*LdsDummy) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f4f7cdedc179239, []int{0}
}
func (m *LdsDummy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LdsDummy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LdsDummy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LdsDummy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdsDummy.Merge(m, src)
}
func (m *LdsDummy) XXX_Size() int {
	return m.Size()
}
func (m *LdsDummy) XXX_DiscardUnknown() {
	xxx_messageInfo_LdsDummy.DiscardUnknown(m)
}

var xxx_messageInfo_LdsDummy proto.InternalMessageInfo

func init() {
	proto.RegisterType((*LdsDummy)(nil), "envoy.service.listener.v3.LdsDummy")
}

func init() {
	proto.RegisterFile("envoy/service/listener/v3/lds.proto", fileDescriptor_2f4f7cdedc179239)
}

var fileDescriptor_2f4f7cdedc179239 = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0xc1, 0xaa, 0xd3, 0x40,
	0x14, 0x86, 0x1d, 0x2f, 0xa8, 0xcc, 0xe2, 0x0a, 0x01, 0xa9, 0x37, 0x5c, 0x82, 0xb7, 0x77, 0x61,
	0x2d, 0x3a, 0xa3, 0xcd, 0xae, 0xcb, 0x12, 0x5c, 0x75, 0x51, 0xec, 0x13, 0x4c, 0x93, 0xd3, 0x38,
	0x90, 0xce, 0x8c, 0x33, 0x93, 0xd1, 0x2c, 0xdc, 0xb8, 0x2a, 0x6e, 0x05, 0x05, 0xd7, 0xbe, 0x8a,
	0xe0, 0x52, 0xd0, 0x07, 0x90, 0xe2, 0x1b, 0xf8, 0x02, 0xd2, 0x24, 0xd3, 0x36, 0x42, 0x45, 0x37,
	0x77, 0x97, 0xe1, 0x7c, 0xe7, 0xff, 0x0f, 0x3f, 0x7f, 0xf0, 0x25, 0x08, 0x27, 0x2b, 0x6a, 0x40,
	0x3b, 0x9e, 0x02, 0x2d, 0xb8, 0xb1, 0x20, 0x40, 0x53, 0x17, 0xd3, 0x22, 0x33, 0x44, 0x69, 0x69,
	0x65, 0x70, 0x56, 0x43, 0xa4, 0x85, 0x88, 0x87, 0x88, 0x8b, 0xc3, 0x61, 0x77, 0x3f, 0xe3, 0x26,
	0x95, 0x0e, 0x74, 0xb5, 0x15, 0xd8, 0x3d, 0x1a, 0x99, 0xf0, 0x3c, 0x97, 0x32, 0x2f, 0x80, 0x32,
	0xc5, 0x29, 0x13, 0x42, 0x5a, 0x66, 0xb9, 0x14, 0xad, 0x49, 0x18, 0xb5, 0xd3, 0xfa, 0xb5, 0x28,
	0x97, 0x34, 0x2b, 0x75, 0x0d, 0x1c, 0x9b, 0xbf, 0xd4, 0x4c, 0x29, 0xd0, 0x7e, 0xff, 0xa2, 0xcc,
	0x14, 0x3b, 0xd4, 0xa5, 0x0e, 0xb4, 0xe1, 0x52, 0x70, 0x91, 0xb7, 0xc8, 0xbd, 0xe6, 0xd8, 0x43,
	0x46, 0x83, 0x91, 0xa5, 0x4e, 0xa1, 0x25, 0x7a, 0x8e, 0x15, 0x3c, 0x63, 0x16, 0xa8, 0xff, 0x68,
	0x06, 0xfd, 0x01, 0xbe, 0x35, 0xcd, 0x4c, 0x52, 0xae, 0x56, 0xd5, 0xf8, 0xfc, 0xe3, 0xe7, 0x75,
	0xd4, 0xc3, 0x77, 0x9a, 0x54, 0x98, 0xe2, 0xc4, 0x8d, 0x88, 0x9f, 0x8e, 0xbe, 0x9f, 0xe0, 0xbb,
	0xd3, 0x36, 0xa1, 0xc4, 0x27, 0x30, 0x6f, 0xf2, 0x09, 0x5e, 0xe3, 0xd3, 0x04, 0x0a, 0xcb, 0x3c,
	0x60, 0x82, 0x27, 0xa4, 0x1b, 0xee, 0x3e, 0x34, 0x17, 0x93, 0x9a, 0xdd, 0x89, 0x3c, 0x83, 0x17,
	0x25, 0x18, 0x1b, 0x8e, 0xfe, 0x67, 0xc5, 0x28, 0x29, 0x0c, 0xf4, 0xaf, 0x0d, 0xd0, 0x63, 0x14,
	0x58, 0x7c, 0x7b, 0x6e, 0x35, 0xb0, 0xd5, 0xde, 0xff, 0xe1, 0x5f, 0xc5, 0xfe, 0xb4, 0x7e, 0xf4,
	0x8f, 0x74, 0xc7, 0xf5, 0x03, 0xc2, 0xa7, 0x4f, 0xc1, 0xa6, 0xcf, 0xaf, 0xc8, 0x75, 0xf0, 0xe6,
	0xdb, 0xcf, 0x77, 0xd7, 0xcf, 0xfa, 0xbd, 0x4e, 0x09, 0xc7, 0xbe, 0xb6, 0xa6, 0x1e, 0x9f, 0x8c,
	0xd1, 0x30, 0x7c, 0xf0, 0xf6, 0xd3, 0xfb, 0x5f, 0x37, 0x2f, 0xf1, 0x45, 0xa3, 0x9f, 0x4a, 0xb1,
	0xe4, 0xf9, 0x61, 0xbf, 0x89, 0x3f, 0x79, 0x92, 0x7c, 0xd9, 0x44, 0xe8, 0xeb, 0x26, 0x42, 0x3f,
	0x36, 0x11, 0xc2, 0xf7, 0xb9, 0x6c, 0x6e, 0x52, 0x5a, 0xbe, 0xaa, 0xc8, 0xd1, 0xff, 0x63, 0xb2,
	0x6d, 0xcd, 0x6c, 0xdb, 0xa0, 0x19, 0x5a, 0x23, 0xb4, 0xb8, 0x51, 0xb7, 0x29, 0xfe, 0x1d, 0x00,
	0x00, 0xff, 0xff, 0xfe, 0x97, 0xaa, 0xf7, 0x77, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ListenerDiscoveryServiceClient is the client API for ListenerDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ListenerDiscoveryServiceClient interface {
	DeltaListeners(ctx context.Context, opts ...grpc.CallOption) (ListenerDiscoveryService_DeltaListenersClient, error)
	StreamListeners(ctx context.Context, opts ...grpc.CallOption) (ListenerDiscoveryService_StreamListenersClient, error)
	FetchListeners(ctx context.Context, in *v3.DiscoveryRequest, opts ...grpc.CallOption) (*v3.DiscoveryResponse, error)
}

type listenerDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewListenerDiscoveryServiceClient(cc *grpc.ClientConn) ListenerDiscoveryServiceClient {
	return &listenerDiscoveryServiceClient{cc}
}

func (c *listenerDiscoveryServiceClient) DeltaListeners(ctx context.Context, opts ...grpc.CallOption) (ListenerDiscoveryService_DeltaListenersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ListenerDiscoveryService_serviceDesc.Streams[0], "/envoy.service.listener.v3.ListenerDiscoveryService/DeltaListeners", opts...)
	if err != nil {
		return nil, err
	}
	x := &listenerDiscoveryServiceDeltaListenersClient{stream}
	return x, nil
}

type ListenerDiscoveryService_DeltaListenersClient interface {
	Send(*v3.DeltaDiscoveryRequest) error
	Recv() (*v3.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type listenerDiscoveryServiceDeltaListenersClient struct {
	grpc.ClientStream
}

func (x *listenerDiscoveryServiceDeltaListenersClient) Send(m *v3.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *listenerDiscoveryServiceDeltaListenersClient) Recv() (*v3.DeltaDiscoveryResponse, error) {
	m := new(v3.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *listenerDiscoveryServiceClient) StreamListeners(ctx context.Context, opts ...grpc.CallOption) (ListenerDiscoveryService_StreamListenersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ListenerDiscoveryService_serviceDesc.Streams[1], "/envoy.service.listener.v3.ListenerDiscoveryService/StreamListeners", opts...)
	if err != nil {
		return nil, err
	}
	x := &listenerDiscoveryServiceStreamListenersClient{stream}
	return x, nil
}

type ListenerDiscoveryService_StreamListenersClient interface {
	Send(*v3.DiscoveryRequest) error
	Recv() (*v3.DiscoveryResponse, error)
	grpc.ClientStream
}

type listenerDiscoveryServiceStreamListenersClient struct {
	grpc.ClientStream
}

func (x *listenerDiscoveryServiceStreamListenersClient) Send(m *v3.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *listenerDiscoveryServiceStreamListenersClient) Recv() (*v3.DiscoveryResponse, error) {
	m := new(v3.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *listenerDiscoveryServiceClient) FetchListeners(ctx context.Context, in *v3.DiscoveryRequest, opts ...grpc.CallOption) (*v3.DiscoveryResponse, error) {
	out := new(v3.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.listener.v3.ListenerDiscoveryService/FetchListeners", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListenerDiscoveryServiceServer is the server API for ListenerDiscoveryService service.
type ListenerDiscoveryServiceServer interface {
	DeltaListeners(ListenerDiscoveryService_DeltaListenersServer) error
	StreamListeners(ListenerDiscoveryService_StreamListenersServer) error
	FetchListeners(context.Context, *v3.DiscoveryRequest) (*v3.DiscoveryResponse, error)
}

// UnimplementedListenerDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedListenerDiscoveryServiceServer struct {
}

func (*UnimplementedListenerDiscoveryServiceServer) DeltaListeners(srv ListenerDiscoveryService_DeltaListenersServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaListeners not implemented")
}
func (*UnimplementedListenerDiscoveryServiceServer) StreamListeners(srv ListenerDiscoveryService_StreamListenersServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamListeners not implemented")
}
func (*UnimplementedListenerDiscoveryServiceServer) FetchListeners(ctx context.Context, req *v3.DiscoveryRequest) (*v3.DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchListeners not implemented")
}

func RegisterListenerDiscoveryServiceServer(s *grpc.Server, srv ListenerDiscoveryServiceServer) {
	s.RegisterService(&_ListenerDiscoveryService_serviceDesc, srv)
}

func _ListenerDiscoveryService_DeltaListeners_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ListenerDiscoveryServiceServer).DeltaListeners(&listenerDiscoveryServiceDeltaListenersServer{stream})
}

type ListenerDiscoveryService_DeltaListenersServer interface {
	Send(*v3.DeltaDiscoveryResponse) error
	Recv() (*v3.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type listenerDiscoveryServiceDeltaListenersServer struct {
	grpc.ServerStream
}

func (x *listenerDiscoveryServiceDeltaListenersServer) Send(m *v3.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *listenerDiscoveryServiceDeltaListenersServer) Recv() (*v3.DeltaDiscoveryRequest, error) {
	m := new(v3.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ListenerDiscoveryService_StreamListeners_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ListenerDiscoveryServiceServer).StreamListeners(&listenerDiscoveryServiceStreamListenersServer{stream})
}

type ListenerDiscoveryService_StreamListenersServer interface {
	Send(*v3.DiscoveryResponse) error
	Recv() (*v3.DiscoveryRequest, error)
	grpc.ServerStream
}

type listenerDiscoveryServiceStreamListenersServer struct {
	grpc.ServerStream
}

func (x *listenerDiscoveryServiceStreamListenersServer) Send(m *v3.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *listenerDiscoveryServiceStreamListenersServer) Recv() (*v3.DiscoveryRequest, error) {
	m := new(v3.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ListenerDiscoveryService_FetchListeners_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListenerDiscoveryServiceServer).FetchListeners(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.listener.v3.ListenerDiscoveryService/FetchListeners",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListenerDiscoveryServiceServer).FetchListeners(ctx, req.(*v3.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ListenerDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.listener.v3.ListenerDiscoveryService",
	HandlerType: (*ListenerDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchListeners",
			Handler:    _ListenerDiscoveryService_FetchListeners_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DeltaListeners",
			Handler:       _ListenerDiscoveryService_DeltaListeners_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamListeners",
			Handler:       _ListenerDiscoveryService_StreamListeners_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/listener/v3/lds.proto",
}

func (m *LdsDummy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LdsDummy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LdsDummy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func encodeVarintLds(dAtA []byte, offset int, v uint64) int {
	offset -= sovLds(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LdsDummy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovLds(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLds(x uint64) (n int) {
	return sovLds(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LdsDummy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLds
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
			return fmt.Errorf("proto: LdsDummy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LdsDummy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipLds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLds
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLds
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipLds(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLds
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
					return 0, ErrIntOverflowLds
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
					return 0, ErrIntOverflowLds
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
				return 0, ErrInvalidLengthLds
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLds
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLds
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLds        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLds          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLds = fmt.Errorf("proto: unexpected end of group")
)