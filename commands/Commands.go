package commands

import info "nombredetuapp/Documents/Proyecto/src/commands/Info"

type CommandHandler func(args []string) string

var Commands = map[string]CommandHandler{
	"info": info.InfoCommand,
	// "help": HelpCommand,
}
