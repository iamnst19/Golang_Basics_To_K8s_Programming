package main

import "fmt"

func main() {
	var x, y int = 210, 20
	y %= x
	fmt.Println(y)
}
