package main

import (
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
Author: Max Koutouzov

Usage example:

go run inout/countingLetters.go -file /usr/share/man/man8/fdisk.8.gz
go run inout/countingLetters.go -file /tmp/test.txtg
*/

func countLetters(r io.Reader) map[string]int {
	/*
		Counting letters in a file using a map
	*/
	buf := make([]byte, 2048)
	out := map[string]int{}
	for {
		n, err := r.Read(buf)
		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}
		if err == io.EOF {
			return out
		}
		if err != nil {
			return nil
		}
	}
}

func buildGzipReader(filename string) (*gzip.Reader, func(), error) {
	/*
		Extracts text from a gzipped file
	*/
	r, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	gr, err := gzip.NewReader(r)
	if err != nil {
		return nil, nil, err
	}
	return gr, func() {
		gr.Close()
		r.Close()
	}, nil
}

func main() {
	var fileName string
	flag.StringVar(&fileName, "file", "", "Path and filename to read.")
	flag.Parse()

	// Checks if filename is .gz or regular text
	res := strings.HasSuffix(fileName, ".gz")
	if !res {
		data, _ := os.Open(fileName)
		checkLetters := countLetters(data)
		fmt.Println(checkLetters)
	} else {
		r, closer, err := buildGzipReader(fileName)
		if err != nil {
			fmt.Println(err)
		}
		defer closer()
		counts := countLetters(r)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(counts)
	}

}
