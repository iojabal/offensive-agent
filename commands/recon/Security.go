package recon

import (
	"fmt"
	"os/exec"
	"os/user"
	"runtime"
	"strings"
)

// GatherSecurityContext collects user, privilege, and domain information
func GatherSecurityContext() string {
	var output strings.Builder

	// Current user information
	currentUser, err := user.Current()
	if err == nil {
		output.WriteString(fmt.Sprintf("│  Username     : %s\n", currentUser.Username))
		output.WriteString(fmt.Sprintf("│  User ID      : %s\n", currentUser.Uid))
		if currentUser.Gid != "" {
			output.WriteString(fmt.Sprintf("│  Group ID     : %s\n", currentUser.Gid))
		}
		if currentUser.HomeDir != "" {
			output.WriteString(fmt.Sprintf("│  Home Dir     : %s\n", currentUser.HomeDir))
		}
	}

	if runtime.GOOS == "windows" {
		// Whoami - Full information
		if out, err := exec.Command("whoami").Output(); err == nil {
			username := strings.TrimSpace(string(out))
			output.WriteString(fmt.Sprintf("│  Full Name    : %s\n", username))
		}

		// Domain membership
		if out, err := exec.Command("wmic", "computersystem", "get", "domain", "/value").Output(); err == nil {
			lines := strings.Split(string(out), "\n")
			for _, line := range lines {
				if strings.HasPrefix(line, "Domain=") {
					domain := strings.TrimPrefix(strings.TrimSpace(line), "Domain=")
					if domain != "" && !strings.EqualFold(domain, "WORKGROUP") {
						output.WriteString(fmt.Sprintf("│  Domain       : %s\n", domain))
					} else {
						output.WriteString("│  Domain       : Not joined (Workgroup)\n")
					}
				}
			}
		}

		// User groups
		if out, err := exec.Command("whoami", "/groups").Output(); err == nil {
			output.WriteString("│\n│  Groups:\n")
			lines := strings.Split(string(out), "\n")
			inGroupSection := false
			groupCount := 0
			for _, line := range lines {
				line = strings.TrimSpace(line)
				if strings.Contains(line, "GROUP INFORMATION") {
					inGroupSection = true
					continue
				}
				if inGroupSection && line != "" && !strings.Contains(line, "---") && !strings.Contains(line, "Group Name") {
					// Extract group name (first column)
					parts := strings.Fields(line)
					if len(parts) > 0 && !strings.HasPrefix(parts[0], "=") {
						groupCount++
						if groupCount <= 10 { // Limit to 10 groups for readability
							output.WriteString(fmt.Sprintf("│    [%d] %s\n", groupCount, parts[0]))
						}
					}
				}
			}
			if groupCount > 10 {
				output.WriteString(fmt.Sprintf("│    ... and %d more groups\n", groupCount-10))
			}
		}

		// User privileges
		if out, err := exec.Command("whoami", "/priv").Output(); err == nil {
			output.WriteString("│\n│  Privileges:\n")
			lines := strings.Split(string(out), "\n")
			inPrivSection := false
			privCount := 0
			enabledPrivs := []string{}
			disabledPrivs := []string{}

			for _, line := range lines {
				line = strings.TrimSpace(line)
				if strings.Contains(line, "PRIVILEGES INFORMATION") {
					inPrivSection = true
					continue
				}
				if inPrivSection && line != "" && !strings.Contains(line, "---") && !strings.Contains(line, "Privilege Name") {
					parts := strings.Fields(line)
					if len(parts) >= 2 {
						privName := parts[0]
						privStatus := parts[len(parts)-1]
						privCount++

						// Check if privilege is enabled
						if strings.Contains(strings.ToLower(privStatus), "enabled") {
							enabledPrivs = append(enabledPrivs, privName)
						} else {
							disabledPrivs = append(disabledPrivs, privName)
						}
					}
				}
			}

			// Show enabled privileges first (most important)
			if len(enabledPrivs) > 0 {
				output.WriteString("│    Enabled:\n")
				for _, priv := range enabledPrivs {
					// Highlight dangerous privileges
					marker := ""
					if isDangerousPrivilege(priv) {
						marker = " [!]"
					}
					output.WriteString(fmt.Sprintf("│      • %s%s\n", priv, marker))
				}
			}

			// Show disabled privileges (collapsed)
			if len(disabledPrivs) > 0 {
				output.WriteString(fmt.Sprintf("│    Disabled: %d privileges\n", len(disabledPrivs)))
			}
		}

		// Integrity level
		if out, err := exec.Command("whoami", "/groups").Output(); err == nil {
			lines := strings.Split(string(out), "\n")
			for _, line := range lines {
				if strings.Contains(line, "Mandatory Label") {
					parts := strings.Fields(line)
					for i, part := range parts {
						if strings.Contains(part, "Mandatory") && i+2 < len(parts) {
							integrityLevel := parts[i+2]
							output.WriteString(fmt.Sprintf("│\n│  Integrity    : %s\n", integrityLevel))
							break
						}
					}
					break
				}
			}
		}

		// Check if elevated (admin)
		isAdmin := IsElevated()
		output.WriteString(fmt.Sprintf("│  Elevated     : %v\n", isAdmin))

	} else {
		// Unix-like systems
		if out, err := exec.Command("id").Output(); err == nil {
			idInfo := strings.TrimSpace(string(out))
			output.WriteString(fmt.Sprintf("│  ID Info      : %s\n", idInfo))
		}

		// Check if root
		if out, err := exec.Command("id", "-u").Output(); err == nil {
			uid := strings.TrimSpace(string(out))
			isRoot := uid == "0"
			output.WriteString(fmt.Sprintf("│  Root         : %v\n", isRoot))
		}

		// Sudo rights
		if out, err := exec.Command("sudo", "-l", "-n").Output(); err == nil {
			output.WriteString("│  Sudo Rights  : Available\n")
			sudoInfo := strings.TrimSpace(string(out))
			if len(sudoInfo) < 200 {
				output.WriteString(fmt.Sprintf("│                 %s\n", sudoInfo))
			}
		}
	}

	return output.String()
}

// isDangerousPrivilege checks if a privilege is commonly used for privilege escalation
func isDangerousPrivilege(priv string) bool {
	dangerous := []string{
		"SeDebugPrivilege",
		"SeImpersonatePrivilege",
		"SeAssignPrimaryTokenPrivilege",
		"SeTcbPrivilege",
		"SeBackupPrivilege",
		"SeRestorePrivilege",
		"SeLoadDriverPrivilege",
		"SeTakeOwnershipPrivilege",
	}

	for _, d := range dangerous {
		if strings.EqualFold(priv, d) {
			return true
		}
	}
	return false
}

// IsElevated checks if the current process is running with elevated privileges
func IsElevated() bool {
	if runtime.GOOS == "windows" {
		// Check if we can write to a protected registry key
		cmd := exec.Command("net", "session")
		err := cmd.Run()
		return err == nil
	}
	// Unix-like: check if UID is 0
	if out, err := exec.Command("id", "-u").Output(); err == nil {
		uid := strings.TrimSpace(string(out))
		return uid == "0"
	}
	return false
}
