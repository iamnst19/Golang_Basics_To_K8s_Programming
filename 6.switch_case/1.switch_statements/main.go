package main

import "fmt"

func main() {
	var i int = 800
	switch i {
	case 10:
		fmt.Println("i is 10")
	case 20:
		fmt.Println("i is 20")
	case 100, 200:
		fmt.Println("i is 100 or 200")
	default:
		fmt.Println("i is not 10, 20 or 100")
	}
}
