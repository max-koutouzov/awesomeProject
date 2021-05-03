package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func main() {
	getit, err := httpGetBody("https://www.bbc.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", getit)
}
