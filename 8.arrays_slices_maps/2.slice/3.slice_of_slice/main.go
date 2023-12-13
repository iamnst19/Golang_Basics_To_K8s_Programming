package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr[0])
	slice_1 := arr[1:8]
	fmt.Println(slice_1)
	slice_2 := slice_1[1:6]
	fmt.Println(slice_2)
}
