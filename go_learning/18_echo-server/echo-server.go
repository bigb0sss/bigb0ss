package main

import (
	"io"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 4096) // Buffer for Storing Data

	for {
		data, err := conn.Read(buffer[0:])
		if err == io.EOF {
			log.Println("[INFO] Client Disconnected!")
			break
		} else if err != nil {
			log.Println("[ERROR] Unexpected Error!")
			break
		}
		log.Println("[INFO] Writing Data...")
		log.Printf("[INFO] Received %d bytes: %s\n", data, string(buffer))

		if _, err := conn.Write(buffer[0:data]); err != nil {
			log.Fatalln("[ERROR] Unable to Write Data!")
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":1337")

	if err != nil {
		log.Fatalln("[ERROR] Unable to Bind to Port!")
	}
	log.Println("[INFO] Listening on 0.0.0.0:1337")

	for {
		connect, err := listener.Accept()
		log.Println("[INFO] Received Connection!")

		if err != nil {
			log.Fatalln("[ERROR] Unable to Accept Connection!")
		}
		go echo(connect) // goroutine for concurrency
	}

}
