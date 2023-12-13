package main
import "fmt"
func main() {
	var name string = "Aira"
	var age int     = 8
	var isCute bool = true
	var grades float32 = 91.24

	fmt.Printf("variable name = %v is of type %T \n", name, name)
	fmt.Printf("variable age = %v is of type %T \n", age, age)
	fmt.Printf("variable isCute = %v is of type %T \n", isCute , isCute)
	fmt.Printf("variable grades = %v is of type %T \n", grades , grades)
}
