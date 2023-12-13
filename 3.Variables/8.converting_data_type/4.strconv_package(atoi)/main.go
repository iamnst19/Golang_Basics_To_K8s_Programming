package main
import (
	"fmt"
	"strconv"
)
func main() {
	var s string = "24"
	i, err := strconv.Atoi(s) //convert string to integer
	fmt.Printf("%v, %T \n", i, i) 
	fmt.Printf("%v, %T \n", err, err) 
	
}