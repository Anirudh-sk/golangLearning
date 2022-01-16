package main

import "fmt"

func main() {
	fmt.Println("array tha da gethu")

	var arr [5]string

	arr[0] = "joey"
	arr[1] = "chandler"
	arr[2] = "rachel"
	// arr[3] = "phoebe"
	arr[4] = "monica"

	fmt.Println(arr) //the length will alwasy be 5 irrevalent of the number of elements
	var arr1 = [5]string{"cherry", "peach"}
	fmt.Println(arr1)

}
