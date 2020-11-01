package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type testReader struct{} // io.Reader to read from stdin
type testWriter struct{} // io.Writer to write to stdout

// Read Data from stdin
func (reader *testReader) Read(b []byte) (int, error) {
	fmt.Print("[IN] > ")
	return os.Stdin.Read(b)
}

// Write data to stdout
func (writer *testWriter) Write(b []byte) (int, error) {
	fmt.Print("[OUT] > ")
	return os.Stdout.Write(b)
}

func main() {
	var (
		reader testReader
		writer testWriter
	)

	if _, err := io.Copy(&writer, &reader); err != nil {
		log.Fatalln("[ERROR] Unable to Write Data!")
	}
}
