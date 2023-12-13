package main

import (
	"fmt"
)

func main() {
	color := map[string]string{"r": "red", "g": "green", "b": "blue"}
	for key, value := range color {
		fmt.Println(key, "=>", value)
	}
}
