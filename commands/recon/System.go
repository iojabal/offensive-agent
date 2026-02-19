package recon

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// GatherSystemInfo collects operating system and hardware information
func GatherSystemInfo() string {
	var output strings.Builder

	// Hostname
	if hostname, err := os.Hostname(); err == nil {
		output.WriteString(fmt.Sprintf("│  Hostname     : %s\n", hostname))
	}

	// Operating System
	output.WriteString(fmt.Sprintf("│  OS           : %s\n", runtime.GOOS))
	output.WriteString(fmt.Sprintf("│  Architecture : %s\n", runtime.GOARCH))

	// Windows-specific information
	if runtime.GOOS == "windows" {
		// OS Version using 'ver' command (lightweight)
		if out, err := exec.Command("cmd", "/c", "ver").Output(); err == nil {
			version := strings.TrimSpace(string(out))
			output.WriteString(fmt.Sprintf("│  Version      : %s\n", version))
		}

		// Get detailed OS info from systeminfo (cached to avoid slow execution)
		if out, err := exec.Command("wmic", "os", "get", "Caption,BuildNumber,OSArchitecture", "/value").Output(); err == nil {
			lines := strings.Split(string(out), "\n")
			for _, line := range lines {
				line = strings.TrimSpace(line)
				if strings.HasPrefix(line, "Caption=") {
					output.WriteString(fmt.Sprintf("│  Caption      : %s\n", strings.TrimPrefix(line, "Caption=")))
				} else if strings.HasPrefix(line, "BuildNumber=") {
					output.WriteString(fmt.Sprintf("│  Build        : %s\n", strings.TrimPrefix(line, "BuildNumber=")))
				}
			}
		}

		// System uptime
		if out, err := exec.Command("wmic", "os", "get", "LastBootUpTime", "/value").Output(); err == nil {
			lines := strings.Split(string(out), "\n")
			for _, line := range lines {
				if strings.HasPrefix(line, "LastBootUpTime=") {
					bootTime := strings.TrimPrefix(strings.TrimSpace(line), "LastBootUpTime=")
					if len(bootTime) >= 14 {
						// Parse WMI datetime format: 20231215103045.500000+000
						year := bootTime[0:4]
						month := bootTime[4:6]
						day := bootTime[6:8]
						hour := bootTime[8:10]
						minute := bootTime[10:12]
						output.WriteString(fmt.Sprintf("│  Last Boot    : %s-%s-%s %s:%s\n", year, month, day, hour, minute))
					}
				}
			}
		}

		// Installed hotfixes (patches) - only recent ones
		if out, err := exec.Command("wmic", "qfe", "get", "HotFixID", "/format:list").Output(); err == nil {
			patches := []string{}
			lines := strings.Split(string(out), "\n")
			for _, line := range lines {
				if strings.HasPrefix(line, "HotFixID=") {
					patch := strings.TrimPrefix(strings.TrimSpace(line), "HotFixID=")
					if patch != "" {
						patches = append(patches, patch)
					}
				}
			}
			if len(patches) > 0 {
				// Show only last 5 patches
				count := len(patches)
				if count > 5 {
					output.WriteString(fmt.Sprintf("│  Patches      : %d installed (last 5 shown)\n", count))
					patches = patches[count-5:]
				} else {
					output.WriteString(fmt.Sprintf("│  Patches      : %d installed\n", count))
				}
				for _, patch := range patches {
					output.WriteString(fmt.Sprintf("│                 %s\n", patch))
				}
			}
		}

		// System locale
		if out, err := exec.Command("wmic", "os", "get", "Locale", "/value").Output(); err == nil {
			lines := strings.Split(string(out), "\n")
			for _, line := range lines {
				if strings.HasPrefix(line, "Locale=") {
					locale := strings.TrimPrefix(strings.TrimSpace(line), "Locale=")
					output.WriteString(fmt.Sprintf("│  Locale       : %s\n", locale))
				}
			}
		}
	} else {
		// Unix-like systems
		if out, err := exec.Command("uname", "-r").Output(); err == nil {
			kernel := strings.TrimSpace(string(out))
			output.WriteString(fmt.Sprintf("│  Kernel       : %s\n", kernel))
		}

		// Uptime
		if out, err := exec.Command("uptime", "-p").Output(); err == nil {
			uptime := strings.TrimSpace(string(out))
			output.WriteString(fmt.Sprintf("│  Uptime       : %s\n", uptime))
		}
	}

	// Current process information
	output.WriteString(fmt.Sprintf("│  Process ID   : %d\n", os.Getpid()))

	// Current time
	output.WriteString(fmt.Sprintf("│  Current Time : %s\n", time.Now().Format("2006-01-02 15:04:05")))

	return output.String()
}
