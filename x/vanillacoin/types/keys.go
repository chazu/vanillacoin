package types

const (
	// ModuleName defines the module name
	ModuleName = "vanillacoin"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_vanillacoin"
)

var (
	ParamsKey = []byte("p_vanillacoin")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
