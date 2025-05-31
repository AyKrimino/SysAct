package tui

import (
	"github.com/AyKrimino/SysAct/internal/logging"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
)

func (cm ConfirmModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case timer.TickMsg:
		cm.timer, cmd = cm.timer.Update(msg)
		return cm, cmd
	case tea.KeyMsg:
		logging.InfoLogger.Printf("Confirm Model: key %s pressed", msg.String())
		switch msg.String() {
		case cm.config.KeyConfirm:
			logging.InfoLogger.Println("confirmed")
			cmd = func() tea.Msg { return ConfirmResultMsg{confirmed: true} }
			return cm, cmd
		case cm.config.KeyCancel:
			logging.InfoLogger.Println("rejected")
			cmd = func() tea.Msg { return ConfirmResultMsg{confirmed: false} }
			return cm, cmd
		}
	case tea.WindowSizeMsg:
		cm.width = msg.Width
		cm.height = msg.Height
		logging.InfoLogger.Printf("Window resized: width=%d, height=%d", cm.width, cm.height)
	}
	return cm, nil
}
