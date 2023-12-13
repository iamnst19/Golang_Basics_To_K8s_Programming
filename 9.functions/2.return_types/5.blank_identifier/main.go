package main

import "fmt"

func f() (int, int) {
	return 42, 27
}

// return only one value
func main() {
	v, _ := f()
	fmt.Println(v)
}
