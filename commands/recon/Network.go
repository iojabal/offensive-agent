package recon

import (
	"fmt"
	"net"
	"os/exec"
	"runtime"
	"strings"
)

// GatherNetworkInfo collects network configuration and connectivity information
func GatherNetworkInfo() string {
	var output strings.Builder

	// Get all network interfaces
	interfaces, err := net.Interfaces()
	if err == nil {
		output.WriteString("│  Network Interfaces:\n")
		for _, iface := range interfaces {
			// Skip loopback and down interfaces
			if iface.Flags&net.FlagLoopback != 0 || iface.Flags&net.FlagUp == 0 {
				continue
			}

			output.WriteString(fmt.Sprintf("│\n│    Interface  : %s\n", iface.Name))
			output.WriteString(fmt.Sprintf("│    MAC        : %s\n", iface.HardwareAddr.String()))

			// Get IP addresses for this interface
			addrs, err := iface.Addrs()
			if err == nil {
				for _, addr := range addrs {
					if ipnet, ok := addr.(*net.IPNet); ok {
						if ipnet.IP.To4() != nil {
							output.WriteString(fmt.Sprintf("│    IPv4       : %s\n", ipnet.String()))
						} else {
							output.WriteString(fmt.Sprintf("│    IPv6       : %s\n", ipnet.String()))
						}
					}
				}
			}
		}
	}

	if runtime.GOOS == "windows" {
		// DNS Servers
		if out, err := exec.Command("ipconfig", "/all").Output(); err == nil {
			lines := strings.Split(string(out), "\n")
			inDNSSection := false
			dnsServers := []string{}

			for _, line := range lines {
				line = strings.TrimSpace(line)
				if strings.Contains(line, "DNS Servers") {
					inDNSSection = true
					// Try to extract DNS from same line
					parts := strings.Split(line, ":")
					if len(parts) > 1 {
						dns := strings.TrimSpace(parts[1])
						if dns != "" && net.ParseIP(dns) != nil {
							dnsServers = append(dnsServers, dns)
						}
					}
					continue
				}
				if inDNSSection && line != "" {
					// Check if this is a continuation of DNS servers
					if net.ParseIP(line) != nil {
						dnsServers = append(dnsServers, line)
					} else if !strings.Contains(line, ":") {
						continue
					} else {
						inDNSSection = false
					}
				}
			}

			if len(dnsServers) > 0 {
				output.WriteString("│\n│  DNS Servers:\n")
				for _, dns := range dnsServers {
					output.WriteString(fmt.Sprintf("│    • %s\n", dns))
				}
			}
		}

		// Default Gateway
		if out, err := exec.Command("route", "print", "0.0.0.0").Output(); err == nil {
			lines := strings.Split(string(out), "\n")
			for _, line := range lines {
				line = strings.TrimSpace(line)
				if strings.HasPrefix(line, "0.0.0.0") {
					parts := strings.Fields(line)
					if len(parts) >= 3 {
						gateway := parts[2]
						output.WriteString(fmt.Sprintf("│\n│  Gateway      : %s\n", gateway))
						break
					}
				}
			}
		}

		// Domain Controller (if in domain)
		if out, err := exec.Command("nltest", "/dsgetdc:").Output(); err == nil {
			lines := strings.Split(string(out), "\n")
			for _, line := range lines {
				line = strings.TrimSpace(line)
				if strings.Contains(line, "DC:") {
					parts := strings.Split(line, "\\\\")
					if len(parts) > 1 {
						dc := strings.TrimSpace(parts[1])
						output.WriteString(fmt.Sprintf("│  Domain DC    : %s\n", dc))
					}
				}
				if strings.Contains(line, "Our Site:") {
					parts := strings.Split(line, ":")
					if len(parts) > 1 {
						site := strings.TrimSpace(parts[1])
						output.WriteString(fmt.Sprintf("│  AD Site      : %s\n", site))
					}
				}
			}
		}

		// Active connections (limited to first 10)
		if out, err := exec.Command("netstat", "-ano").Output(); err == nil {
			output.WriteString("│\n│  Active Connections (top 10):\n")
			lines := strings.Split(string(out), "\n")
			connectionCount := 0
			establishedConns := []string{}

			for _, line := range lines {
				line = strings.TrimSpace(line)
				if strings.Contains(line, "ESTABLISHED") {
					parts := strings.Fields(line)
					if len(parts) >= 4 {
						proto := parts[0]
						localAddr := parts[1]
						foreignAddr := parts[2]
						connStr := fmt.Sprintf("%s %s -> %s", proto, localAddr, foreignAddr)
						establishedConns = append(establishedConns, connStr)
						connectionCount++
						if connectionCount >= 10 {
							break
						}
					}
				}
			}

			for i, conn := range establishedConns {
				output.WriteString(fmt.Sprintf("│    [%d] %s\n", i+1, conn))
			}

			if connectionCount == 0 {
				output.WriteString("│    No established connections\n")
			}
		}

		// Listening ports (services)
		if out, err := exec.Command("netstat", "-ano").Output(); err == nil {
			output.WriteString("│\n│  Listening Ports (sample):\n")
			lines := strings.Split(string(out), "\n")
			listenCount := 0

			for _, line := range lines {
				line = strings.TrimSpace(line)
				if strings.Contains(line, "LISTENING") {
					parts := strings.Fields(line)
					if len(parts) >= 4 {
						proto := parts[0]
						localAddr := parts[1]
						pid := parts[len(parts)-1]
						output.WriteString(fmt.Sprintf("│    %s %s (PID: %s)\n", proto, localAddr, pid))
						listenCount++
						if listenCount >= 10 {
							output.WriteString("│    ... (more listening ports not shown)\n")
							break
						}
					}
				}
			}
		}

		// ARP cache (neighboring hosts)
		if out, err := exec.Command("arp", "-a").Output(); err == nil {
			output.WriteString("│\n│  ARP Cache (nearby hosts):\n")
			lines := strings.Split(string(out), "\n")
			arpCount := 0

			for _, line := range lines {
				line = strings.TrimSpace(line)
				if strings.Contains(line, "dynamic") || strings.Contains(line, "static") {
					parts := strings.Fields(line)
					if len(parts) >= 2 {
						ip := parts[0]
						mac := parts[1]
						output.WriteString(fmt.Sprintf("│    %s -> %s\n", ip, mac))
						arpCount++
						if arpCount >= 10 {
							output.WriteString("│    ... (more entries not shown)\n")
							break
						}
					}
				}
			}
		}

	} else {
		// Unix-like systems
		// Default route
		if out, err := exec.Command("ip", "route", "show", "default").Output(); err == nil {
			route := strings.TrimSpace(string(out))
			output.WriteString(fmt.Sprintf("│\n│  Default Route: %s\n", route))
		}

		// DNS resolvers
		if out, err := exec.Command("cat", "/etc/resolv.conf").Output(); err == nil {
			lines := strings.Split(string(out), "\n")
			output.WriteString("│\n│  DNS Servers:\n")
			for _, line := range lines {
				if strings.HasPrefix(strings.TrimSpace(line), "nameserver") {
					parts := strings.Fields(line)
					if len(parts) >= 2 {
						output.WriteString(fmt.Sprintf("│    • %s\n", parts[1]))
					}
				}
			}
		}

		// Active connections
		if out, err := exec.Command("ss", "-tun").Output(); err == nil {
			output.WriteString("│\n│  Active Connections:\n")
			lines := strings.Split(string(out), "\n")
			count := 0
			for _, line := range lines {
				if count > 0 && count <= 10 && strings.TrimSpace(line) != "" {
					output.WriteString(fmt.Sprintf("│    %s\n", strings.TrimSpace(line)))
				}
				count++
				if count > 10 {
					break
				}
			}
		}
	}

	return output.String()
}
