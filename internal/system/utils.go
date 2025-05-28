package system

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

var commands = map[string]string{
	"GNOME":      "gnome-session-quit --logout --no-prompt",
	"KDE Plasma": "qdbus org.kde.ksmserver /KSMServer logout 0 0 0",
	"XFCE":       "xfce4-session-logout --logout --fast",
	"MATE":       "mate-session-save --logout --force",
	"Cinnamon":   "cinnamon-session-quit --logout --no-prompt",
	"LXDE":       "lxde-logout",
	"LXQt":       "lxqt-leave --logout",
	"Unity":      "gnome-session-quit --logout --no-prompt",
	"i3":         "i3-msg exit",
}

func getLogoutCommandBasedOnDesktopEnvironmentOrWindowManager() (string, error) {
	if runtime.GOOS != "linux" {
		return "", fmt.Errorf("Can't execute this on a %s machine", runtime.GOOS)
	}

	output := os.Getenv("XDG_CURRENT_DESKTOP")

	if output == "" {
		return "", fmt.Errorf("Could not recognize which Window Manager/ Desktop Environment")
	}

	command := ""
	for key, val := range commands {
		if strings.Contains(strings.ToLower(output), strings.ToLower(key)) {
			command = val
		}
	}

	if command == "" {
		return "", fmt.Errorf("Could not find the Logout command for this Window Manager/Desktop Environment %s", output)
	}
	return command, nil
}
