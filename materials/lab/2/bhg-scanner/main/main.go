package main

import "bhg-scanner/scanner"

func main() {
	scanner.PortScanner("scanme.nmap.org", 0, 100, 1, true)
}
