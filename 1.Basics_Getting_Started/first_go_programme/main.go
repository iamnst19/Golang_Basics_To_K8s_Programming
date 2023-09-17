//every go code must start with a package declaration; write a package keyword and the name of the package
// if you are build reusable package of code you would build package as a shared library
// use package main to make it an executagle programme
package main

// use import to use code from other code
// importing fmt package into the code as we need fmt functionality
import "fmt"

// functions are building block of go programme
// main function is the special type of function ie the entrypoint of a go programme
// no need to call main function explicitly
// curly braces denote starting and ending of programming

func main() {
	fmt.Println("Hello World") //name of the module is fmt and Println is the function '.' operator is the members access operator. it is used to access members of the package
}
