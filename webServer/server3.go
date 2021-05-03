package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu2 sync.Mutex
var count2 int

func main() {
	http.HandleFunc("/", handler3)
	http.HandleFunc("/count", counter3)
	log.Fatal(http.ListenAndServe("localhost:7000", nil))

}

func handler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q]\n", k, v)
	}
}

// counter echoes the number of calls so far.
func counter3(w http.ResponseWriter, r *http.Request) {
	mu2.Lock()
	fmt.Fprintf(w, "Count %d\n", count2)
	mu2.Unlock()
}
