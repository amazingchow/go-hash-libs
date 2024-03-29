package hash

const (
	// See https://en.wikipedia.org/wiki/Fowler–Noll–Vo_hash_function#FNV_hash_parameters
	_Offset32 = 2166136261
	_Prime32  = 16777619
	_Offset64 = 14695981039346656037
	_Prime64  = 1099511628211
)

func NewFnv32() Uint32Hasher {
	return fnv32{}
}

type fnv32 struct{}

func (f fnv32) Sum32(key string) uint32 {
	var hash uint32 = _Offset32
	for i := 0; i < len(key); i++ {
		hash *= _Prime32
		hash ^= uint32(key[i])
	}
	return hash
}

func NewFnv32a() Uint32Hasher {
	return fnv32a{}
}

type fnv32a struct{}

func (f fnv32a) Sum32(key string) uint32 {
	var hash uint32 = _Offset32
	for i := 0; i < len(key); i++ {
		hash ^= uint32(key[i])
		hash *= _Prime32
	}
	return hash
}

func NewFnv64() Uint64Hasher {
	return fnv64{}
}

type fnv64 struct{}

func (f fnv64) Sum64(key string) uint64 {
	var hash uint64 = _Offset64
	for i := 0; i < len(key); i++ {
		hash *= _Prime64
		hash ^= uint64(key[i])
	}
	return hash
}

func NewFnv64a() Uint64Hasher {
	return fnv64a{}
}

type fnv64a struct{}

func (f fnv64a) Sum64(key string) uint64 {
	var hash uint64 = _Offset64
	for i := 0; i < len(key); i++ {
		hash ^= uint64(key[i])
		hash *= _Prime64
	}
	return hash
}

// More info: http://www.isthe.com/chongo/tech/comp/fnv/index.html#FNV-source
