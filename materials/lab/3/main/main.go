// Build and Use this File to interact with the shodan package
// In this directory lab/3/shodan/main:
// go build main.go
// SHODAN_API_KEY=YOURAPIKEYHERE ./main <search term>

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"shodan/shodan"
)

func main() {

	//Parse the flags
	hostnamePtr := flag.String("dnsresolve", "*", "a string")
	creditsPtr := flag.Bool("showcredits", false, "a bool")
	hostSearchPtr := flag.String("hosttosearch", "*", "a string")
	flag.Parse()

	apiKey := os.Getenv("SHODAN_API_KEY")

	s := shodan.New(apiKey)
	info, err := s.APIInfo()
	if err != nil {
		log.Panicln(err)
	}

	if *hostnamePtr != "*" {
		hostname_ip, err := s.DNSInfo(*hostnamePtr)

		if err != nil {
			log.Panicln("Error with hostname ip!\n")
		}
		fmt.Printf("\nDNS of hostname %s is %s\n", *hostnamePtr, hostname_ip)
		json.MarshalIndent(hostname_ip, "", "\t")
	}

	if *creditsPtr == true {
		fmt.Printf(
			"Query Credits: %d\nScan Credits:  %d\n\n",
			info.QueryCredits,
			info.ScanCredits)
	}

	if *hostSearchPtr != "*" {
		hostSearch, err := s.HostSearch(*hostSearchPtr)
		if err != nil {
			log.Panicln(err)
		}

		fmt.Printf("Host Data Dump\n")
		for _, host := range hostSearch.Matches {
			fmt.Println("==== start ", host.IPString, "====")
			h, _ := json.Marshal(host)
			fmt.Println(string(h))
			fmt.Println("==== end ", host.IPString, "====")
			//fmt.Println("Press the Enter Key to continue.")
			//fmt.Scanln()
		}

		fmt.Printf("IP, Port\n")

		for _, host := range hostSearch.Matches {
			fmt.Printf("%s, %d\n", host.IPString, host.Port)
		}
	}

}
