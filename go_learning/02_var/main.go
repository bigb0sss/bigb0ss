package main

import "fmt"

func main() {
	// Main Types
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	var name string = "bigb0ss"
	var age int32 = 15
	const isCool = true // const = you cannot re-define it

	// Shorthand
	// name2 := "bigb0ss2"
	size := 1.3
	// email := "bigb0ss@gamil.com"

	name2, email := "bigb0ss", "bigb0ss@gmail.com"

	fmt.Println(name, age, isCool, name2, email)
	fmt.Printf("%T\n", size)
}
