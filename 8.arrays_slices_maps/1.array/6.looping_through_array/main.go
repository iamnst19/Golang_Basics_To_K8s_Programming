package main

import "fmt"

func main() {
	var grades [5]int = [5]int{70, 80, 90, 85, 95}
	for i := 0; i < len(grades); i++ {
		fmt.Println(grades[i])
	}
}
