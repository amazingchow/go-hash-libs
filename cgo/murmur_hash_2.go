package cgo

/*
#cgo LDFLAGS: -lmurmurhash2
#include <murmurhash/MurmurHash2.h>

#include <stdlib.h>
#include <stdint.h>
*/
import "C"

import "unsafe"

func MurmurHash2(key string) uint32 {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))

	var x C.uint32_t = C.MurmurHash2(unsafe.Pointer(ckey), C.int(len(key)), C.uint32_t(0xbc9f1d34))

	return uint32(x)
}
