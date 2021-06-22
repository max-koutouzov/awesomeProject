package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type Order struct {
	ID          string    `json:"id"`
	DateOrdered time.Time `json:"date_ordered"`
	CustomerID  string    `json:"customer_id"`
	Items       []Item    `json:"items"`
}

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	data, err := ioutil.ReadFile("encoding/data.json")
	if err != nil {
		fmt.Println("File reading error: ", err)
	}
	var o Order
	err = json.Unmarshal(data, &o)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(o)
}
