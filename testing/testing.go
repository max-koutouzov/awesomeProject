package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func Strings() {
	string1 := flag.String("string1", "", "")
	string2 := flag.String("string2", "", "")
	string3 := flag.String("string3", "", "")
	flag.Parse()

	fmt.Printf("%s, %s, %s\n", *string1, *string2, *string3)
}

func httpGet(url string) {
	getstuff, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(getstuff.Header)
}

func main() {
	Strings()
	httpGet("https://www.bbc.com")
}
