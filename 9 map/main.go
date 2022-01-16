package main

import "fmt"

func main() {
	fmt.Println("maps are not for mapping they are key value pairs")

	lanusges := make(map[string]string)

	lanusges["JS"] = "javascript"
	lanusges["py"] = "python"
	lanusges["html"] = "hypertext markup language"

	fmt.Println("languages are : ", lanusges)
	fmt.Println(" js stanfs for ", lanusges["JS"])

	delete(lanusges, "py")
	fmt.Println("languages are : ", lanusges)

	//looping for maps

	for key, val := range lanusges {
		fmt.Println(key, " stands for ", val)
	}

}
