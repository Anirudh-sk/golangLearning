package main

import (
	"fmt"
	"net/url"
)

const myURL string = "https://lco.dev:3000/learn?coursename=reactjs&username=anirudh"

func main() {
	fmt.Println("Handling URL")

	//parsing the url
	resutl, _ := url.Parse(myURL)

	fmt.Println(resutl.Scheme)
	fmt.Println(resutl.Host)
	fmt.Println(resutl.Path)
	fmt.Println(resutl.RawQuery)
	fmt.Println(resutl.Port())

	params := resutl.Query()

	fmt.Println(params)
	fmt.Println(params["username"])

	//itration through

	for k, v := range params {
		fmt.Println(v)
		fmt.Println(k)
	}

	partsofURL := &url.URL{
		Scheme:   "https",
		Host:     "anirudh.in",
		Path:     "/King",
		RawQuery: "user=nathanda",
		RawPath:  "hello",
	}
	fmt.Println(partsofURL)

}
