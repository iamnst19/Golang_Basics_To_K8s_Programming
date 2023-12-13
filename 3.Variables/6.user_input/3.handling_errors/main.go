package main
import "fmt"
func main() {
	var name string
	var month int
	fmt.Print("Enter your name and your birth month ")
	count, err := fmt.Scanf("%s %d", &name, &month)
	fmt.Println("count : " ,count)
	fmt.Println("error : " ,err)
	fmt.Println("name :" ,name)
	fmt.Println("month :" ,month)
}