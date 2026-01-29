package info

import (
	"nombredetuapp/Documents/Proyecto/src/shell"
	"os"
	"runtime"
)

type InfoOutput struct {
	Os   string
	Arch string
	User string
	CWD  string
	PID  int
}

func NewInfoOutput() *InfoOutput {
	username := os.Getenv("USERNAME")
	if username == "" {
		username = os.Getenv("USER")
	}
	return &InfoOutput{
		Os:   runtime.GOOS,
		Arch: runtime.GOARCH,
		User: username,
		CWD:  shell.NewShell().GetCWD(),
		PID:  os.Getpid(),
	}
}

func InfoCommand(args []string) string {

	info := NewInfoOutput()
	output := "\nSystem Information:\n"
	output += "-------------------\n"
	output += "Operating System: " + info.Os + "\n"
	output += "Architecture: " + info.Arch + "\n"
	output += "Current User: " + info.User + "\n"
	output += "Current Directory: " + info.CWD + "\n"
	output += "Process ID: " + string(rune(info.PID)) + "\n\n"
	return output

}
