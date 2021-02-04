package hash

import (
	"hash/fnv"
	"testing"
)

var text = "Hello World!!!"

func BenchmarkFnv32(b *testing.B) {
	h := NewFnv32()
	for i := 0; i < b.N; i++ {
		h.Sum32(text)
	}
}

func stdlibFnv32(key string) uint32 {
	h := fnv.New32()
	h.Write([]byte(key))
	return h.Sum32()
}

func BenchmarkStdlibFnv32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stdlibFnv32(text)
	}
}

func BenchmarkFnv32a(b *testing.B) {
	h := NewFnv32a()
	for i := 0; i < b.N; i++ {
		h.Sum32(text)
	}
}

func stdlibFnv32a(key string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32()
}

func BenchmarkStdlibFnv32a(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stdlibFnv32a(text)
	}
}

func BenchmarkFnv64(b *testing.B) {
	h := NewFnv64()
	for i := 0; i < b.N; i++ {
		h.Sum64(text)
	}
}

func stdlibFnv64(key string) uint64 {
	h := fnv.New64()
	h.Write([]byte(key))
	return h.Sum64()
}

func BenchmarkStdlibFnv64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stdlibFnv64(text)
	}
}

func BenchmarkFnv64a(b *testing.B) {
	h := NewFnv64a()
	for i := 0; i < b.N; i++ {
		h.Sum64(text)
	}
}

func stdlibFnv64a(key string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(key))
	return h.Sum64()
}

func BenchmarkStdlibFnv64a(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stdlibFnv64a(text)
	}
}
