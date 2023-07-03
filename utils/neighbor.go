package utils

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"strconv"
	"time"
)

func IsFoundHost(host string, port uint16) bool {
	target := fmt.Sprintf("%s:%d", host, port)

	_, err := net.DialTimeout("tcp", target, 1*time.Second)
	if err != nil {
		fmt.Printf("%s %v\n", target, err)
		return false
	}
	return true
}

var PATTERN = regexp.MustCompile(`((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?\.){3})(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)`)

// TODO: Implement broadcast
func FindNeighbors(myHost string, myPort uint16, startIp uint8, endIp uint8, startPort uint16, endPort uint16) []string {
	address := fmt.Sprintf("%s:%d", myHost, myPort)

	m := PATTERN.FindStringSubmatch(myHost)
	if m == nil {
		return nil
	}
	prefixHost := m[1]
	lastIp, _ := strconv.Atoi(m[len(m)-1])
	neighbors := make([]string, 0)

	for port := startPort; port <= endPort; port += 1 {
		for ip := startIp; ip <= endIp; ip += 1 {
			guessHost := fmt.Sprintf("%s%d", prefixHost, lastIp+int(ip))
			guessTarget := fmt.Sprintf("%s:%d", guessHost, port)
			if guessTarget != address && IsFoundHost(guessHost, port) {
				neighbors = append(neighbors, guessTarget)
			}
		}
	}
	return neighbors
}

func GetHost() string {
	// TODO: Hosty data should be retrieved from a config file
	hostname, err := os.Hostname()
	if err != nil {
		// TODO: For production, default should be first nodes IP address
		return "miner_1" // Default to "miner_1" if hostname retrieval fails
	}

	address, err := net.LookupHost(hostname)
	if err != nil {
		return "miner_1" // Default to "miner_1" if host lookup fails
	}

	ip := address[0] // Assuming only the first IP address is needed
	// Extract the last digit from the IP address
	lastDigit := string(ip[len(ip)-1])

	// Convert the last digit to an integer
	ipNum, err := strconv.Atoi(lastDigit)
	if err != nil || ipNum < 1 || ipNum > 3 {
		return "miner_1" // Default to "miner_1" if the last digit is not a valid number
	}

	return fmt.Sprintf("miner_%d", ipNum)
}
