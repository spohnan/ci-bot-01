package ci_bot

import (
	"net"
)

func IsAddrInCIDR(addr string, cidr string) bool {
	_, cidrnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return false
	}

	ip := net.ParseIP(addr)
	if ip == nil {
		return false
	}

	return cidrnet.Contains(ip)
}
