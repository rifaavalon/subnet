package main

import (
	"fmt"
	"net"
)

func main() {
	var uint32

	a := "98.210.237.192/26"
	b := "98.210.237.193/26"

	ipa, ipneta,_ := net.ParseCIDR(a)
	ipb, ipnetb,_ := net.ParseCIDR(b)

	fmt.Println("", a)
	fmt.Println("", b)
	fmt.Println("", ipa)
	fmt.Println("", ipb)
	fmt.Println("", ipneta)
	fmt.Println("", ipnetb)

	fmt.Printf("\nDoes a (%s) contain: b (%s)?\n", ipneta, ipB)

	if ipneta.Contains(ipb){
		fmt.Println("True")
	} else {
	 fmt.Println("False")
	}
}
