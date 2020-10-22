// Simple Port Scan Tool in Go
// No concurrent port scan 

package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i <= 100; i++ {
		ip := "scanme.nmap.org"
		scan := fmt.Sprintf("%s:%d", ip, i)
		conn, err := net.Dial("tcp", scan)
		if err != nil {
			//panic("[ERROR] Connection Failed!")
			continue
		}
		conn.Close()
		fmt.Printf("[INFO] Port %d Open\n", i)
	}
}
