// bhg-scanner/scanner.go modified from Black Hat Go > CH2 > tcp-scanner-final > main.go
// Code : https://github.com/blackhat-go/bhg/blob/c27347f6f9019c8911547d6fc912aa1171e6c362/ch-2/tcp-scanner-final/main.go
// License: {$RepoRoot}/materials/BHG-LICENSE
// Useage:
// {TODO 1: FILL IN}

package scanner

import (
	"fmt"
	"net"
	"sort"
	"time"
)

func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.DialTimeout("tcp", address, 1*time.Second)
		if err != nil {
			//negative numbers for an efficient channel idea from Dr. Borowczak
			results <- -1 * p
			continue
		}
		conn.Close()
		results <- p
	}
}

// for Part 5 - consider
// easy: taking in a variable for the ports to scan (int? slice? ); a target address (string?)?
// med: easy + return  complex data structure(s?) (maps or slices) containing the ports.
// hard: restructuring code - consider modification to class/object
// No matter what you do, modify scanner_test.go to align; note the single test currently fails
func PortScanner() (int, int) {
	var openports []int // notice the capitalization here. access limited!
	var closedports []int

	ports := make(chan int, 100)
	results := make(chan int)

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-results
		if port > 0 {
			openports = append(openports, port)
		} else {
			closedports = append(closedports, -1*port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	sort.Ints(closedports)

	for _, port := range openports {
		fmt.Printf("%d,open\n", port)
	}
	fmt.Printf("\n\n")
	for _, port := range closedports {
		fmt.Printf("%d,closed\n", port)
	}
	fmt.Printf("\n\n")

	return len(openports), len(closedports)
}
