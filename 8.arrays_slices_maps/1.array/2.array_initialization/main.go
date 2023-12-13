package main

import "fmt"

func main() {
	var fruits [5]string = [5]string{"apple", "orange", "banana", "grape", "mango"}
	fmt.Println(fruits)

	marks := [3]int{90, 80, 70}
	fmt.Println(marks)

	names := [...]string{"John", "Jane", "Mary"}
	fmt.Println(names)
}
