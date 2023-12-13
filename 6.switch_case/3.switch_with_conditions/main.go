package main

import "fmt"

func main() {
	var i, j int = 10, 20
	switch {
	case i+j == 30:
		fmt.Println("i+j is equal to 30") //implicit break
	case i+j <= 30:
		fmt.Println("i+j is less than or equal to 30")
	default:
		fmt.Println("i+j is greater than 30")
	}
}
