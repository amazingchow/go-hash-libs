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
