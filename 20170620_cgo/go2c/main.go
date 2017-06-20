package main

// #cgo LDFLAGS: ./libfizzbuzz.a
// #include <fizzbuzz.h>
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	buf := [16]byte{}
	for i := 0; i < 18; i++ {
		len := C.fizzbuzz((C.int)(i), (*C.char)(unsafe.Pointer(&buf[0])), 16)
		fmt.Println(string(buf[:len]))
	}
}
