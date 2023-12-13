package main
import (
	"fmt"
    "reflect"
)
func main() {
	fmt.Printf("Type: %v \n", reflect.TypeOf(1000))
	fmt.Printf("Type: %v \n", reflect.TypeOf("Ruben"))
	fmt.Printf("Type: %v \n", reflect.TypeOf(22.45))
	fmt.Printf("Type: %v \n", reflect.TypeOf(true))
}
