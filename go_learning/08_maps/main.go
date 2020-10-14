package main

import "fmt"

func main() {
	// Define map
	// emails := make(map[string]string)

	// // Assign KV (Key Value)
	// emails["001"] = "test1@gmail.com"
	// emails["002"] = "test2@gmail.com"

	// Declare map and add kV
	emails := map[string]string{"001": "test1@gmail.com", "002": "test2@gmail.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["001"])

	// Delete from map
	delete(emails, "001")
	fmt.Println(emails)
}
