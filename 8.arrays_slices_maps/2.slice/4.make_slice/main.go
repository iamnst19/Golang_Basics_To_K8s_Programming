package main

import "fmt"

func main() {
	slice := make([]int, 10)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}
