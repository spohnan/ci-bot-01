// Package bot is the top level container for ci-bot-01 code
package bot

import (
	"encoding/csv"
	"log"
	"net"
	"os"
	"strings"
)

// IsRequestAuthorized checks to see if the requesting IP address
// has been allowed explicitly or as part of a CIDR range
func IsRequestAuthorized(addr string) bool {

	wl := os.Getenv("CI_BOT_IP_WHITELIST")
	if wl == "" {
		return false
	}

	wlRecords := csv.NewReader(strings.NewReader(wl))
	addrOK, err := wlRecords.ReadAll()
	if err != nil {
		log.Println(err)
	}

	for records := range addrOK {
		r := addrOK[records]
		for i := 0; i < len(r); i++ {
			// Test to see if the IP addr is part of an allowed CIDR
			// range or if it's a direct match to a white listed address
			if isRange(r[i]) && isAddrInCIDR(addr, r[i]) || addr == r[i] {
				return true
			}
		}
	}

	return false

}

// isRange looks for a / character to see if it's a CIDR notation
// range of IPs
func isRange(addr string) bool {
	return strings.ContainsRune(addr, '/')
}

// isAddrInCIDR checks to see if a given IP is within a CIDR block
func isAddrInCIDR(addr string, cidr string) bool {

	ip := net.ParseIP(addr)
	if ip == nil {
		return false
	}

	_, cidrnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return false
	}

	return cidrnet.Contains(ip)

}
