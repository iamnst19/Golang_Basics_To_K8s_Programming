package main

import "fmt"

func main() {
	colors := map[string]string{"g": "green", "r": "red", "b": "blue"}
	colors["g"] = "gray"
	fmt.Println(colors)
}
