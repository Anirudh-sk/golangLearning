package main

import "fmt"

func main() {

	fmt.Println("structs are similar to classes and there is no inheritance")

	anirudh := User{"Ani", "ani@ani.com", 69, false}
	fmt.Println(anirudh)
	fmt.Println(anirudh.Name)
	fmt.Println(anirudh.Age)
	anirudh.GetStatus()
	anirudh.NewEmail()
	fmt.Println(anirudh)

}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func (u User) GetStatus() { //this is a method
	fmt.Println("user active ? ", u.Status)
}

func (u User) NewEmail() {
	u.Email = "ani@123.commmmm"
	fmt.Println("email", u.Email)
}
