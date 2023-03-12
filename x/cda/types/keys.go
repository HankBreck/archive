package types

import (
	"encoding/binary"
)

var (
	// ModuleName defines the module name
	ModuleName = "cda"
	// StoreKey defines the primary module store key
	StoreKey = ModuleName
	// RouterKey is the message route for slashing
	RouterKey = ModuleName
	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName
	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_cda"

	// CDAKey defines the prefix for storing CDA objects
	CDAKey = "CDA-value-"
	// CDACountKey defines the prefix for storing the current number of CDA objects
	CDACountKey = "CDA-count-"
	// CDASignerKey defines the prefix for storing all ids of CDAs that a signer is a part of.
	// It is a subprefix that requires the signer's identity ID be appended to the end.
	//
	// e.g. "CDA-signer-10-" --> {1: 0x1, 5: 0x1, 6: 0x1, 11: 0x1}
	CDASignerKey = []byte("CDA-signer-")
	// CDAApprovalKey defines the prefix for storing the approvals for a specific CDA.
	// It is intended to be suffixed with the CDA's id
	//
	// e.g. "CDA-approval-1"
	CDAApprovalKey = "CDA-approval-"
	// CDAMetadataKey defines the prefix for storing the signing metadata for a specific CDA.
	// It is intended to be suffixed with the CDA's id
	//
	// e.g. "CDA-approval-1"
	CDAMetadataKey = "CDA-metadata-"
	// ContractKey defines the prefix under which Contract objects are stored
	ContractKey = "contracts-"
	// ContractCountKey defines the prefix under which the current count contracts is stored.
	// This is used to determine the next ID a contract should be stored under.
	ContractCountKey = "contract_count-"
	// SigningDataKey defines the prefix under which signing data is stored
	SigningDataKey = "signing_data-"

	// Key for the CDA ID attribute in an event
	AttributeKeyCdaId = "cda-id"
	// Key for the contract ID attribute in an event
	AttributeKeyContractId = "contract-id"

	// WitnessApprovalKey specifies the length 1 key used to store CDA
	// approvals from the CDA's witness
	WitnessApprovalKey = []byte{0x01}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func SignerCdaStoreKey(signerId uint64) []byte {
	bzSignerId := make([]byte, 8)
	binary.BigEndian.PutUint64(bzSignerId, signerId)
	return append(CDASignerKey, bzSignerId...)
}
