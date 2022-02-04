// bhg-scanner/scanner.go modified from Black Hat Go > CH2 > tcp-scanner-final > main.go
// Code : https://github.com/blackhat-go/bhg/blob/c27347f6f9019c8911547d6fc912aa1171e6c362/ch-2/tcp-scanner-final/main.go
// License: {$RepoRoot}/materials/BHG-LICENSE
// Useage:
// To test, run "go test"
// This program accepts an address, port range, timeout time, and display
// parameter and returns the number of open and closed ports. The output
// can be directed into a .csv file for easy usage.

package scanner

import (
	"fmt"
	"net"
	"sort"
	"time"
)

//This function is launched as a goroutine and scans ports according to the parameters
//in PortScanner. It returns open ports as positive port numbers and closed ports as
//negative port values.
func worker(ports, results chan int, addressToScan string, timeoutSeconds int) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", addressToScan, p)
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

// for Part 5
// easy: taking in a variable for the ports to scan (int? slice? ); a target address (string?)?

//PortScanner takes 5 parameters: an address to scan, ports to start and end at, a timeout
//value (in seconds), and a boolean parameter specifying whether or not to print the closed ports.
//The number of open and closed ports are returned separately.
func PortScanner(addressToScan string, start, end, timeoutSeconds int, showAll bool) (int, int) {
	//This segment checks the validity of the specified port range, and if there
	//is an invalid range, it sets defaults
	if start < 1 || end > 1024 {
		fmt.Printf("Input error: reverting port range to default\n")
		start = 1
		end = 100
	}

	var openports []int // notice the capitalization here. access limited!
	var closedports []int

	//These are the channels the goroutines use to communicate
	ports := make(chan int, 100)
	results := make(chan int)

	//This segment launches the goroutines
	for i := start; i < cap(ports); i++ {
		go worker(ports, results, addressToScan, timeoutSeconds)
	}

	//This segment collects the data from the goroutines
	go func() {
		for i := start; i <= end; i++ {
			ports <- i
		}
	}()

	//This segment sorts the port data into open or closed ports
	for i := start; i <= end; i++ {
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

	//This segment prints the open ports as well as
	// the closed ports, only if the showAll parameter was true
	fmt.Printf("Scanned %s over ports %d to %d with a timeout time of %d second(s). Closed ports were selected to be ", addressToScan, start, end, timeoutSeconds)
	if showAll {
		fmt.Printf("shown.\n")
	} else {
		fmt.Printf("hidden.\n")
	}
	for _, port := range openports {
		fmt.Printf("%d,open\n", port)
	}
	fmt.Printf("\n\n")
	if showAll {
		for _, port := range closedports {
			fmt.Printf("%d,closed\n", port)
		}
		fmt.Printf("\n\n")
	}

	return len(openports), len(closedports)
}
