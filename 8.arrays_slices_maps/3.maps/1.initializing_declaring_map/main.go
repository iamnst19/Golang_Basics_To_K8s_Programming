package main

import "fmt"

func main() {
	var colors map[string]string
	colors["g"] = "green"
	fmt.Println(colors)
}

// Output:
// panic: assignment to entry in nil map
//this is because we have not initialized the map
