package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Anirudh server")
	GetReq()
}

func GetReq() {
	const url = "http://localhost:8000/get"

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Println(res)
	content, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))
}
