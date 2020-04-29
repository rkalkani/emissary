// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/data/tap/v2alpha/common.proto

package envoy_data_tap_v2alpha

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

// Wrapper for tapped body data. This includes HTTP request/response body, transport socket received
// and transmitted data, etc.
type Body struct {
	// Types that are valid to be assigned to BodyType:
	//	*Body_AsBytes
	//	*Body_AsString
	BodyType isBody_BodyType `protobuf_oneof:"body_type"`
	// Specifies whether body data has been truncated to fit within the specified
	// :ref:`max_buffered_rx_bytes
	// <envoy_api_field_service.tap.v2alpha.OutputConfig.max_buffered_rx_bytes>` and
	// :ref:`max_buffered_tx_bytes
	// <envoy_api_field_service.tap.v2alpha.OutputConfig.max_buffered_tx_bytes>` settings.
	Truncated            bool     `protobuf:"varint,3,opt,name=truncated,proto3" json:"truncated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Body) Reset()         { *m = Body{} }
func (m *Body) String() string { return proto.CompactTextString(m) }
func (*Body) ProtoMessage()    {}
func (*Body) Descriptor() ([]byte, []int) {
	return fileDescriptor_a560f1e2899ebe7a, []int{0}
}
func (m *Body) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Body) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Body.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Body) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Body.Merge(m, src)
}
func (m *Body) XXX_Size() int {
	return m.Size()
}
func (m *Body) XXX_DiscardUnknown() {
	xxx_messageInfo_Body.DiscardUnknown(m)
}

var xxx_messageInfo_Body proto.InternalMessageInfo

type isBody_BodyType interface {
	isBody_BodyType()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Body_AsBytes struct {
	AsBytes []byte `protobuf:"bytes,1,opt,name=as_bytes,json=asBytes,proto3,oneof" json:"as_bytes,omitempty"`
}
type Body_AsString struct {
	AsString string `protobuf:"bytes,2,opt,name=as_string,json=asString,proto3,oneof" json:"as_string,omitempty"`
}

func (*Body_AsBytes) isBody_BodyType()  {}
func (*Body_AsString) isBody_BodyType() {}

func (m *Body) GetBodyType() isBody_BodyType {
	if m != nil {
		return m.BodyType
	}
	return nil
}

func (m *Body) GetAsBytes() []byte {
	if x, ok := m.GetBodyType().(*Body_AsBytes); ok {
		return x.AsBytes
	}
	return nil
}

func (m *Body) GetAsString() string {
	if x, ok := m.GetBodyType().(*Body_AsString); ok {
		return x.AsString
	}
	return ""
}

func (m *Body) GetTruncated() bool {
	if m != nil {
		return m.Truncated
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Body) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Body_AsBytes)(nil),
		(*Body_AsString)(nil),
	}
}

func init() {
	proto.RegisterType((*Body)(nil), "envoy.data.tap.v2alpha.Body")
}

func init() {
	proto.RegisterFile("envoy/data/tap/v2alpha/common.proto", fileDescriptor_a560f1e2899ebe7a)
}

var fileDescriptor_a560f1e2899ebe7a = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x41, 0x4a, 0xc4, 0x30,
	0x14, 0x40, 0xe7, 0xab, 0x68, 0x9b, 0x71, 0xd5, 0x85, 0x14, 0xd4, 0x52, 0xd4, 0x45, 0x57, 0x09,
	0xe8, 0x09, 0x8c, 0x9b, 0x59, 0x0e, 0xf5, 0x00, 0xe5, 0x77, 0x12, 0x74, 0xc0, 0xe6, 0x87, 0xe4,
	0x3b, 0x98, 0x1b, 0xba, 0xf4, 0x08, 0xd2, 0x93, 0x48, 0x83, 0xe0, 0xc6, 0xed, 0xe3, 0x2d, 0xde,
	0x13, 0xb7, 0xd6, 0x1d, 0x28, 0x29, 0x83, 0x8c, 0x8a, 0xd1, 0xab, 0xc3, 0x3d, 0xbe, 0xf9, 0x57,
	0x54, 0x3b, 0x9a, 0x26, 0x72, 0xd2, 0x07, 0x62, 0xaa, 0x2e, 0xb2, 0x24, 0x17, 0x49, 0x32, 0x7a,
	0xf9, 0x2b, 0xdd, 0x4c, 0xe2, 0x44, 0x93, 0x49, 0xd5, 0xa5, 0x28, 0x30, 0x0e, 0x63, 0x62, 0x1b,
	0x6b, 0x68, 0xa1, 0x3b, 0xdf, 0xac, 0xfa, 0x33, 0x8c, 0x7a, 0x01, 0xd5, 0xb5, 0x28, 0x31, 0x0e,
	0x91, 0xc3, 0xde, 0xbd, 0xd4, 0x47, 0x2d, 0x74, 0xe5, 0x66, 0xd5, 0x17, 0x18, 0x9f, 0x33, 0xa9,
	0xae, 0x44, 0xc9, 0xe1, 0xdd, 0xed, 0x90, 0xad, 0xa9, 0x8f, 0x5b, 0xe8, 0x8a, 0xfe, 0x0f, 0xe8,
	0xb5, 0x28, 0x47, 0x32, 0x69, 0xe0, 0xe4, 0xad, 0x7e, 0xfc, 0x9c, 0x1b, 0xf8, 0x9a, 0x1b, 0xf8,
	0x9e, 0x1b, 0x10, 0x77, 0x7b, 0x92, 0xb9, 0xcb, 0x07, 0xfa, 0x48, 0xf2, 0xff, 0x44, 0xbd, 0x7e,
	0xca, 0x23, 0xdb, 0xe5, 0x63, 0x0b, 0xe3, 0x69, 0x1e, 0x7a, 0xf8, 0x09, 0x00, 0x00, 0xff, 0xff,
	0xec, 0x7d, 0x68, 0x4e, 0xf7, 0x00, 0x00, 0x00,
}

func (m *Body) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Body) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Body) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Truncated {
		i--
		if m.Truncated {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.BodyType != nil {
		{
			size := m.BodyType.Size()
			i -= size
			if _, err := m.BodyType.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *Body_AsBytes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Body_AsBytes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.AsBytes != nil {
		i -= len(m.AsBytes)
		copy(dAtA[i:], m.AsBytes)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.AsBytes)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *Body_AsString) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Body_AsString) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.AsString)
	copy(dAtA[i:], m.AsString)
	i = encodeVarintCommon(dAtA, i, uint64(len(m.AsString)))
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}
func encodeVarintCommon(dAtA []byte, offset int, v uint64) int {
	offset -= sovCommon(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Body) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BodyType != nil {
		n += m.BodyType.Size()
	}
	if m.Truncated {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Body_AsBytes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AsBytes != nil {
		l = len(m.AsBytes)
		n += 1 + l + sovCommon(uint64(l))
	}
	return n
}
func (m *Body_AsString) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AsString)
	n += 1 + l + sovCommon(uint64(l))
	return n
}

func sovCommon(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCommon(x uint64) (n int) {
	return sovCommon(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Body) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: Body: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Body: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AsBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, dAtA[iNdEx:postIndex])
			m.BodyType = &Body_AsBytes{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AsString", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BodyType = &Body_AsString{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Truncated", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Truncated = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCommon
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
func skipCommon(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommon
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
					return 0, ErrIntOverflowCommon
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
					return 0, ErrIntOverflowCommon
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
				return 0, ErrInvalidLengthCommon
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCommon
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCommon
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCommon        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommon          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCommon = fmt.Errorf("proto: unexpected end of group")
)
