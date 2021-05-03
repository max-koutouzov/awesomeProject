package main

import (
	"fmt"
	"net/url"
)

// Values maps a string key to a list of values.
type Values map[string][]string

// Get returns the first value associated
// or "" if there are none.
func (v Values) Get(key string) string {
	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}
	return ""
}

// Add adds the value to key.
// It appends to any existing vlaues associated with key.
func (v Values) Add(key, value string) {
	v[key] = append(v[key], value)
}

func main() {
	m := url.Values{"lang": {"en"}}
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang"))
	fmt.Println(m.Get("q"))
	fmt.Println(m.Get("item"))
	fmt.Println(m["item"])

	m = nil
	fmt.Println(m.Get("item"))

}
