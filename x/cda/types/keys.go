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
	// e.g. "CDA-owner-10-" --> {1: 0x1, 5: 0x1, 6: 0x1, 11: 0x1}
	CDASignerKey = []byte("CDA-signer-")

	// CDASignerCountKey defines the prefix for storing the current number of CDAs an account
	// is an owner of.
	// It is a subprefix that requires the owner's account be appended to the end.
	//
	// e.g. "CDA-owner-count-10-" --> 4
	// TODO: Clean up
	// CDASignerCountKey = []byte("CDA-signer-count-")

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

	// Key for the CDA attribute in a cosmos event
	AttributeKeyCdaId = "cda-id"

	// ContractKey defines the prefix under which Contract objects are stored
	ContractKey = "contracts-"

	// ContractCountKey defines the prefix under which the current count contracts is stored.
	// This is used to determine the next ID a contract should be stored under.
	ContractCountKey = "contract_count-"

	// SigningDataKey defines the prefix under which signing data is stored
	SigningDataKey = "signing_data-"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func SignerCdaStoreKey(signerId uint64) []byte {
	bzSignerId := make([]byte, 8)
	binary.BigEndian.PutUint64(bzSignerId, signerId)
	return append(CDASignerKey, bzSignerId...)
}
