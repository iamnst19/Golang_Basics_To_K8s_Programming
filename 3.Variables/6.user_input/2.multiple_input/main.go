package main
import "fmt"
func main() {
	var name string
	var is_fun bool
	fmt.Print("Enter your name and are you fun at work: ")
	fmt.Scanf("%s %t", &name, &is_fun)
	fmt.Println(name, is_fun)
}