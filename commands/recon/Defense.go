package recon

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

// GatherDefenseInfo detects security products and defensive mechanisms
func GatherDefenseInfo() string {
	var output strings.Builder

	if runtime.GOOS == "windows" {
		// Windows Defender Status
		output.WriteString("│  Antivirus:\n")
		if out, err := exec.Command("powershell", "-Command", "Get-MpComputerStatus | Select-Object RealTimeProtectionEnabled,IoavProtectionEnabled,AntispywareEnabled | Format-List").Output(); err == nil {
			defenderInfo := string(out)
			if strings.Contains(defenderInfo, "RealTimeProtectionEnabled") {
				lines := strings.Split(defenderInfo, "\n")
				for _, line := range lines {
					line = strings.TrimSpace(line)
					if line != "" {
						output.WriteString(fmt.Sprintf("│    %s\n", line))
					}
				}
			}
		} else {
			// Fallback: Check Windows Defender service
			if out, err := exec.Command("sc", "query", "windefend").Output(); err == nil {
				if strings.Contains(string(out), "RUNNING") {
					output.WriteString("│    Windows Defender: RUNNING\n")
				} else {
					output.WriteString("│    Windows Defender: STOPPED\n")
				}
			}
		}

		// Check for other AV products via WMI
		if out, err := exec.Command("wmic", "/namespace:\\\\root\\SecurityCenter2", "path", "AntiVirusProduct", "get", "displayName,pathToSignedProductExe", "/format:list").Output(); err == nil {
			lines := strings.Split(string(out), "\n")
			avProducts := []string{}
			var currentAV string

			for _, line := range lines {
				line = strings.TrimSpace(line)
				if strings.HasPrefix(line, "displayName=") {
					currentAV = strings.TrimPrefix(line, "displayName=")
				}
				if strings.HasPrefix(line, "pathToSignedProductExe=") && currentAV != "" {
					avProducts = append(avProducts, currentAV)
					currentAV = ""
				}
			}

			if len(avProducts) > 0 {
				output.WriteString("│\n│  Detected AV Products:\n")
				for _, av := range avProducts {
					if av != "" {
						output.WriteString(fmt.Sprintf("│    • %s\n", av))
					}
				}
			}
		}

		// Firewall Status
		output.WriteString("│\n│  Firewall:\n")
		if out, err := exec.Command("netsh", "advfirewall", "show", "allprofiles", "state").Output(); err == nil {
			lines := strings.Split(string(out), "\n")
			for _, line := range lines {
				line = strings.TrimSpace(line)
				if strings.Contains(line, "State") && strings.Contains(line, "ON") {
					output.WriteString("│    Firewall is ENABLED\n")
					break
				} else if strings.Contains(line, "State") && strings.Contains(line, "OFF") {
					output.WriteString("│    Firewall is DISABLED\n")
					break
				}
			}
		}

		// Check for EDR/Security products via processes
		output.WriteString("│\n│  Security Processes:\n")
		securityProcesses := []string{
			"MsMpEng.exe",          // Windows Defender
			"CrowdStrike",          // CrowdStrike Falcon
			"CSFalconService.exe",  // CrowdStrike
			"SentinelAgent.exe",    // SentinelOne
			"CylanceSvc.exe",       // Cylance
			"cb.exe",               // Carbon Black
			"MfeAVSvc.exe",         // McAfee
			"TmListen.exe",         // Trend Micro
			"SavService.exe",       // Sophos
			"TaniumClient.exe",     // Tanium
			"elastic-agent.exe",    // Elastic EDR
			"elastic-endpoint.exe", // Elastic EDR
		}

		if out, err := exec.Command("tasklist").Output(); err == nil {
			taskList := string(out)
			foundProcesses := []string{}

			for _, proc := range securityProcesses {
				if strings.Contains(taskList, proc) {
					foundProcesses = append(foundProcesses, proc)
				}
			}

			if len(foundProcesses) > 0 {
				for _, proc := range foundProcesses {
					output.WriteString(fmt.Sprintf("│    [!] %s\n", proc))
				}
			} else {
				output.WriteString("│    No common EDR processes detected\n")
			}
		}

		// PowerShell Execution Policy
		output.WriteString("│\n│  PowerShell:\n")
		if out, err := exec.Command("powershell", "-Command", "Get-ExecutionPolicy").Output(); err == nil {
			policy := strings.TrimSpace(string(out))
			output.WriteString(fmt.Sprintf("│    Execution Policy: %s\n", policy))
		}

		// Check for PowerShell logging
		if out, err := exec.Command("reg", "query", "HKLM\\SOFTWARE\\Policies\\Microsoft\\Windows\\PowerShell\\ScriptBlockLogging", "/v", "EnableScriptBlockLogging").Output(); err == nil {
			if strings.Contains(string(out), "0x1") {
				output.WriteString("│    Script Block Logging: ENABLED [!]\n")
			} else {
				output.WriteString("│    Script Block Logging: Disabled\n")
			}
		} else {
			output.WriteString("│    Script Block Logging: Not configured\n")
		}

		// Check for module logging
		if out, err := exec.Command("reg", "query", "HKLM\\SOFTWARE\\Policies\\Microsoft\\Windows\\PowerShell\\ModuleLogging", "/v", "EnableModuleLogging").Output(); err == nil {
			if strings.Contains(string(out), "0x1") {
				output.WriteString("│    Module Logging: ENABLED [!]\n")
			} else {
				output.WriteString("│    Module Logging: Disabled\n")
			}
		} else {
			output.WriteString("│    Module Logging: Not configured\n")
		}

		// AppLocker Status
		output.WriteString("│\n│  Application Control:\n")
		if out, err := exec.Command("reg", "query", "HKLM\\SOFTWARE\\Policies\\Microsoft\\Windows\\SrpV2").Output(); err == nil {
			if strings.Contains(string(out), "Exe") || strings.Contains(string(out), "Script") {
				output.WriteString("│    AppLocker: CONFIGURED [!]\n")
			} else {
				output.WriteString("│    AppLocker: Not configured\n")
			}
		} else {
			output.WriteString("│    AppLocker: Not configured\n")
		}

		// UAC Settings
		if out, err := exec.Command("reg", "query", "HKLM\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Policies\\System", "/v", "EnableLUA").Output(); err == nil {
			if strings.Contains(string(out), "0x1") {
				output.WriteString("│    UAC: ENABLED\n")
			} else {
				output.WriteString("│    UAC: DISABLED\n")
			}
		}

		// Check for Sysmon
		if out, err := exec.Command("sc", "query", "Sysmon").Output(); err == nil {
			if strings.Contains(string(out), "RUNNING") {
				output.WriteString("│\n│  [!] SYSMON DETECTED - System is heavily monitored!\n")
			}
		} else if out, err := exec.Command("sc", "query", "Sysmon64").Output(); err == nil {
			if strings.Contains(string(out), "RUNNING") {
				output.WriteString("│\n│  [!] SYSMON64 DETECTED - System is heavily monitored!\n")
			}
		}

		// AMSI Status (for script-based attacks)
		output.WriteString("│\n│  Script Protection:\n")
		if out, err := exec.Command("powershell", "-Command", "[Ref].Assembly.GetType('System.Management.Automation.AmsiUtils')").Output(); err == nil {
			if strings.Contains(string(out), "AmsiUtils") {
				output.WriteString("│    AMSI: Present\n")
			}
		}

	} else {
		// Unix-like systems
		output.WriteString("│  Security Features:\n")

		// Check for SELinux
		if out, err := exec.Command("getenforce").Output(); err == nil {
			selinux := strings.TrimSpace(string(out))
			output.WriteString(fmt.Sprintf("│    SELinux: %s\n", selinux))
		}

		// Check for AppArmor
		if out, err := exec.Command("aa-status").Output(); err == nil {
			if strings.Contains(string(out), "apparmor module is loaded") {
				output.WriteString("│    AppArmor: LOADED\n")
			}
		}

		// Check firewall (iptables/ufw)
		if out, err := exec.Command("ufw", "status").Output(); err == nil {
			if strings.Contains(string(out), "active") {
				output.WriteString("│    UFW Firewall: ACTIVE\n")
			}
		}

		// Check for security monitoring tools
		securityProcs := []string{
			"ossec",
			"wazuh",
			"falco",
			"auditd",
		}

		if out, err := exec.Command("ps", "aux").Output(); err == nil {
			procList := string(out)
			for _, proc := range securityProcs {
				if strings.Contains(procList, proc) {
					output.WriteString(fmt.Sprintf("│    [!] %s detected\n", proc))
				}
			}
		}
	}

	return output.String()
}
