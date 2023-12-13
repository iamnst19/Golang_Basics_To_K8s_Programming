package main
import "fmt"
func main() {
	city:= "London"
	{
		country:= "UK"
		fmt.Println(city)
		fmt.Println(country)
	}
//	fmt.Println(country) - this would throw error
	fmt.Println(city)
}