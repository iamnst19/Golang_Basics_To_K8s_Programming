package main

import "fmt"

func operation(a int, b int) (int, int) {
	sum := a + b
	diff := a - b
	return sum, diff
}

func main() {
	sum, difference := operation(8, 4)
	fmt.Println(sum, difference)
}
