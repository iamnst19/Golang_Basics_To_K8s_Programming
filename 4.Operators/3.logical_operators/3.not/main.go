package main

import "fmt"

func main() {
	var x, y int = 10, 20
	fmt.Println(!(x > 10))
	fmt.Println(!(true))
	fmt.Println(!(false))
	fmt.Println(!(y < 30))
}
