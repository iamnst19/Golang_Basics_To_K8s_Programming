package main

import "fmt"

func main() {
	var x int = 10
	fmt.Println(x > 5 || x < 15)
	fmt.Println(x > 5 || x < 10)
}
