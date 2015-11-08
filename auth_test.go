package ci_bot

import (
	"testing"
)

type item struct {
	addr    string
	cidr    string
	matches bool
}

var testdata = []item{
	item{"192.168.1.67", "192.168.1.0/24", false},
	item{"192.168.1.67", "192.168.1.0/28", false},
	item{"192.168.1.67", "0.0.0.0/0", true},
}

func TestCIDRMatch(t *testing.T) {
	for _, it := range testdata {
		if IsAddrInCIDR(it.addr, it.cidr) != it.matches {
			t.Fatalf("%s in %s should be %t", it.addr, it.cidr, it.matches)
		}
	}
}
