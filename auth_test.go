package bot

import (
	"testing"
)

type item struct {
	addr    string
	cidr    string
	matches bool
}

var testdata = []item{
	{"192.168.1.67", "192.168.1.0/24", true},
	{"192.168.1.67", "192.168.1.0/28", false},
	{"192.168.1.67", "0.0.0.0/0", true},
	{"not_an_ip", "0.0.0.0/0", false},
	{"192.168.1.67", "not_a_cidr", false},
	{"", "0.0.0.0/0", false},
	{"192.168.1.67", "", false},
	{"", "", false},
}

func TestCIDRMatch(t *testing.T) {
	for _, it := range testdata {
		if isAddrInCIDR(it.addr, it.cidr) != it.matches {
			t.Fatalf("%s in %s should be %t", it.addr, it.cidr, it.matches)
		}
	}
}
