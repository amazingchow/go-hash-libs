package hash

// Uint64Hasher is responsible for generating unsigned-64-bit hash of provided string.
type Uint64Hasher interface {
	Sum64(string) uint64
}

// Uint32Hasher is responsible for generating unsigned-32-bit hash of provided string.
type Uint32Hasher interface {
	Sum64(string) uint64
}
