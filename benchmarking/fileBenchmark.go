package main

import (
	"flag"
	"fmt"
	"os"
)

func FileLen(f string, bufsize int) (int, error) {
	file, err := os.Open(f)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	count := 0
	for {
		buf := make([]byte, bufsize)
		num, err := file.Read(buf)
		count += num
		if err != nil {
			break
		}
	}
	return count, nil
}

func main() {
	var filename string
	var buffer int

	flag.StringVar(&filename, "filename", "", "Enter filename and path.")
	flag.IntVar(&buffer, "buffer", 1024, "Enter buffer amount for file reading.")
	flag.Parse()

	result, err := FileLen(filename, buffer)
	if err != nil {
		fmt.Printf("An error occurred %s\n", err)
	}
	fmt.Printf("Total numbers of characters in the file: %v\n", result)

}
