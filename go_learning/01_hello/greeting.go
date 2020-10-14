package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("[*] Please enter your name: ")
	var name string
	fmt.Scanln(&name) // User input. Wait until user press ENTER
	fmt.Printf("[+] My name is %s.\n", name)
	name = strings.TrimSpace(name) // Remove any new line or spaces
}
