package commands

type CommandHandler func(args []string) string

var Commands = map[string]CommandHandler{
	"info": InfoCommand,
	// "help": HelpCommand,
}

func InfoCommand(args []string) string {
	return "This is a custom reverse shell application.\n"
}
