package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

//TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	iS := make([]string, len(ip))
	for i, v := range ip {
		iS[i] = strconv.Itoa(int(v))
	}
	return strings.Join(iS, ".")
}
func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
