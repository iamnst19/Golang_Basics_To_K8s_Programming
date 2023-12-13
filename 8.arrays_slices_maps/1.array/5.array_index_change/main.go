package main

import "fmt"

func main() {
	fruits := [3]string{"apple", "grapes", "banana"}
	fmt.Println(fruits)
	fruits[2] = "mango"
	fmt.Println(fruits)
	fmt.Println(fruits[2])
}
