package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "bigb0ss"
	fmt.Println(strings.ToUpper(s))
	fmt.Println(strings.ToLower(s))

	// String search
	fmt.Println(strings.HasPrefix(s, "big"))
	fmt.Println(strings.HasSuffix(s, "b0ss"))
	fmt.Println(strings.Contains(s, "big"))
	fmt.Println(strings.Count(s, "b"))

	fmt.Println(len(s))

	// strings.Join
	// strings.Split
	// strings.ReplaceAll
	
	// Converting Strings and Bytes
	a := "bigb0ss"
	b := []byte(a)
	c := string(b)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}
