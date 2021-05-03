package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// WaitForServer attempts to contact the server of a URL.
// It tries for one minute using exponential back-off
// It reports and error if all attempts fail.

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}
		log.Printf("Server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) // exponential back off
	}
	return fmt.Errorf("Server %s failed to respond after %s", url, timeout)
}

func main() {
	url := flag.String("url", "", "Use: --url <https://www.google.com/")
	flag.Args()
	if err := WaitForServer(*url); err != nil {
		fmt.Printf("Site is down or unreachable: %v\n", err)
		os.Exit(1)
	}
}
