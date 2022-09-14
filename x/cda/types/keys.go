package types

const (
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

	// CDAOwnerKey defines the prefix for storing all ids of CDAs owned by an account.
	// It is a subprefix that requires the owner's account be appended to the end.
	//
	// e.g. "CDA-owner-archive1ps3rtvcqw3p9megamtg8mrq3nn7fvzw2de6e62-" --> [1, 5, 6, 11]
	CDAOwnerKey = "CDA-owner-"

	// CDAOwnerCountKey defines the prefix for storing the current number of CDAs an account
	// is an owner of.
	// It is a subprefix that requires the owner's account be appended to the end.
	//
	// e.g. "CDA-owner-count-archive1ps3rtvcqw3p9megamtg8mrq3nn7fvzw2de6e62-" --> 4
	CDAOwnerCountKey = "CDA-owner-count-"

	//
	AttributeKeyCdaId = "cda-id"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
