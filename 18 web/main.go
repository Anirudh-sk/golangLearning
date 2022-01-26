package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://aivirex.in"

func main() {
	fmt.Println("WEB req")

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close() // very important

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

}
