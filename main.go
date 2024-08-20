package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const (
	green = "\033[32m"
	reset = "\033[0m"
)

func main() {
	printHeader()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		domain := scanner.Text()
		checkDomain(domain)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: could not read from input: %v\n", err)
	}
}

// printHeader prints the table header
func printHeader() {
	fmt.Printf("%-20s %-10s %-10s %-50s %-10s %-50s\n", "Domain", "hasMX", "hasSPF", "SPF Record", "hasDMARC", "DMARC Record")
	fmt.Println(strings.Repeat("-", 140))
}

// checkDomain checks the DNS records for the given domain
func checkDomain(domain string) {
	hasMX, hasSPF, hasDMARC := false, false, false
	var spfRecord, dmarcRecord string

	// Check for MX records
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error looking up MX records for domain %s: %v\n", domain, err)
	} else if len(mxRecords) > 0 {
		hasMX = true
	}

	// Check for SPF records
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error looking up TXT records for domain %s: %v\n", domain, err)
	} else {
		for _, record := range txtRecords {
			if strings.HasPrefix(record, "v=spf1") {
				hasSPF = true
				spfRecord = record
				break
			}
		}
	}

	// Check for DMARC records
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error looking up DMARC records for domain %s: %v\n", domain, err)
	} else {
		for _, record := range dmarcRecords {
			if strings.HasPrefix(record, "v=DMARC1") {
				hasDMARC = true
				dmarcRecord = record
				break
			}
		}
	}

	// Print the results in a formatted table row
	printResult(domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}

// printResult prints the results of the domain check in a formatted row
func printResult(domain string, hasMX, hasSPF bool, spfRecord string, hasDMARC bool, dmarcRecord string) {
	fmt.Printf("%-20s %-10s %-10s %-50s %-10s %-50s\n",
		domain,
		colorize(hasMX),
		colorize(hasSPF),
		spfRecord,
		colorize(hasDMARC),
		dmarcRecord,
	)
}

// colorize returns a string with green color if the value is true
func colorize(value bool) string {
	if value {
		return fmt.Sprintf("%s%t%s", green, value, reset)
	}
	return fmt.Sprintf("%t", value)
}
