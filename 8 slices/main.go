package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("slice na its go not katrina")

	var slice = []string{"cherry", "banana"}

	slice = append(slice, "peach")
	fmt.Println("the fav fruits are", slice)

	slice = append(slice[1:])
	fmt.Println("sliced ones are", slice)

	score := make([]int, 4)

	score[0] = 100
	score[1] = 169
	score[2] = 102
	score[3] = 103
	//score[4]=103  produces error

	score = append(score, 969, 669, 696, 69)

	fmt.Println(sort.IntsAreSorted(score))
	fmt.Println(score)

	sort.Ints(score)

	fmt.Println(sort.IntsAreSorted(score))
	fmt.Println(score)

	// slice manupulation with index

	var tech = []string{"c", "c++", "python", "javascript"}
	var i int = 2

	fmt.Println(tech)

	tech = append(tech[:i], tech[i+1:]...)
	fmt.Println(tech)

}
