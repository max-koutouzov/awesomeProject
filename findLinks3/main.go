package main

import (
	"fmt"
	"gopl.io/ch5/links"
	"log"
	"os"
)

// Enforces a limit of 20 concurrent tokens
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens

	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	unseenLinks := make(chan string)

	// start wih the command line argumnets.
	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()

		// Main go routine de-duplicates worklist items
		// and sends the unseen ones to the crawlers.
		seen := make(map[string]bool)
		for list := range worklist {
			for _, link := range list {
				if !seen[link] {
					seen[link] = true
					unseenLinks <- link
				}
			}
		}
	}
}
