// bhg-scanner/scanner.go modified from Black Hat Go > CH2 > tcp-scanner-final > main.go
// Code : https://github.com/blackhat-go/bhg/blob/c27347f6f9019c8911547d6fc912aa1171e6c362/ch-2/tcp-scanner-final/main.go
// License: {$RepoRoot}/materials/BHG-LICENSE
// Useage:
// {TODO 1: FILL IN}
// run: go test
//

package scanner

import (
	"fmt"
	"net"
	"sort"
	"time"
)

//TODO 3 : ADD closed ports; currently code only tracks open ports
var openports []int // notice the capitalization here. access limited!
var closedports []int

func worker(addressToScan string, ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", addressToScan, p)
		conn, err := net.DialTimeout("tcp", address, 5*time.Second) // TODO 2 : REPLACE THIS WITH DialTimeout (before testing!)
		if err != nil {
			results <- 0
			closedports = append(closedports, p)
			continue
		} else {
			openports = append(openports, p)
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
func PortScanner(addressToScan string) (int, int) {

	ports := make(chan int, 100) // TODO 4: TUNE THIS FOR CODEANYWHERE / LOCAL MACHINE
	results := make(chan int)

	for i := 0; i < cap(ports); i++ {
		go worker(addressToScan, ports, results)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			//openports = append(openports, port)
		} else {
			//closedports = append(closedports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	sort.Ints(closedports)

	printPorts()

	lengthOpen := len(openports)
	lengthClosed := len(closedports)
	openports = []int{}
	closedports = []int{}
	return lengthOpen, lengthClosed // TODO 6 : Return total number of ports scanned (number open, number closed);
	//you'll have to modify the function parameter list in the defintion and the values in the scanner_test
}

func printPorts() {
	//TODO 5 : Enhance the output for easier consumption, include closed ports
	fmt.Printf("Open ports: ")
	if len(openports) == 0 {
		fmt.Printf("None\n")
	} else {
		condense(openports)
	}

	fmt.Printf("\nClosed ports: ")
	if len(closedports) == 0 {
		fmt.Printf("None\n")
	} else {
		condense(closedports)
	}
	fmt.Printf("\n")
}

func condense(list []int) {
	fmt.Printf("%v", list[0])

	for i := 1; i < len(list); i++ {
		if list[i] == list[i-1]+1 {
			fmt.Printf("-")
		} else {
			fmt.Printf("%v, %v", list[i-1], list[i])
		}
	}
	fmt.Printf("\n")
}
