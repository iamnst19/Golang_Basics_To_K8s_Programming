package main

import "fmt"

func main() {
	arr := [5]int{10, 20, 90, 70, 60}
	slice := arr[:3]
	fmt.Println(cap(slice))

	slice_2 := make([]int, 5, 20) // 5 is length and 20 is capacity
	new_slice := append(slice, slice_2...)
	fmt.Println(cap(new_slice))
	fmt.Println(new_slice) // [10 20 90 0 0 0 0 0 0 0] because slice_2 is of length 5 and capacity 20
}
