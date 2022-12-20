package types

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

	// IssuerCountKey defines the prefix under which the current count of issuers is stored.
	// This is used to determine the next ID an issuer should be stored under.
	IssuerCountKey = "id-issuer-count-"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
