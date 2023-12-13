package main

import "fmt"

func main() {
	src_slice := []int{1, 2, 3, 4, 5}
	dst_slice := make([]int, 3)
	num := copy(dst_slice, src_slice)
	fmt.Println(dst_slice)
	fmt.Println(num)
}
