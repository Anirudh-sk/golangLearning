package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Dice Golang")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1
	fmt.Println(diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice Value is", diceNumber, "you can enter or move one position")
	case 2:
		fmt.Println("Dice Value is", diceNumber, "you can move two position")
	case 3:
		fmt.Println("Dice Value is", diceNumber, "you can move three position")
	case 4:
		fmt.Println("Dice Value is", diceNumber, "you can move four position")
	case 5:
		fmt.Println("Dice Value is", diceNumber, "you can move five position")
	case 6:
		fmt.Println("Dice Value is", diceNumber, "you can move six position and play again")
	default:
		fmt.Println("ena nae solringa")

	}

}
