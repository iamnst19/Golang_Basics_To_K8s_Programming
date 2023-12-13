package main

import "fmt"

func returnCube(n int) int {
	return n * n * n
	// fmt.Println(n)
}
func main() {
	result := returnCube(3)
	fmt.Println(result)
}
