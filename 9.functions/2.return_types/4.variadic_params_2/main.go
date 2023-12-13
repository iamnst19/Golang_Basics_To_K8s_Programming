package main

import "fmt"

func employesDetails(name string, skills ...string) {
	fmt.Println("hi", name, "your skills are - ")
	for _, s := range skills {
		fmt.Printf("%s", s)
	}

}

func main() {
	employesDetails("Jack ", "C ", " C++", " Python", " Java")
}
