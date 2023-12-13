package main

import "fmt"

func main() {
	arr := [5]int{10, 20, 30, 40, 500}
	i := 2
	fmt.Println(arr)
	slices := arr[:i]
	slices_2 := arr[i+1:]
	new_slice := append(slices, slices_2...)
	fmt.Println(new_slice)
}
