package main

import "fmt"

func main() {

	fmt.Println("structs are similar to classes and there is no inheritance")

	anirudh := User{"Ani", "ani@ani.com", 69, false}
	fmt.Println(anirudh)
	fmt.Println(anirudh.Name)
	fmt.Println(anirudh.Age)

}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}
