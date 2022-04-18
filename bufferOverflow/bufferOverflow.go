package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"unsafe"
)

func win() {
	fmt.Println("{<G0V3RF10W!>}")
}

// initialize the reader outside of the main function to simplify POC development, as
// there are less local variables on the stack.
var reader = bufio.NewReader(os.Stdin)

type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

func main() {
	// this is a harmless buffer, containing some harmless data
	harmlessData := [8]byte{'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A'}

	// create a slice of length 512 byte, but assign the address of the harmless data as
	// its buffer. Use the reflect.SliceHeader to change the slice
	confusedSlice := make([]byte, 512)
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&confusedSlice))
	harmlessDataAddress := uintptr(unsafe.Pointer(&(harmlessData[0])))
	sliceHeader.Data = harmlessDataAddress

	// now read into the confused slice from STDIN. This is not quite as bad as a gets()
	// call in C, but almost. The function will read up to 512 byte, but the underlying
	// buffer is only 8 bytes. This function is pretty much the complete vulnerability
	_, _ = reader.Read(confusedSlice)
}
