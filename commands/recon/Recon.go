package recon

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// ReconCommand executes a comprehensive system reconnaissance
// Usage: recon [--quick|--full|--json]
func ReconCommand(args []string) string {
	mode := "standard"
	outputJSON := false

	// Parse arguments
	if len(args) > 0 {
		switch args[0] {
		case "--quick":
			mode = "quick"
		case "--full":
			mode = "full"
		case "--json":
			outputJSON = true
		case "--help":
			return helpText()
		}
	}

	// Gather information concurrently with timeout
	results := make(map[string]string)
	var wg sync.WaitGroup
	var mu sync.Mutex

	timeout := 30 * time.Second
	ctx := make(chan struct{})

	// Launch concurrent reconnaissance
	tasks := []struct {
		name string
		fn   func() string
	}{
		{"system", GatherSystemInfo},
		{"security", GatherSecurityContext},
		{"network", GatherNetworkInfo},
		{"defense", GatherDefenseInfo},
	}

	for _, task := range tasks {
		wg.Add(1)
		go func(taskName string, taskFn func() string) {
			defer wg.Done()
			select {
			case <-ctx:
				return
			default:
				result := taskFn()
				mu.Lock()
				results[taskName] = result
				mu.Unlock()
			}
		}(task.name, task.fn)
	}

	// Wait with timeout
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		// All tasks completed
	case <-time.After(timeout):
		close(ctx)
		return "[!] Reconnaissance timed out after 30 seconds\n"
	}

	// Format output
	if outputJSON {
		return formatJSON(results)
	}

	return formatOutput(results, mode)
}

func formatOutput(results map[string]string, mode string) string {
	var output strings.Builder

	output.WriteString("\n")
	output.WriteString("═══════════════════════════════════════════════════════════\n")
	output.WriteString("                 SYSTEM RECONNAISSANCE                      \n")
	output.WriteString("═══════════════════════════════════════════════════════════\n\n")

	// System Information
	if sysInfo, ok := results["system"]; ok {
		output.WriteString("┌─ SYSTEM INFORMATION\n")
		output.WriteString("│\n")
		output.WriteString(sysInfo)
		output.WriteString("\n")
	}

	// Security Context
	if secInfo, ok := results["security"]; ok {
		output.WriteString("┌─ SECURITY CONTEXT\n")
		output.WriteString("│\n")
		output.WriteString(secInfo)
		output.WriteString("\n")
	}

	// Network Information
	if netInfo, ok := results["network"]; ok {
		output.WriteString("┌─ NETWORK INFORMATION\n")
		output.WriteString("│\n")
		output.WriteString(netInfo)
		output.WriteString("\n")
	}

	// Defense Mechanisms (only in full mode)
	if mode == "full" {
		if defInfo, ok := results["defense"]; ok {
			output.WriteString("┌─ DEFENSE MECHANISMS\n")
			output.WriteString("│\n")
			output.WriteString(defInfo)
			output.WriteString("\n")
		}
	}

	output.WriteString("═══════════════════════════════════════════════════════════\n")

	return output.String()
}

func formatJSON(results map[string]string) string {
	// Simple JSON formatting (you could use encoding/json for proper formatting)
	var output strings.Builder
	output.WriteString("{\n")
	output.WriteString("  \"timestamp\": \"" + time.Now().Format(time.RFC3339) + "\",\n")
	output.WriteString("  \"recon\": {\n")

	keys := []string{"system", "security", "network", "defense"}
	for i, key := range keys {
		if value, ok := results[key]; ok {
			output.WriteString(fmt.Sprintf("    \"%s\": %q", key, value))
			if i < len(keys)-1 {
				output.WriteString(",\n")
			} else {
				output.WriteString("\n")
			}
		}
	}

	output.WriteString("  }\n")
	output.WriteString("}\n")

	return output.String()
}

func helpText() string {
	return `
Reconnaissance Command - Gather system information

Usage: recon [OPTIONS]

Options:
  --quick    Quick reconnaissance (essential info only)
  --full     Full reconnaissance (includes defense mechanisms)
  --json     Output in JSON format
  --help     Show this help message

Examples:
  recon              # Standard reconnaissance
  recon --quick      # Fast essential info
  recon --full       # Complete system analysis
  recon --json       # JSON output for parsing

Information Gathered:
  • System: OS, architecture, hostname, uptime
  • Security: User, groups, privileges, domain membership
  • Network: IPs, routes, DNS, domain controllers
  • Defense: AV, firewall, EDR, logging (--full only)
`
}