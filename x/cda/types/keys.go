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

	/// Both CDA keys will be appended with the creator's address via CDA-value-{wallet_addr}-
	// CDA Key
	CDAKey = "CDA-value-"
	// CDA Count Key
	CDACountKey = "CDA-count-"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
