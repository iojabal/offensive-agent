package commands

import (
	info "nombredetuapp/Documents/Proyecto/src/commands/Info"
	"nombredetuapp/Documents/Proyecto/src/commands/persistence"
	"nombredetuapp/Documents/Proyecto/src/commands/recon"
)

type CommandHandler func(args []string) string

var Commands = map[string]CommandHandler{
	"info":        info.InfoCommand,
	"persistence": persistence.PersistenceCommand,
	"recon":       recon.ReconCommand,
	// "help": HelpCommand,
}
