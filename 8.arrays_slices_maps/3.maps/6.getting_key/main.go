package main

import "fmt"

func main() {
	code := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(code)
	value, found := code["b"]
	fmt.Println(value, found)
	value, found = code["c"]
	fmt.Println(value, found)
	value, found = code["d"]
	fmt.Println(value, found)
}
