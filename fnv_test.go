package hash

import (
	"hash/fnv"
	"testing"
)

var text = "Hello World!!!"

func stdlibFnv32(key string) uint32 {
	h := fnv.New32()
	h.Write([]byte(key))
	return h.Sum32()
}

func TestFnv32(t *testing.T) {
	h := NewFnv32()
	if h.Sum32(text) != stdlibFnv32(text) {
		t.Fatal("implementation of Fnv32 is invalid")
	}
}

func stdlibFnv32a(key string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32()
}

func TestFnv32a(t *testing.T) {
	h := NewFnv32a()
	if h.Sum32(text) != stdlibFnv32a(text) {
		t.Fatal("implementation of Fnv32a is invalid")
	}
}

func stdlibFnv64(key string) uint64 {
	h := fnv.New64()
	h.Write([]byte(key))
	return h.Sum64()
}

func TestFnv64(t *testing.T) {
	h := NewFnv64()
	if h.Sum64(text) != stdlibFnv64(text) {
		t.Fatal("implementation of Fnv64 is invalid")
	}
}

func stdlibFnv64a(key string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(key))
	return h.Sum64()
}

func TestFnv64a(t *testing.T) {
	h := NewFnv64a()
	if h.Sum64(text) != stdlibFnv64a(text) {
		t.Fatal("implementation of Fnv64a is invalid")
	}
}
