package main

import (
	"fmt"
)

func main() {
	a := "Hello, Hi"
	for i, c := range a {
		fmt.Printf("%d: %s\n", i, string(c))
	}
	fmt.Println("Length of 'Hello, Hi': ", len(a))
}
