package types

const (
	// ModuleName defines the module name
	ModuleName = "contractregistry"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_contractregistry"

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
