package hash

import (
	"testing"
)

func BenchmarkFnv32(b *testing.B) {
	h := NewFnv32()
	for i := 0; i < b.N; i++ {
		h.Sum32(text)
	}
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

func BenchmarkStdlibFnv64a(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stdlibFnv64a(text)
	}
}
