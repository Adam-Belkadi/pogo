package scanner

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

var (
	mostUsedPorts = []int{20, 21, 22, 23, 25, 37, 42, 43, 49, 53, 67, 68, 69, 70, 79, 80, 88, 102, 110, 111, 113, 119, 123, 135, 137, 138, 139, 143, 161, 162, 179, 194, 201, 264, 318, 381, 383, 389, 411, 427, 443, 445, 464, 465, 487, 500, 512, 513, 514, 515, 520, 521, 540, 548, 554, 563, 587, 591, 593, 631, 636, 639, 646, 691, 860, 873, 902, 989, 990, 993, 995, 1025, 1026, 1027, 1080, 1099, 1194, 1214, 1241, 1311, 1337, 1352, 1433, 1434, 1494, 1512, 1521, 1589, 1701, 1720, 1723, 1755, 1812, 1813, 1863, 1900, 1935, 1947, 1985, 2000, 2001, 2049, 2082, 2083, 2086, 2087, 2095, 2096, 2100, 2222, 2302, 2483, 2484, 2967, 3000, 3050, 3074, 3128, 3222, 3260, 3268, 3269, 3306, 3389, 3690, 4333, 4444, 4500, 4567, 4662, 4672, 4899, 5000, 5001, 5004, 5005, 5037, 5060, 5061, 5101, 5190, 5222, 5223, 5432, 5433, 5500, 5631, 5666, 5800, 5900, 6000, 6001, 6112, 6379, 6443, 6500, 6566, 6660, 6661, 6662, 6663, 6664, 6665, 6666, 6667, 6668, 6669, 6697, 6881, 6882, 6883, 6884, 6885, 6886, 6887, 6888, 6889, 6890, 6901, 6969, 7000, 7001, 7070, 8000, 8008, 8009, 8010, 8080, 8081, 8086, 8087, 8088, 8090, 8091, 8096, 8118, 8123, 8181, 8222, 8332, 8333, 8400, 8443, 8500, 8600, 8649, 8880, 8888, 9000, 9001, 9042, 9060, 9080, 9090, 9091, 9200, 9300, 9418, 9443, 9999, 10000, 11211, 12000, 16080, 18080, 19132, 19294, 20000, 25565, 27015, 27017, 27018, 27019, 28015, 31337, 50000, 50070, 50075, 50090}
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

func check_port(port_string string) []int {
	var ports []int
	if strings.Contains(port_string, ",") {
		parts := strings.Split(port_string, ",")
		ports = make([]int, 0, len(parts))
		for _, p := range parts {
			port, err := strconv.Atoi(strings.TrimSpace(p))
			if err == nil {
				ports = append(ports, port)
			}
		}
	} else if port_string == "-" {
		ports = make([]int, 65535)
		for i := range ports {
			ports[i] = i + 1
		}
		fmt.Println("[#] Scanning all ports (1-65535)...")
	} else if strings.Contains(port_string, "-") {
		ports = make([]int, 0)
		parts := strings.Split(port_string, "-")
		if len(parts) == 2 {
			start, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
			end, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
			if err1 == nil && err2 == nil && start <= end {
				ports = make([]int, 0, end-start+1)
				for i := start; i <= end; i++ {
					ports = append(ports, i)
				}
			}
		}
	} else {
		if port, err := strconv.Atoi(strings.TrimSpace(port_string)); err == nil {
			ports = []int{port}
		}
	}
	return ports
}

func Scan(target string, port_string string) {
	evaluated_target, err := check_ip(target)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	ports := mostUsedPorts

	if port_string != "most used" {
		ports = check_port(port_string)
	}

	fmt.Printf("[#] Scanning Target: %s --- %s\n\n", evaluated_target, target)

	workerCount := 100

	tasks := make([]Task, len(ports))
	for i, port := range ports {
		tasks[i] = Task{ID: port}
	}

	wp := WorkerPool{
		Tasks:       tasks,
		concurrency: workerCount,
	}

	wp.Run(evaluated_target)
	fmt.Printf("\n[#] Scan Complete: %s --- %s\n", evaluated_target, target)
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

	conn.SetReadDeadline(time.Now().Add(1 * time.Second))

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Printf("[+] Port %d is open: Unable to detect version\n", port)
		return
	}

	banner := strings.TrimRight(string(buffer[:n]), "\r\n")

	if len(banner) > 0 {
		fmt.Printf("[+] Port %d is open: %s\n", port, banner)
	} else {
		fmt.Printf("[+] Port %d is open: No banner retrieved\n", port)
	}

}
