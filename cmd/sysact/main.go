package main

import (
	"github.com/AyKrimino/SysAct/internal/logging"
	"github.com/AyKrimino/SysAct/internal/tui"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	logging.InitLogger()

	l := []list.Item{
		tui.NewAction("Logout", "End the current user session and return to the login screen."),
		tui.NewAction("Suspend", "Pause the system, saving the session to memory and entering a low-power state."),
		tui.NewAction("Reboot", "Restart the system, closing all applications and reinitializing the operating system."),
		tui.NewAction("Poweroff", "Completely shut down the system, turning off all hardware components."),
	}

	m := tui.NewMainModel(l, "System Actions")

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		logging.FatalLogger.Fatalf("Unable to run the app %v", err)
	}
}
