//what are format specifier: They are specifers that help to format differnt kinds of data types
package main
import "fmt"
func main() {
  var name string = "Moti Nagar" //%v formats the value in default format
  fmt.Printf("Nice to see you here, at %v", name)
}