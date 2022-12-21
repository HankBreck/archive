// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: archive/identity/certificate.proto

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

type Certificate struct {
	Id                uint64            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IssuerAddress     string            `protobuf:"bytes,2,opt,name=issuer_address,json=issuerAddress,proto3" json:"issuer_address,omitempty"`
	Salt              string            `protobuf:"bytes,3,opt,name=salt,proto3" json:"salt,omitempty"`
	MetadataSchemaUri string            `protobuf:"bytes,4,opt,name=metadata_schema_uri,json=metadataSchemaUri,proto3" json:"metadata_schema_uri,omitempty"`
	Hashes            map[string]string `protobuf:"bytes,5,rep,name=hashes,proto3" json:"hashes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *Certificate) Reset()         { *m = Certificate{} }
func (m *Certificate) String() string { return proto.CompactTextString(m) }
func (*Certificate) ProtoMessage()    {}
func (*Certificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7da3374219bca8c, []int{0}
}
func (m *Certificate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Certificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Certificate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Certificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Certificate.Merge(m, src)
}
func (m *Certificate) XXX_Size() int {
	return m.Size()
}
func (m *Certificate) XXX_DiscardUnknown() {
	xxx_messageInfo_Certificate.DiscardUnknown(m)
}

var xxx_messageInfo_Certificate proto.InternalMessageInfo

func (m *Certificate) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Certificate) GetIssuerAddress() string {
	if m != nil {
		return m.IssuerAddress
	}
	return ""
}

func (m *Certificate) GetSalt() string {
	if m != nil {
		return m.Salt
	}
	return ""
}

func (m *Certificate) GetMetadataSchemaUri() string {
	if m != nil {
		return m.MetadataSchemaUri
	}
	return ""
}

func (m *Certificate) GetHashes() map[string]string {
	if m != nil {
		return m.Hashes
	}
	return nil
}

func init() {
	proto.RegisterType((*Certificate)(nil), "archive.identity.Certificate")
	proto.RegisterMapType((map[string]string)(nil), "archive.identity.Certificate.HashesEntry")
}

func init() {
	proto.RegisterFile("archive/identity/certificate.proto", fileDescriptor_b7da3374219bca8c)
}

var fileDescriptor_b7da3374219bca8c = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xcf, 0x6a, 0x3a, 0x31,
	0x10, 0xc7, 0xcd, 0xfa, 0x07, 0x8c, 0x28, 0x9a, 0x9f, 0x3f, 0x08, 0x1e, 0xa2, 0x78, 0xb2, 0x97,
	0x15, 0xec, 0xc5, 0xf6, 0xa6, 0x45, 0x68, 0x0f, 0xbd, 0xa4, 0xf4, 0xd2, 0x8b, 0xa4, 0x9b, 0xa9,
	0x1b, 0xaa, 0xae, 0x24, 0x51, 0xba, 0x6f, 0xd1, 0x47, 0xe9, 0x63, 0xf4, 0xe8, 0xb1, 0x27, 0x29,
	0xeb, 0x8b, 0x14, 0xb3, 0x8a, 0x8b, 0xb7, 0x99, 0xf9, 0x7c, 0x26, 0xf0, 0xcd, 0xe0, 0xae, 0xd0,
	0x41, 0xa8, 0x36, 0xd0, 0x57, 0x12, 0x96, 0x56, 0xd9, 0xb8, 0x1f, 0x80, 0xb6, 0xea, 0x4d, 0x05,
	0xc2, 0x82, 0xbf, 0xd2, 0x91, 0x8d, 0x48, 0xfd, 0xe8, 0xf8, 0x27, 0xa7, 0xd5, 0x9c, 0x45, 0xb3,
	0xc8, 0xc1, 0xfe, 0xa1, 0x4a, 0xbd, 0xee, 0x97, 0x87, 0x2b, 0x77, 0xe7, 0x6d, 0x52, 0xc3, 0x9e,
	0x92, 0x14, 0x75, 0x50, 0xaf, 0xc0, 0x3d, 0x25, 0xc9, 0x10, 0xd7, 0x94, 0x31, 0x6b, 0xd0, 0x53,
	0x21, 0xa5, 0x06, 0x63, 0xa8, 0xd7, 0x41, 0xbd, 0xf2, 0xb8, 0x91, 0xec, 0xda, 0xd5, 0x07, 0x47,
	0x46, 0x29, 0xe0, 0x55, 0x95, 0x6d, 0x09, 0xc1, 0x05, 0x23, 0xe6, 0x96, 0xe6, 0x0f, 0x3e, 0x77,
	0x35, 0x99, 0xe0, 0x7f, 0x0b, 0xb0, 0x42, 0x0a, 0x2b, 0xa6, 0x26, 0x08, 0x61, 0x21, 0xa6, 0x6b,
	0xad, 0x68, 0xc1, 0x3d, 0xf9, 0x3f, 0xd9, 0xb5, 0x1b, 0x8f, 0x47, 0xfc, 0xe4, 0xe8, 0xb3, 0x56,
	0xbc, 0xb1, 0xb8, 0x1c, 0x91, 0x11, 0x2e, 0x85, 0xc2, 0x84, 0x60, 0x68, 0xb1, 0x93, 0xef, 0x55,
	0x06, 0x57, 0xfe, 0x65, 0x5a, 0x3f, 0x93, 0xc9, 0xbf, 0x77, 0xee, 0x64, 0x69, 0x75, 0xcc, 0x8f,
	0x8b, 0xad, 0x1b, 0x5c, 0xc9, 0x8c, 0x49, 0x1d, 0xe7, 0xdf, 0x21, 0x76, 0xb9, 0xcb, 0xfc, 0x50,
	0x92, 0x26, 0x2e, 0x6e, 0xc4, 0x7c, 0x0d, 0x69, 0x5e, 0x9e, 0x36, 0xb7, 0xde, 0x10, 0x8d, 0x07,
	0xdf, 0x09, 0x43, 0xdb, 0x84, 0xa1, 0xdf, 0x84, 0xa1, 0xcf, 0x3d, 0xcb, 0x6d, 0xf7, 0x2c, 0xf7,
	0xb3, 0x67, 0xb9, 0x17, 0x7a, 0x3a, 0xcc, 0xc7, 0xf9, 0x34, 0x36, 0x5e, 0x81, 0x79, 0x2d, 0xb9,
	0xdf, 0xbe, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x31, 0xbd, 0xda, 0x3a, 0xbb, 0x01, 0x00, 0x00,
}

func (m *Certificate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Certificate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Certificate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Hashes) > 0 {
		for k := range m.Hashes {
			v := m.Hashes[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintCertificate(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintCertificate(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintCertificate(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.MetadataSchemaUri) > 0 {
		i -= len(m.MetadataSchemaUri)
		copy(dAtA[i:], m.MetadataSchemaUri)
		i = encodeVarintCertificate(dAtA, i, uint64(len(m.MetadataSchemaUri)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Salt) > 0 {
		i -= len(m.Salt)
		copy(dAtA[i:], m.Salt)
		i = encodeVarintCertificate(dAtA, i, uint64(len(m.Salt)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.IssuerAddress) > 0 {
		i -= len(m.IssuerAddress)
		copy(dAtA[i:], m.IssuerAddress)
		i = encodeVarintCertificate(dAtA, i, uint64(len(m.IssuerAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintCertificate(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCertificate(dAtA []byte, offset int, v uint64) int {
	offset -= sovCertificate(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Certificate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovCertificate(uint64(m.Id))
	}
	l = len(m.IssuerAddress)
	if l > 0 {
		n += 1 + l + sovCertificate(uint64(l))
	}
	l = len(m.Salt)
	if l > 0 {
		n += 1 + l + sovCertificate(uint64(l))
	}
	l = len(m.MetadataSchemaUri)
	if l > 0 {
		n += 1 + l + sovCertificate(uint64(l))
	}
	if len(m.Hashes) > 0 {
		for k, v := range m.Hashes {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovCertificate(uint64(len(k))) + 1 + len(v) + sovCertificate(uint64(len(v)))
			n += mapEntrySize + 1 + sovCertificate(uint64(mapEntrySize))
		}
	}
	return n
}

func sovCertificate(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCertificate(x uint64) (n int) {
	return sovCertificate(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Certificate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCertificate
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
			return fmt.Errorf("proto: Certificate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Certificate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificate
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
				return fmt.Errorf("proto: wrong wireType = %d for field IssuerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificate
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
				return ErrInvalidLengthCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IssuerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Salt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificate
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
				return ErrInvalidLengthCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Salt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetadataSchemaUri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificate
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
				return ErrInvalidLengthCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MetadataSchemaUri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hashes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificate
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
				return ErrInvalidLengthCertificate
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Hashes == nil {
				m.Hashes = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCertificate
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCertificate
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthCertificate
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthCertificate
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCertificate
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthCertificate
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthCertificate
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipCertificate(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthCertificate
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Hashes[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCertificate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCertificate
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
func skipCertificate(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCertificate
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
					return 0, ErrIntOverflowCertificate
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
					return 0, ErrIntOverflowCertificate
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
				return 0, ErrInvalidLengthCertificate
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCertificate
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCertificate
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCertificate        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCertificate          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCertificate = fmt.Errorf("proto: unexpected end of group")
)
