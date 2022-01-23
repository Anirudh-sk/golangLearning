package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("Files manupulation")

	content := "Files contents are this"

	file, err := os.Create("./chummaOruFile.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content)
	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	fmt.Println("length", length)
	defer file.Close()

	readFile("./chummaOruFile.txt")
}

func readFile(filename string) {
	byte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("DAta is : ", string(byte))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
