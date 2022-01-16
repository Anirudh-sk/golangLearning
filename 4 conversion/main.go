package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	println("Welcome to Vanga saapadalam")
	println("vanthathukku rating potutu po da venna")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	println("kalambu kalambu")
	println(input)

	Rating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
		fmt.Println(Rating)
		println(err)
		println(Rating)
	} else {
		println(Rating + 1)
		fmt.Println(Rating + 1)
	}
}
