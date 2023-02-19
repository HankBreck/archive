package types

import "encoding/binary"

const (
	// ModuleName defines the module name
	ModuleName = "identity"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_identity"

	// IssuerKey defines the prefix under which Contract objects are stored
	IssuerKey = "id-issuers-"

	// CertificateKey defines the prefix under which Certificate objects are stored
	CertificateKey = "id-certs-"

	// CertificateCountKey defines the prefix under which the current count of certificates is stored.
	// This is used to determine the next ID a Certificate should be stored under.
	CertificateCountKey = "id-cert-count-"

	// MembershipKey defines the prefix under which Membership lists are stored
	MembershipKey = "id-memberships-"

	// OperatorKey defines the prefix under which Membership lists are stored
	OperatorKey = "id-operators-"

	// FreezeKey defines the prefix under which frozen identity IDs are stored
	FreezeKey = "id-frozen-"
)

func MembershipKeyPrefix(id uint64, isPending bool) []byte {
	bzId := make([]byte, 8)
	binary.BigEndian.PutUint64(bzId, id)

	var prefix []byte
	if isPending {
		prefix = append([]byte(MembershipKey), byte(0))
	} else {
		prefix = append([]byte(MembershipKey), byte(1))
	}

	return append(prefix, bzId...)
}

func OperatorKeyPrefix(id uint64) []byte {
	bzId := make([]byte, 8)
	binary.BigEndian.PutUint64(bzId, id)
	return append([]byte(OperatorKey), bzId...)
}
