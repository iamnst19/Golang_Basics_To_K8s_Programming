package main

import (
	"fmt"
	"strings"
)

func main() {
	x := func(s string) {
		fmt.Println(strings.ToLower(s))
	}("RacheL")
	fmt.Printf("%T \n", x)
}
