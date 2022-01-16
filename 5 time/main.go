package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time is relative")

	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) // format

	createdDate := time.Date(2003, time.June, 4, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
}
