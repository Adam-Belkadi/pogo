package main

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

func check_ip(target string) (string, error) {
	// Try parsing as IP
	ip := net.ParseIP(target)
	if ip != nil {
		return ip.String(), nil 
	}

	// Otherwise resolve hostname
	ips, err := net.LookupIP(target)
	if err != nil || len(ips) == 0 {
		return "", err
	}

	return ips[0].String(), nil
}

func scan(target string) {
	evaluated_target, err := check_ip(target)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(evaluated_target)

	fmt.Printf("\n[# Scanning Target]---- %s\n", target)

	var wg sync.WaitGroup
	for port := 0; port < 1000; port++ {
		wg.Add(1)

		go scan_port(string(evaluated_target), port, &wg)
	}

	wg.Wait()
}

func scan_port(target string, port int, wg *sync.WaitGroup) {
	defer wg.Done()

	address := fmt.Sprintf("%s:%d", target, port)

	conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
	if err != nil {
		// connection failed → port likely closed or filtered
		return
	}
	defer conn.Close()

	fmt.Printf("[+] Port %d is open\n", port)
}

func main() {
	var input string
	fmt.Println("[+] Enter target/s to scan (split multiple targets with ','): ")
	fmt.Scanf("%s", &input)

	if strings.Contains(input, ",") {
		targets := strings.Split(input, ",")
		for _, t := range targets {
			scan(strings.TrimSpace(t))
		}
	} else {
		scan(strings.TrimSpace(input))
	}

}