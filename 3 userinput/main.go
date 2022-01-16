package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to sathyam cinemas enjoy your snacks mmmmm !!!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Saptiya da venna")

	input, err := reader.ReadString('\n')
	println(input, err)
	println("mass aa irukku da ithu")

}
