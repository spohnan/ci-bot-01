package ci_bot

import (
	"net"
)

// Checks to see if a given IP is within a CIDR block
func IsAddrInCIDR(addr string, cidr string) bool {

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
