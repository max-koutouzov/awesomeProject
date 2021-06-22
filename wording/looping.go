package main

import (
	"fmt"
	"sort"
)

func switching() {
	fmt.Print("===============================\n")
	words := []string{"a", "cow", "smile", "gopher", "watermelon", "toilets"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length", wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word")
		}
	}
}

func blankSwitch() {
	fmt.Print("===============================\n")
	words := []string{"hi", "salutations", "hello"}
	for _, word := range words {
		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Println(word, "is a short word!")
		case wordLen > 10:
			fmt.Println(word, "is a long word!")
		default:
			fmt.Println(word, "is exactly the right length")
		}
	}
}

func inner() {
	for i := 0; i < 5; i++ {
		func(j int) {
			fmt.Println("printing", j, "inside from anonymous function")
		}(i)
	}
}

func peopleSlice() {
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}

	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobbert", 23},
		{"Fred", "Fredson", 18},
	}
	fmt.Println(people)
	// Sort by last name
	sort.Slice(people, func(i int, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println(people)
	// sort by age
	sort.Slice(people, func(i int, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)
}

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func main() {
	envVals := []int{2, 4, 6, 8, 10}
	strings := map[string][]string{
		"Marines": []string{"Infantry", "Supply", "Airwing"},
		"Navy":    []string{"Air", "Fire", "Sea"},
	}
	uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}

	mapster := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	for k := range uniqueNames {
		fmt.Println(k)
	}
	for _, v := range envVals {
		fmt.Println(v)
	}
	for k, v := range strings {
		fmt.Println(k, v)
	}

	for i := 0; i < 3; i++ {
		fmt.Println("Loop", i)
		for k, v := range mapster {
			fmt.Println(k, v)
		}
	}
	switching()
	blankSwitch()
	inner()
	peopleSlice()

	twoBase := makeMult(2)
	threeBase := makeMult(3)
	for i := 0; i < 3; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}
}
