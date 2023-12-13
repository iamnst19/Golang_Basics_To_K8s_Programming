package main

import "fmt"

func main() {
	arr := [4]int{10, 20, 30, 40}
	slice := arr[1:3]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	slice = append(slice, 50)
	fmt.Println("appemded 50")
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	slice = append(slice, 50, 10000)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

}
