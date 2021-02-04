package hash

import (
	"unsafe"
)

// Forked from Austin Appleby's cpp version.
// See https://github.com/aappleby/smhasher/blob/master/src/MurmurHash2.cpp
//
//-----------------------------------------------------------------------------
// MurmurHash2 was written by Austin Appleby, and is placed in the public
// domain. The author hereby disclaims copyright to this source code.

// Note - This code makes a few assumptions about how your machine behaves -

// 1. We can read a 4-byte value from any address without crashing
// 2. sizeof(int) == 4

// And it has a few limitations -

// 1. It will not work incrementally.
// 2. It will not produce the same results on little-endian and big-endian
//    machines.

func NewMurmurHash2_32() Uint32Hasher {
	return murmurhash2_32{}
}

type murmurhash2_32 struct{}

func (f murmurhash2_32) Sum32(key string) uint32 {
	// 'm' and 'r' are mixing constants generated offline.
	// They're not really 'magic', they just happen to work well.

	const (
		m uint32 = 0x5bd1e995
		r uint   = 24
	)

	// Initialize the hash to a 'random' value.

	var (
		// seed is inspired from Jeff Dean's LevelDB
		_l   int    = len(key)
		l    uint32 = *(*uint32)(unsafe.Pointer(&_l))
		seed uint32 = 0xbc9f1d34
		h    uint32 = seed ^ l
	)

	// Mix 4 bytes at a time into the hash.

	var (
		data []byte = []byte(key)
		idx  int    = 0
	)
	for l >= 4 {
		var k uint32 = *(*uint32)(unsafe.Pointer(&data[idx]))

		k *= m
		k ^= k >> r
		k *= m

		h *= m
		h ^= k

		idx += 4
		l -= 4
	}

	// Handle the last few bytes of the input array.

	var x uint
	switch l {
	case 3:
		x = uint(data[idx+2]) << 16
		h ^= *(*uint32)(unsafe.Pointer(&x))
		fallthrough
	case 2:
		x = uint(data[idx+1]) << 8
		h ^= *(*uint32)(unsafe.Pointer(&x))
		fallthrough
	case 1:
		x = uint(data[idx])
		h ^= *(*uint32)(unsafe.Pointer(&x))
		h *= m
	}

	// Do a few final mixes of the hash to ensure the last few
	// bytes are well-incorporated.

	h ^= h >> 13
	h *= m
	h ^= h >> 15

	return h
}

func NewMurmurHash2_64() Uint64Hasher {
	return murmurhash2_64{}
}

type murmurhash2_64 struct{}

func (f murmurhash2_64) Sum64(key string) uint64 {
	const (
		m uint64 = 0xc6a4a7935bd1e995
		r uint   = 47
	)

	var (
		// seed is inspired from Jeff Dean's LevelDB
		_l   int    = len(key)
		l    uint64 = *(*uint64)(unsafe.Pointer(&_l))
		seed uint64 = 0xbc9f1d34
		h    uint64 = seed ^ (l * m)
	)

	var (
		data   []byte = []byte(key)
		idx    int    = 0
		data_e uint64 = *(*uint64)(unsafe.Pointer(&data[idx]))
		end    uint64 = *(*uint64)(unsafe.Pointer(&data[idx+int(l)/8]))
	)

	for data_e != end {
		var k uint64 = *(*uint64)(unsafe.Pointer(&data[idx]))

		k *= m
		k ^= k >> r
		k *= m

		h ^= k
		h *= m

		idx++
		data_e = *(*uint64)(unsafe.Pointer(&data[idx]))
	}

	var (
		data2 []byte = data[idx:]
	)

	var x uint
	switch l & 7 {
	case 7:
		x = uint(data2[6])
		h ^= (*(*uint64)(unsafe.Pointer(&x)) << 48)
		fallthrough
	case 6:
		x = uint(data2[5])
		h ^= (*(*uint64)(unsafe.Pointer(&x)) << 40)
		fallthrough
	case 5:
		x = uint(data2[4])
		h ^= (*(*uint64)(unsafe.Pointer(&x)) << 32)
		fallthrough
	case 4:
		x = uint(data2[3])
		h ^= (*(*uint64)(unsafe.Pointer(&x)) << 24)
		fallthrough
	case 3:
		x = uint(data2[2])
		h ^= (*(*uint64)(unsafe.Pointer(&x)) << 16)
		fallthrough
	case 2:
		x = uint(data2[1])
		h ^= (*(*uint64)(unsafe.Pointer(&x)) << 8)
		fallthrough
	case 1:
		x = uint(data2[0])
		h ^= *(*uint64)(unsafe.Pointer(&x))
		h *= m
	}

	h ^= h >> r
	h *= m
	h ^= h >> r

	return h
}
