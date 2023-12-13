package main
import "fmt"
func main() {
	const name = "Jack"
	name = "Peter"
	const is_Smart = true
	const age = 45
	fmt.Printf("%v: %T \n", name, name)
	fmt.Printf("%v: %T \n", is_Smart, is_Smart)
	fmt.Printf("%v: %T \n", age, age)
	

}