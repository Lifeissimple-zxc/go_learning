package main

import (
	"fmt"
	"strings"
)

type IPAddr [4]byte

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

// TODO: Add a "String() string" method to IPAddr.
func (i IPAddr) String() string {
	var sb strings.Builder
	for _, val := range i {
		sb.WriteString(fmt.Sprintf("%v.", val))
	}
	return sb.String()[:len(sb.String())-1]
}
