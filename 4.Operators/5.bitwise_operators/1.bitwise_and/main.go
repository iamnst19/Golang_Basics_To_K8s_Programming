package main

import "fmt"

func main() {
	//bitwise and
	var x, y int = 12, 25 // 1010 &
	// 10100 = 0
	z := x & y     // 10100	= 20
	fmt.Println(z) // 20
}
