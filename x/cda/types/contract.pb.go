// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: archive/cda/contract.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

type ContactMethod int32

const (
	ContactMethod_Phone ContactMethod = 0
	ContactMethod_Email ContactMethod = 1
)

var ContactMethod_name = map[int32]string{
	0: "Phone",
	1: "Email",
}

var ContactMethod_value = map[string]int32{
	"Phone": 0,
	"Email": 1,
}

func (x ContactMethod) String() string {
	return proto.EnumName(ContactMethod_name, int32(x))
}

func (ContactMethod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f9ad2df853cc13a9, []int{0}
}

type Contract struct {
	Id                uint64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Description       string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Authors           []string     `protobuf:"bytes,3,rep,name=authors,proto3" json:"authors,omitempty"`
	ContactInfo       *ContactInfo `protobuf:"bytes,4,opt,name=contact_info,json=contactInfo,proto3" json:"contact_info,omitempty"`
	MoreInfoUri       string       `protobuf:"bytes,5,opt,name=more_info_uri,json=moreInfoUri,proto3" json:"more_info_uri,omitempty"`
	TemplateUri       string       `protobuf:"bytes,6,opt,name=template_uri,json=templateUri,proto3" json:"template_uri,omitempty"`
	TemplateSchemaUri string       `protobuf:"bytes,7,opt,name=template_schema_uri,json=templateSchemaUri,proto3" json:"template_schema_uri,omitempty"`
}

func (m *Contract) Reset()         { *m = Contract{} }
func (m *Contract) String() string { return proto.CompactTextString(m) }
func (*Contract) ProtoMessage()    {}
func (*Contract) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9ad2df853cc13a9, []int{0}
}
func (m *Contract) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Contract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Contract.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Contract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contract.Merge(m, src)
}
func (m *Contract) XXX_Size() int {
	return m.Size()
}
func (m *Contract) XXX_DiscardUnknown() {
	xxx_messageInfo_Contract.DiscardUnknown(m)
}

var xxx_messageInfo_Contract proto.InternalMessageInfo

func (m *Contract) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Contract) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Contract) GetAuthors() []string {
	if m != nil {
		return m.Authors
	}
	return nil
}

func (m *Contract) GetContactInfo() *ContactInfo {
	if m != nil {
		return m.ContactInfo
	}
	return nil
}

func (m *Contract) GetMoreInfoUri() string {
	if m != nil {
		return m.MoreInfoUri
	}
	return ""
}

func (m *Contract) GetTemplateUri() string {
	if m != nil {
		return m.TemplateUri
	}
	return ""
}

func (m *Contract) GetTemplateSchemaUri() string {
	if m != nil {
		return m.TemplateSchemaUri
	}
	return ""
}

