package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"runtime"
)

var addrs chan string

func main() {
	network := flag.String("network", "8.8.8.8/32", "the network to do PTR lookups for")
	flag.Parse()

	addrs = make(chan string)

	for w := 1; w <= 8; w++ {
		go printPtr(addrs)
	}

	ip, ipnet, err := net.ParseCIDR(*network)
	if err != nil {
		log.Printf("Error: %s", err.Error())
	}
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		addrs <- ip.String()
	}

}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func printPtr(addrs chan string) {
	for addr := range addrs {
		result, _ := net.LookupAddr(addr)
		if len(result) > 0 {
			fmt.Printf("%s, %s \n", addr, result[0])
		}
	}
}
