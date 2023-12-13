package main

import "fmt"

func main() {
	x := func(l int, b int) int {
		return l * b
	}(5, 6) // here we are passing the arguments along with the function call
	fmt.Printf("%T \n", x)
	fmt.Println(x)
}
