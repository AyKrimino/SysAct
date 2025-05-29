package tui

import (
	"github.com/AyKrimino/SysAct/internal/logging"
	tea "github.com/charmbracelet/bubbletea"
)

func (cm ConfirmModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd = nil
	switch msg := msg.(type) {
	case tea.KeyMsg:
		logging.InfoLogger.Printf("Confirm Model: key %s pressed", msg.String())
		switch msg.String() {
		case "y", "Y":
			logging.InfoLogger.Println("confirmed")
			cmd = func() tea.Msg { return ConfirmResultMsg{confirmed: true} }
			return cm, cmd
		default:
			logging.InfoLogger.Println("rejected")
			cmd = func() tea.Msg { return ConfirmResultMsg{confirmed: false} }
			return cm, cmd
		}
	}
	return cm, cmd
}
