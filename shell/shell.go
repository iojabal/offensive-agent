package shell

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type Shell struct {
	osUsed    string
	shellPath string
	shellArgs []string
	cwd       string
}

func NewShell() *Shell {
	osUsed := runtime.GOOS
	var shellPath string
	var shellArgs []string
	if osUsed == "windows" {
		shellPath = "powershell.exe"
		shellArgs = []string{"-NoLogo", "-NoProfile", "-Command"}
	} else {
		shellPath = "/bin/sh"
		shellArgs = []string{"-c"}
	}

	cwd, _ := os.Getwd()
	return &Shell{
		osUsed:    osUsed,
		shellPath: shellPath,
		shellArgs: shellArgs,
		cwd:       cwd,
	}
}

func (s *Shell) Execute(input string) string {
	cmdStr := strings.TrimSpace(input)

	if strings.HasPrefix(cmdStr, "cd ") {
		newDir := strings.TrimSpace(cmdStr[3:])

		if newDir == "" {
			if s.osUsed == "windows" {
				newDir = os.Getenv("USERPROFILE")
			} else {
				newDir = os.Getenv("HOME")
			}
		}
		err := os.Chdir(newDir)
		if err != nil {
			return fmt.Sprintf("cd: %s: %s\n", newDir, err.Error())
		}
		s.cwd, _ = os.Getwd()
		return ""
	}
	fullCmd := append(s.shellArgs, cmdStr)
	cmd := exec.Command(s.shellPath, fullCmd...)
	cmd.Dir = s.cwd

	output, err := cmd.CombinedOutput()
	result := string(output)

	if err != nil {
		result += fmt.Sprintf("\n[!] execution error: %v", err)
	}

	return strings.ReplaceAll(result, "\n", "\r\n")
}
func (s *Shell) GetCWD() string {
	return s.cwd
}
