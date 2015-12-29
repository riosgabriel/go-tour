package main

import (
	"fmt"
	"strings"
)

type IPAddr [4]byte

func (ipAdd IPAddr) String() string {
	s := []string{}

	for i := 0; i < len(ipAdd); i++ {
		s = append(s, fmt.Sprintf("%v", ipAdd[i]))
	}
	return strings.Join(s, ".")
}

// TODO: Add a "String() string" method to IPAddr.

func main() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
}
