package types

import proto "github.com/gogo/protobuf/proto"

type SigningDataExtension struct {
	SigningData RawSigningData
}

func (m *SigningDataExtension) Reset()         { *m = SigningDataExtension{} }
func (m *SigningDataExtension) String() string { return proto.CompactTextString(m) }
func (*SigningDataExtension) ProtoMessage()    {}
