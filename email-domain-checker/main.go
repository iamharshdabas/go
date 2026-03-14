package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
)

func main() {
	var domain string
	flag.StringVar(&domain, "domain", "gmail.com", "domain to search")
	flag.Parse()

	if domain == "gmail.com" {
		fmt.Println("using default domain: gmail.com")
		fmt.Println("use -domain flag to specify a different domain")
	}

	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	if mxRecords, err := net.LookupMX(domain); err != nil {
		fmt.Println("error looking up MX records:", err)
	} else {
		hasMX = len(mxRecords) > 0
	}

	if txtRecords, err := net.LookupTXT(domain); err != nil {
		fmt.Println("error looking up TXT records:", err)
	} else {
		for _, txt := range txtRecords {
			if strings.HasPrefix(txt, "v=spf1") {
				hasSPF = true
				spfRecord = txt
				break
			}
		}
	}

	if dmarcRecords, err := net.LookupTXT("_dmarc." + domain); err != nil {
		fmt.Println("error looking up DMARC records:", err)
	} else {
		for _, txt := range dmarcRecords {
			if strings.HasPrefix(txt, "v=DMARC1") {
				hasDMARC = true
				dmarcRecord = txt
				break
			}
		}
	}

	fmt.Printf("\nDomain: %s\n", domain)
	fmt.Printf("Has MX records: %t\n", hasMX)
	fmt.Printf("\nHas SPF record: %t\n", hasSPF)
	if hasSPF {
		fmt.Printf("SPF record: %s\n", spfRecord)
	}
	fmt.Printf("\nHas DMARC record: %t\n", hasDMARC)
	if hasDMARC {
		fmt.Printf("DMARC record: %s\n", dmarcRecord)
	}
}
