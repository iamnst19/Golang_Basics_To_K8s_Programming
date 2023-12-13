package main

import "fmt"

func main() {
	lan := map[string]string{"en": "English", "fr": "French", "es": "Spanish"}

	// Collect keys to be deleted
	keysToDelete := make([]string, 0)
	for key := range lan {
		keysToDelete = append(keysToDelete, key)
	}
	fmt.Println(keysToDelete)
	// Delete keys outside the loop
	for _, key := range keysToDelete {
		delete(lan, key)
	}

	fmt.Println(lan)
}
