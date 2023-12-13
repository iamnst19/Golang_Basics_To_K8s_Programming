package main
import (
	"fmt"
	"strconv"
)
func main() {
	var i int = 23
	var s string = strconv.Itoa(i) //convert into a string
	fmt.Printf("%q\n", s)
	
}