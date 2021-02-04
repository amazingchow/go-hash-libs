package hash

const (
	// See https://en.wikipedia.org/wiki/Fowler–Noll–Vo_hash_function#FNV_hash_parameters
	offset32 = 2166136261
	prime32  = 16777619
	offset64 = 14695981039346656037
	prime64  = 1099511628211
)

func NewFnv32() Uint32Hasher {
	return fnv32{}
}

type fnv32 struct{}

func (h fnv32) Sum32(key string) uint32 {
	var hash uint32 = offset32
	for i := 0; i < len(key); i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}

func NewFnv32a() Uint32Hasher {
	return fnv32a{}
}

type fnv32a struct{}

func (h fnv32a) Sum32(key string) uint32 {
	var hash uint32 = offset32
	for i := 0; i < len(key); i++ {
		hash ^= uint32(key[i])
		hash *= prime32
	}
	return hash
}

func NewFnv64() Uint64Hasher {
	return fnv64{}
}

type fnv64 struct{}

func (h fnv64) Sum64(key string) uint64 {
	var hash uint64 = offset64
	for i := 0; i < len(key); i++ {
		hash *= prime64
		hash ^= uint64(key[i])
	}
	return hash
}

func NewFnv64a() Uint64Hasher {
	return fnv64a{}
}

type fnv64a struct{}

func (h fnv64a) Sum64(key string) uint64 {
	var hash uint64 = offset64
	for i := 0; i < len(key); i++ {
		hash ^= uint64(key[i])
		hash *= prime64
	}
	return hash
}

// More info: http://www.isthe.com/chongo/tech/comp/fnv/index.html#FNV-source
