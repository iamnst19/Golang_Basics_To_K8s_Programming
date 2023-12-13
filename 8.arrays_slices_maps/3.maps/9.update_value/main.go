package main

import "fmt"

func main() {
	lan := map[string]string{"en": "English", "fr": "French", "es": "Spanish"}
	lan["es"] = "Hindi"
	fmt.Println(lan)
}
