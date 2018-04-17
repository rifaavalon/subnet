package main

import (
	"fmt"
	"net"
)

var err error

func IntToIP(ip uint32) net.IP {
	result := make(net.IP, 4)
	result[0] = byte(ip)
	result[1] = byte(ip >> 8)
	result[2] = byte(ip >> 16)
	result[3] = byte(ip >> 24)
	return result
}

func main() {
	z := IntToIP(0x62D2ED4B).String()

	a := "98.210.237.192/26"
	b := z

	ipa, ipneta, _ := net.ParseCIDR(a)
	ipb, ipnetb, _ := net.ParseCIDR(b)

	fmt.Println("Network address A: ", a)
	fmt.Println("IP address      B: ", b)
	fmt.Println("ipA              : ", ipa)
	fmt.Println("ipnetA           : ", ipneta)
	fmt.Println("ipB              : ", ipb)
	fmt.Println("ipnetB           : ", ipnetb)

	fmt.Printf("\nDoes a (%s) contain: b (%s)?\n", ipneta, ipb)

	if ipneta.Contains(ipb) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
