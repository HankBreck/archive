// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: archive/identity/issuer.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type Issuer struct {
	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Creator     string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	MoreInfoUri string `protobuf:"bytes,4,opt,name=more_info_uri,json=moreInfoUri,proto3" json:"more_info_uri,omitempty"`
	Cost        uint64 `protobuf:"varint,5,opt,name=cost,proto3" json:"cost,omitempty"`
}

func (m *Issuer) Reset()         { *m = Issuer{} }
func (m *Issuer) String() string { return proto.CompactTextString(m) }
func (*Issuer) ProtoMessage()    {}
func (*Issuer) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d0df9580870907e, []int{0}
}
func (m *Issuer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Issuer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Issuer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Issuer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Issuer.Merge(m, src)
}
func (m *Issuer) XXX_Size() int {
	return m.Size()
}
func (m *Issuer) XXX_DiscardUnknown() {
	xxx_messageInfo_Issuer.DiscardUnknown(m)
}

var xxx_messageInfo_Issuer proto.InternalMessageInfo

func (m *Issuer) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Issuer) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Issuer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Issuer) GetMoreInfoUri() string {
	if m != nil {
		return m.MoreInfoUri
	}
	return ""
}

func (m *Issuer) GetCost() uint64 {
	if m != nil {
		return m.Cost
	}
	return 0
}

func init() {
	proto.RegisterType((*Issuer)(nil), "archive.identity.Issuer")
}

func init() { proto.RegisterFile("archive/identity/issuer.proto", fileDescriptor_5d0df9580870907e) }

var fileDescriptor_5d0df9580870907e = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0x2c, 0x4a, 0xce,
	0xc8, 0x2c, 0x4b, 0xd5, 0xcf, 0x4c, 0x49, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9, 0xd4, 0xcf, 0x2c, 0x2e,
	0x2e, 0x4d, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x80, 0x4a, 0xeb, 0xc1, 0xa4,
	0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x92, 0xfa, 0x20, 0x16, 0x44, 0x9d, 0x52, 0x3f, 0x23,
	0x17, 0x9b, 0x27, 0x58, 0xa3, 0x10, 0x1f, 0x17, 0x53, 0x66, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06,
	0x4b, 0x10, 0x53, 0x66, 0x8a, 0x90, 0x04, 0x17, 0x7b, 0x72, 0x51, 0x6a, 0x62, 0x49, 0x7e, 0x91,
	0x04, 0x93, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c, 0x2b, 0x24, 0xc4, 0xc5, 0x92, 0x97, 0x98, 0x9b,
	0x2a, 0xc1, 0x0c, 0x16, 0x06, 0xb3, 0x85, 0x8c, 0xb9, 0x78, 0x73, 0xf3, 0x8b, 0x52, 0xe3, 0x33,
	0xf3, 0xd2, 0xf2, 0xe3, 0x4b, 0x8b, 0x32, 0x25, 0x58, 0x40, 0x92, 0x4e, 0xfc, 0x8f, 0xee, 0xc9,
	0x73, 0xfb, 0xe6, 0x17, 0xa5, 0x7a, 0xe6, 0xa5, 0xe5, 0x87, 0x16, 0x65, 0x06, 0x71, 0xe7, 0x22,
	0x38, 0x20, 0x83, 0x92, 0xf3, 0x8b, 0x4b, 0x24, 0x58, 0xc1, 0x96, 0x82, 0xd9, 0x4e, 0x46, 0x27,
	0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c,
	0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x25, 0x01, 0xf3, 0x72, 0x05, 0xc2, 0xd3,
	0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0xcf, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x3d, 0x89, 0x39, 0x05, 0x15, 0x01, 0x00, 0x00,
}

func (m *Issuer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Issuer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Issuer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Cost != 0 {
		i = encodeVarintIssuer(dAtA, i, uint64(m.Cost))
		i--
		dAtA[i] = 0x28
	}
	if len(m.MoreInfoUri) > 0 {
		i -= len(m.MoreInfoUri)
		copy(dAtA[i:], m.MoreInfoUri)
		i = encodeVarintIssuer(dAtA, i, uint64(len(m.MoreInfoUri)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintIssuer(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintIssuer(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintIssuer(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintIssuer(dAtA []byte, offset int, v uint64) int {
	offset -= sovIssuer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Issuer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovIssuer(uint64(m.Id))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovIssuer(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovIssuer(uint64(l))
	}
	l = len(m.MoreInfoUri)
	if l > 0 {
		n += 1 + l + sovIssuer(uint64(l))
	}
	if m.Cost != 0 {
		n += 1 + sovIssuer(uint64(m.Cost))
	}
	return n
}

func sovIssuer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIssuer(x uint64) (n int) {
	return sovIssuer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Issuer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIssuer
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
			return fmt.Errorf("proto: Issuer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Issuer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssuer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssuer
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
				return ErrInvalidLengthIssuer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIssuer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssuer
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
				return ErrInvalidLengthIssuer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIssuer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MoreInfoUri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssuer
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
				return ErrInvalidLengthIssuer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIssuer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MoreInfoUri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cost", wireType)
			}
			m.Cost = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssuer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Cost |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIssuer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIssuer
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
func skipIssuer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIssuer
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
					return 0, ErrIntOverflowIssuer
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
					return 0, ErrIntOverflowIssuer
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
				return 0, ErrInvalidLengthIssuer
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIssuer
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIssuer
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIssuer        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIssuer          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIssuer = fmt.Errorf("proto: unexpected end of group")
)
