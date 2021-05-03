package main

import "fmt"

func zfunc() {
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z)
}

func main() {
	zfunc()
}
