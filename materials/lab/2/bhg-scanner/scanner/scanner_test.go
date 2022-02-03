package scanner

import (
	"testing"
)

// THESE TESTS ARE LIKELY TO FAIL IF YOU DO NOT CHANGE HOW the worker connects (e.g., you should use DialTimeout)
func TestOpenPort(t *testing.T) {

	open, closed := PortScanner() // Currently function returns number of open and closed ports
	want := 2                     // default value when passing in 1024 TO scanme; also only works because currently PortScanner only returns
	//consider what would happen if you parameterize the portscanner address and ports to scan

	if open != want {
		t.Errorf("got %d, wanted %d, ports closed: %d", open, want, closed)
	}
}

func TestTotalPortsScanned(t *testing.T) {
	// THIS TEST WILL FAIL - YOU MUST MODIFY THE OUTPUT OF PortScanner()

	open, closed := PortScanner() // Currently function returns number of open and closed ports
	want := 1024                  // default value; consider what would happen if you parameterize the portscanner ports to scan

	if (open + closed) != want {
		t.Errorf("got %d, wanted %d", open+closed, want)
	}
}
