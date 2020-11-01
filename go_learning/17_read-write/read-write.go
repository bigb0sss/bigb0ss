package main

import (
	"fmt"
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

	input := make([]byte, 4096) // Set buffer for input/output

	r, err := reader.Read(input)
	if err != nil {
		log.Fatalln("[ERROR] Unable to Read Data!")
	}
	fmt.Printf("[INFO] Read %d bytes from stdin\n", r)

	w, err := writer.Write(input)
	if err != nil {
		log.Fatalln("[ERROR] Unable to Write Data!")
	}
	fmt.Printf("[INFO] Wrote %d bytes from stdout\n", w)

}
