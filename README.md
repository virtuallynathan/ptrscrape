# ptrscrape
So you want to dump some PTR records? 

This code is designed to quickly dump PTR (rDNS) records for a given IPv4 CIDR block. 

Usage:
./ptrscrape -network=192.168.0.0/16

GOMAXPROCS is currently set at 4, with 8 worker goroutines. 
