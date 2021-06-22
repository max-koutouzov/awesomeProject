package main

import (
	"encoding/csv"
	"fmt"
	"strings"
)

func main() {
	data := `name,age,has_pet
	Jon,"100",true
	"Fred ""The Hammer"" Smith",42,false
	Martha,37,"true"
	`
	r := csv.NewReader(strings.NewReader(data))
	allData, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	var entries []MyData
	Unmarshal(allData, &entries)
	fmt.Println(entries)

	//now to turn entries into output
	out, err := Marshal(entries)
	if err != nil {
		panic(err)
	}
	sb := &strings.Builder{}
	w := csv.NewWriter(sb)
	w.WriteAll(out)
	fmt.Println(sb)
}
