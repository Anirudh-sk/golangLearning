package main

import "fmt"

func main() {
	defer fmt.Println("hello world")
	defer fmt.Println("2")
	defer fmt.Println("1")
	fmt.Println("Hello")
	MyDefer()

}

func MyDefer() {

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

}
