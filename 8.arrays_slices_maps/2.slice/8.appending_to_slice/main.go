package main

import (
	"fmt"
)

func main() {
	arr := [5]int{10, 20, 30, 40, 50}
	slices := arr[:2]
	fmt.Println(arr)
	fmt.Println(slices)
	arr_2 := [5]int{15, 25, 35, 45, 55}
	slices_2 := arr_2[:2]
	fmt.Println(arr_2)
	fmt.Println(slices_2)
	new_slice := append(slices, slices_2...)
	fmt.Println(new_slice)

}
