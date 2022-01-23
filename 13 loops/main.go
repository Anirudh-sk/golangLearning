package main

import "fmt"

func main() {

	fmt.Println("vanthan suttan sethan repeatu")

	days := []string{"sun", "tue", "wed", "fri", "sat"}

	fmt.Println(days)

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])

	}

	//another for loop syntax

	for i := range days {
		fmt.Println(days[i])

	}

	//another syntax similar to foreach

	for index, value := range days {
		fmt.Println("index is ", index, "value is ", value)
	}

	//while loop

	loop := 1
	for loop < 10 {
		if loop == 6 {

			goto poda
		}

		fmt.Println("value ", loop)
		loop++
	}

	//goto
poda:
	fmt.Println("engayo pohuthu")

}
