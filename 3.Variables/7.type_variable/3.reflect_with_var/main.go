package main
import (
	"fmt"
	"reflect"
)
func main() {
	var name string = "Aira"
	var age int     = 8
	var isCute bool = true
	var grades float32 = 91.24

	fmt.Printf("variable name = %v is of type %v \n", name, reflect.TypeOf(name))
	fmt.Printf("variable age = %v is of type %v \n", age, reflect.TypeOf(age))
	fmt.Printf("variable isCute = %v is of type %v \n", isCute , reflect.TypeOf(isCute))
	fmt.Printf("variable grades = %v is of type %v \n", grades , reflect.TypeOf(grades))
}
