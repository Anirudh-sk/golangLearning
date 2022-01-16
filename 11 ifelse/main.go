package main

import "fmt"

func main() {
	fmt.Println("if else")

	count := 69
	var result string

	if count < 40 {
		result = "yaaru sami ivan"
	} else if count > 30 {
		result = "eatho oru result"
	} else {
		result = "eatho oru result but in else"
	}

	fmt.Println(result)

	//another syntax

	if num := 3; num < 10 {
		fmt.Println("less than 10")
	} else {
		fmt.Println("greateer then 10")
	}

}
