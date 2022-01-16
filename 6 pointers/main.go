package main

import "fmt"

func main() {
	fmt.Println("Pointer puluthi da nanga")

	// var ptr *int
	// fmt.Println(ptr)   declaration

	muNum := 100

	var ptr = &muNum //reference
	fmt.Println(ptr, *ptr)

}
