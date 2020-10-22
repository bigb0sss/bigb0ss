// Concurrently scanning ports using WaitGroup

package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	ip := "scanme.nmap.org"
	fmt.Printf("[INFO] Starting Port Scan: %s\n", ip)

	var wg sync.WaitGroup // Syncronized Counter

	for i := 0; i <= 100; i++ {
		wg.Add(1) // Increment the counter each time
		go func(j int) {
			defer wg.Done()
			scan := fmt.Sprintf("%s:%d", ip, j)
			conn, err := net.Dial("tcp", scan)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("[INFO] Port %d OPEN\n", j)
			wg.Done() // This needs to end the script. But it can generate some verbose error sometimes.
		}(i)
	}
	wg.Wait()
	fmt.Println("[INFO] Done!")
}
