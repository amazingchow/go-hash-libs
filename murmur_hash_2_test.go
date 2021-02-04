package hash

import (
	"testing"

	"github.com/amazingchow/photon-dance-hash-lib/cgo"
)

func TestMurmurHash2_32(t *testing.T) {
	h := NewMurmurHash2_32()
	if h.Sum32(text) != cgo.MurmurHash2(text) {
		t.Fatal("implementation of MurmurHash2_32 is invalid")
	}
}

func TestMurmurHash2_64(t *testing.T) {
	h := NewMurmurHash2_64()
	if h.Sum64(text) != cgo.MurmurHash64A(text) {
		// TODO: not worker under Linux 5.4.0-26-generic + Linux Mint 20 x86-64
		t.Fatal("implementation of NewMurmurHash2_64 is invalid")
	}
}
