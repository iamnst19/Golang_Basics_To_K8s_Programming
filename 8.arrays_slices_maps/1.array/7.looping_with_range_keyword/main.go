package main

import "fmt"

func main() {
	var grades [3]int = [3]int{97, 85, 93}
	for index, element := range grades {
		fmt.Println(index, "=>", element)
	}
}
