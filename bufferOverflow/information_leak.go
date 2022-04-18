package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"time"
	"unsafe"
)

/*

This buffer overflow program will take input when running the program or
via piped input. This will create a race condition against the garbage collector
using the unsafe package.

*/

func main() {
	go heapHeapHeap()

	readAndHaveFun()
}

func unsafeStringToBytes(s *string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(s))
	sliceHeader := &reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}

	// CHANGE:
	time.Sleep(1 * time.Nanosecond)

	return *(*[]byte)(unsafe.Pointer(sliceHeader))
}

func readAndHaveFun() {
	reader := bufio.NewReader(os.Stdin)
	count := 1
	var firstChar byte

	fmt.Println("Taking strings to buffer. Takes many 'iterations' of input :)")
	fmt.Println("Automating input is encouraged ;)")
	for {
		s, _ := reader.ReadString('\n')
		if len(s) == 0 {
			continue
		}
		firstChar = s[0]

		// Overflow after iterations
		bytes := unsafeStringToBytes(&s)

		_, _ = reader.ReadString('\n')

		if len(bytes) > 0 && bytes[0] != firstChar {
			fmt.Printf("{0V3R_Fl3W} after %d iterations\n", count)
			os.Exit(0)
		}
		//fmt.Printf("%s\n", s)
		count++
	}
}

func heapHeapHeap() {
	var a *[]byte
	for {
		tmp := make([]byte, 1000000, 1000000)
		a = &tmp
		_ = a
	}
}
