package main

import "fmt"

func main() {
	fmt.Println("ithu main function")
	hello()

	result := adder(10, 30)
	fmt.Println("result is ", result)
}

func adder(a int, b int) int {
	return a + b
}

func hello() {
	fmt.Println("hello function da ithu")

}
