package main

import (
	"fmt"
)

const favColor string = "red"

func main() {
	var guess string

	// Input loop
	for {
		fmt.Println("[*] What is my favorite color? ")

		if _, err := fmt.Scanln(&guess); err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		// Guess the correct color?
		if favColor == guess {
			fmt.Printf("[+] %q is my favorite color!\n", guess)
			return
		}
		// Wrong guess
		fmt.Printf("[-] Wrong guess. %q is not my favorite color. Guess again!\n\n", guess)
	}

}