type ContactInfo struct {
	Method ContactMethod `protobuf:"varint,1,opt,name=method,proto3,enum=archive.cda.ContactMethod" json:"method,omitempty"`
	Value  string        `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *ContactInfo) Reset()         { *m = ContactInfo{} }
func (m *ContactInfo) String() string { return proto.CompactTextString(m) }
func (*ContactInfo) ProtoMessage()    {}
func (*ContactInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9ad2df853cc13a9, []int{1}
}
func (m *ContactInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContactInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContactInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContactInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactInfo.Merge(m, src)
}
func (m *ContactInfo) XXX_Size() int {
	return m.Size()
}
func (m *ContactInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ContactInfo proto.InternalMessageInfo

func (m *ContactInfo) GetMethod() ContactMethod {
	if m != nil {
		return m.Method
	}
	return ContactMethod_Phone
}

func (m *ContactInfo) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterEnum("archive.cda.ContactMethod", ContactMethod_name, ContactMethod_value)
	proto.RegisterType((*Contract)(nil), "archive.cda.Contract")
	proto.RegisterType((*ContactInfo)(nil), "archive.cda.ContactInfo")
}

func init() { proto.RegisterFile("archive/cda/contract.proto", fileDescriptor_f9ad2df853cc13a9) }

var fileDescriptor_f9ad2df853cc13a9 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xc1, 0xaa, 0xd3, 0x40,
	0x14, 0x86, 0x33, 0xb9, 0xb7, 0xbd, 0x76, 0x72, 0xef, 0xb5, 0x1d, 0x2b, 0x84, 0x2e, 0xd2, 0x50,
	0x10, 0x82, 0x8b, 0x04, 0xd2, 0x37, 0x48, 0x29, 0x28, 0x58, 0x90, 0xa8, 0x08, 0x6e, 0xca, 0x74,
	0x32, 0x6d, 0x86, 0x36, 0x99, 0x30, 0x99, 0x14, 0x7d, 0x0b, 0x1f, 0xcb, 0x65, 0x17, 0x2e, 0x5c,
	0x15, 0x49, 0x5f, 0x44, 0x32, 0x49, 0x4d, 0x94, 0xbb, 0x9b, 0xff, 0x3f, 0xff, 0x77, 0x38, 0x73,
	0x38, 0x70, 0x82, 0x05, 0x89, 0xd9, 0x91, 0x7a, 0x24, 0xc2, 0x1e, 0xe1, 0xa9, 0x14, 0x98, 0x48,
	0x37, 0x13, 0x5c, 0x72, 0x64, 0x34, 0x35, 0x97, 0x44, 0x78, 0x32, 0xde, 0xf1, 0x1d, 0x57, 0xbe,
	0x57, 0xbd, 0xea, 0xc8, 0xec, 0xa7, 0x0e, 0x9f, 0x2d, 0x1a, 0x0a, 0x3d, 0x42, 0x9d, 0x45, 0x26,
	0xb0, 0x81, 0x73, 0x1b, 0xea, 0x2c, 0x42, 0x36, 0x34, 0x22, 0x9a, 0x13, 0xc1, 0x32, 0xc9, 0x78,
	0x6a, 0xea, 0x36, 0x70, 0x06, 0x61, 0xd7, 0x42, 0x26, 0xbc, 0xc3, 0x85, 0x8c, 0xb9, 0xc8, 0xcd,
	0x1b, 0xfb, 0xc6, 0x19, 0x84, 0x57, 0x89, 0xde, 0xc1, 0xfb, 0x6a, 0x1a, 0x4c, 0xe4, 0x9a, 0xa5,
	0x5b, 0x6e, 0xde, 0xda, 0xc0, 0x31, 0x7c, 0xd3, 0xed, 0x8c, 0xe4, 0x2e, 0xea, 0xc0, 0xdb, 0x74,
	0xcb, 0x83, 0xe7, 0xe5, 0x79, 0x6a, 0x74, 0x8c, 0xd0, 0x20, 0xad, 0x40, 0x73, 0xf8, 0x90, 0x70,
	0x41, 0x55, 0xab, 0x75, 0x21, 0x98, 0xd9, 0xab, 0x66, 0xa9, 0xa1, 0x15, 0x17, 0xb4, 0x0a, 0x7d,
	0x12, 0x2c, 0x34, 0x92, 0x56, 0x20, 0x1f, 0xde, 0x4b, 0x9a, 0x64, 0x07, 0x2c, 0xa9, 0x62, 0xfa,
	0x2d, 0xf3, 0xb1, 0xf1, 0x15, 0x23, 0x5b, 0x81, 0x96, 0xf0, 0xc5, 0x5f, 0x26, 0x27, 0x31, 0x4d,
	0xb0, 0x42, 0xef, 0x14, 0xfa, 0xb2, 0x3c, 0x4f, 0x47, 0x57, 0xf4, 0x83, 0xaa, 0x56, 0x0d, 0x46,
	0xf2, 0x7f, 0x6b, 0xf6, 0x19, 0x76, 0xff, 0x82, 0x7c, 0xd8, 0x4f, 0xa8, 0x8c, 0x79, 0xbd, 0xdc,
	0x47, 0x7f, 0xf2, 0xd4, 0x1a, 0x56, 0x2a, 0x11, 0x36, 0x49, 0x34, 0x86, 0xbd, 0x23, 0x3e, 0x14,
	0xb4, 0x59, 0x7b, 0x2d, 0x5e, 0xbf, 0x82, 0x0f, 0xff, 0xc4, 0xd1, 0x00, 0xf6, 0xde, 0xc7, 0x3c,
	0xa5, 0x43, 0xad, 0x7a, 0x2e, 0x13, 0xcc, 0x0e, 0x43, 0x10, 0x04, 0x3f, 0x4a, 0x0b, 0x9c, 0x4a,
	0x0b, 0xfc, 0x2e, 0x2d, 0xf0, 0xfd, 0x62, 0x69, 0xa7, 0x8b, 0xa5, 0xfd, 0xba, 0x58, 0xda, 0x17,
	0x67, 0xc7, 0x64, 0x5c, 0x6c, 0x5c, 0xc2, 0x13, 0xef, 0x0d, 0x4e, 0xf7, 0x81, 0xa0, 0x64, 0xef,
	0x5d, 0x8f, 0xe8, 0xab, 0x3a, 0x23, 0xf9, 0x2d, 0xa3, 0xf9, 0xa6, 0xaf, 0x2e, 0x64, 0xfe, 0x27,
	0x00, 0x00, 0xff, 0xff, 0x02, 0x09, 0xee, 0x4d, 0x62, 0x02, 0x00, 0x00,
}

func (m *Contract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Contract) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Contract) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TemplateSchemaUri) > 0 {
		i -= len(m.TemplateSchemaUri)
		copy(dAtA[i:], m.TemplateSchemaUri)
		i = encodeVarintContract(dAtA, i, uint64(len(m.TemplateSchemaUri)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.TemplateUri) > 0 {
		i -= len(m.TemplateUri)
		copy(dAtA[i:], m.TemplateUri)
		i = encodeVarintContract(dAtA, i, uint64(len(m.TemplateUri)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.MoreInfoUri) > 0 {
		i -= len(m.MoreInfoUri)
		copy(dAtA[i:], m.MoreInfoUri)
		i = encodeVarintContract(dAtA, i, uint64(len(m.MoreInfoUri)))
		i--
		dAtA[i] = 0x2a
	}
	if m.ContactInfo != nil {
		{
			size, err := m.ContactInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintContract(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Authors) > 0 {
		for iNdEx := len(m.Authors) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Authors[iNdEx])
			copy(dAtA[i:], m.Authors[iNdEx])
			i = encodeVarintContract(dAtA, i, uint64(len(m.Authors[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintContract(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintContract(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ContactInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContactInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContactInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintContract(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if m.Method != 0 {
		i = encodeVarintContract(dAtA, i, uint64(m.Method))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintContract(dAtA []byte, offset int, v uint64) int {
	offset -= sovContract(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Contract) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovContract(uint64(m.Id))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovContract(uint64(l))
	}
	if len(m.Authors) > 0 {
		for _, s := range m.Authors {
			l = len(s)
			n += 1 + l + sovContract(uint64(l))
		}
	}
	if m.ContactInfo != nil {
		l = m.ContactInfo.Size()
		n += 1 + l + sovContract(uint64(l))
	}
	l = len(m.MoreInfoUri)
	if l > 0 {
		n += 1 + l + sovContract(uint64(l))
	}
	l = len(m.TemplateUri)
	if l > 0 {
		n += 1 + l + sovContract(uint64(l))
	}
	l = len(m.TemplateSchemaUri)
	if l > 0 {
		n += 1 + l + sovContract(uint64(l))
	}
	return n
}

func (m *ContactInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Method != 0 {
		n += 1 + sovContract(uint64(m.Method))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovContract(uint64(l))
	}
	return n
}

func sovContract(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozContract(x uint64) (n int) {
	return sovContract(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Contract) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContract
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
			return fmt.Errorf("proto: Contract: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Contract: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
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
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
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
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authors", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
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
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authors = append(m.Authors, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContactInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
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
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ContactInfo == nil {
				m.ContactInfo = &ContactInfo{}
			}
			if err := m.ContactInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MoreInfoUri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
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
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MoreInfoUri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TemplateUri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
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
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TemplateUri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TemplateSchemaUri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
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
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TemplateSchemaUri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContract(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthContract
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
func (m *ContactInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContract
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
			return fmt.Errorf("proto: ContactInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContactInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Method", wireType)
			}
			m.Method = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Method |= ContactMethod(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
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
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContract(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthContract
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
func skipContract(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowContract
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
					return 0, ErrIntOverflowContract
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
					return 0, ErrIntOverflowContract
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
				return 0, ErrInvalidLengthContract
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupContract
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthContract
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthContract        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowContract          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupContract = fmt.Errorf("proto: unexpected end of group")
)
