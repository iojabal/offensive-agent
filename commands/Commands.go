package commands

import (
	info "nombredetuapp/Documents/Proyecto/src/commands/Info"
	"nombredetuapp/Documents/Proyecto/src/commands/persistence"
)

type CommandHandler func(args []string) string

var Commands = map[string]CommandHandler{
	"info":        info.InfoCommand,
	"persistence": persistence.PersistenceCommand,
	// "help": HelpCommand,
}
