package main

import "fmt"

func main() {
	arr := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	slice := arr[1:8]
	fmt.Println(len(arr))
	fmt.Println(cap(arr))
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice)) // 9 because the slice starts at index 1 and ends till end of the array hence 9
}
