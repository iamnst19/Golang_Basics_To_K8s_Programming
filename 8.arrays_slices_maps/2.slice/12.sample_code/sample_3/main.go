package main

import "fmt"

func main() {
	arr := []int{10, 20, 90, 70, 60}
	slice := make([]int, 10)
	num := copy(slice, arr) //source is arr and destination is slice
	fmt.Println(slice)
	fmt.Println(num) //prints the number of elements copied
}
