package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func slicing(inputs []int) []int {
	x := []int{1, 2, 3}
	x = append(x, inputs...)
	return x
}

func wordConcat(start string, finish string) (string, int) {

	starting, err := ioutil.ReadFile("../index.html")
	if err != nil{
		fmt.Printf("Error opening and concatenating words %s\n", err)
	}
	printIt, _ := fmt.Printf("Here you go: %s \n\n", starting)
	printMessage, _ := fmt.Printf("Here's the function string output: %s %s\n", start, finish)

	return string(rune(printIt)), printMessage
}

func capacityChecking()  {
	fmt.Print("\n\nSlice\tLength\tCapacity\n")
	var x []int
	fmt.Println(x, len(x), cap(x))
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 20)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 30)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 50)
	fmt.Println(x, len(x), cap(x))
}

func shareMemory() {
	x := []int{1, 2, 3, 4}
	y := x[:2]
	z := x[1:]
	x[1] = 20
	y[0] = 10
	z[1] = 30
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}


func makingThngs() []int {
	x := make([]int, 0, 10)
	x = append(x, 100, 5, 6, 3, 2)
	return x
}

func appendSlicing() {
	x := []int{1, 2, 3, 4}
	y := x[:2]
	fmt.Println(cap(x), cap(y))
	y = append(y, 30)
	y = append(y, 87, 88, 44)
	print("Append Slicing\n")
	fmt.Println("x:", x)
	fmt.Println("y:", y)
}

func sliciy(x []string) []string {
	var (
		z = append(x, "boogers")
	)
	return z
}

func confusingSlices() {
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	y := x[:2:2]
	z := x[2:4:4]
	fmt.Println(cap(x), cap(y), cap(z))
	y = append(y, 30, 40, 50)
	x = append(x, 60)
	z = append(z, 70)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}


func stringy(input string) (string, string, string) {
	var s2 = input[4:7]
	var s3 = input[:5]
	var s4 = input[6:]
	//fmt.Println(s2, s3, s4)
	return s2, s3, s4
}

func stringConversion(input string) (bs []byte, rs []rune) {
	bs = []byte(input)
	rs = []rune(input)
	return bs, rs
}

func mappy() map[string]int {
	totalWins := map[string]int{}
	return totalWins
}

func mappyTeams() (map[string][]string, map[int][]string) {
	teams := map[string][]string {
		"Orcas": []string{"Fred", "Ralph", "Bijou"},
		"Lions": []string{"Sarah", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"},
	}
	ages := make(map[int][]string,10)


	return teams, ages
}


func results() {
	totalWins := mappy()
	mappyTeams()
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2
	fmt.Printf("Orcas won: %v\n", totalWins["Orcas"])
	fmt.Printf("Kittens won: %v\n", totalWins["Kittens"])
	totalWins["Kittens"]++
	fmt.Printf("New Kittens score: %v\n",totalWins["Kittens"])
	totalWins["Lions"] = 3
	fmt.Printf("Lions win: %v\n", totalWins["Lions"])
}


func inSetting() (map[int]bool, []int) {
	inSet := map[int]bool{}
	vals := []int{5, 7, 3, 4,6, 10, 11, 5, 4, 17, 18}
	for _, v := range vals {
		inSet[v] = true
	}
	fmt.Println(len(vals), len(inSet))
	fmt.Println(inSet[5])
	fmt.Println(inSet[500])
	if inSet[6] {
		fmt.Println("6 in set")
	}
	return inSet, vals
}

func structing() {
	var person struct{
		name string
		age int
		pet string
	}

	person.name = "Bob"
	person.age = 50
	person.pet = "Dog"

	pet := struct{
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}
	fmt.Println(json.Marshal(person))
	fmt.Println(json.Marshal(pet))
}



func main() {
	wordConcat("Printed", "Output")
	fmt.Print(slicing([]int{3, 5, 66, 39}))
	capacityChecking()
	fmt.Print(makingThngs())
	x := sliciy([]string{"hello", "cheese"})
	fmt.Print(x)
	shareMemory()
	appendSlicing()
	confusingSlices()
	fmt.Println(stringy("Hello there 🌞"))
	fmt.Println(stringConversion("Hello, 🌞"))
	results()
	inSetting()
	structing()
}