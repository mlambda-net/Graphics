package lib

/*
#cgo LDFLAGS: -framework opencl ${SRCDIR}/libmath.a -ldl
#include "./math.h"
*/
import "C"
import "unsafe"

func Sum(items []int32) int32 {
	p := unsafe.Pointer(&items[0])
	x := C.sum((*C.int)(p), C.int(len(items)))
	return int32(x)
}
