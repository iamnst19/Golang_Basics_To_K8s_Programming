package main

import "fmt"

func main() {
	arr := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	slice := arr[0:5]
	fmt.Println(arr)
	fmt.Println(slice)
	slice[1] = 100000
	fmt.Println(slice)
}
