package scanner

import (
	"fmt"
	"net"
	"strings"
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

func Scan(target string) {
	evaluated_target, err := check_ip(target)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	
	fmt.Printf("\n[# Scanning Target]---- %s\n", target)
	fmt.Println(evaluated_target)

	workerCount := 100

	// ports := make(chan int, workerCount)

	// var wg sync.WaitGroup
	// for i := 0; i < workerCount; i++ {
	// 	wg.Add(1)

	// 	go func() {
	// 		defer wg.Done()
			
	// 		for port := range ports {
	// 			scan_port(evaluated_target, port)
	// 		}
	// 	} ()
	// }

	// for p := 1; p <= 1024; p++ {
    // 	ports <- p
	// }

	// close(ports)

	// wg.Wait()

	// build task list with port IDs 1..1024
	tasks := make([]Task, 1024)
	for i := range tasks {
		tasks[i] = Task{ID: i + 1}
	}

	wp := WorkerPool{
		Tasks:       tasks,
		concurrency: workerCount,
	}

	wp.Run(evaluated_target)
	fmt.Printf("\n[# Scan Complete]---- %s\n", target)
}

func scan_port(target string, port int) {
	var address string
	if strings.Contains(target, ":") {
		address = fmt.Sprintf("[%s]:%d", target, port)
	} else {
		address = fmt.Sprintf("%s:%d", target, port)
	}

	conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
	if err != nil {
		// connection failed → port likely closed or filtered
		return
	}
	defer conn.Close()

	fmt.Printf("[+] Port %d is open\n", port)
}

func main2() {
	var input string
	fmt.Println("[+] Enter target/s to scan (split multiple targets with ','): ")
	fmt.Scanf("%s", &input)

	if strings.Contains(input, ",") {
		targets := strings.Split(input, ",")
		for _, t := range targets {
			Scan(strings.TrimSpace(t))
		}
	} else {
		Scan(strings.TrimSpace(input))
	}

}

