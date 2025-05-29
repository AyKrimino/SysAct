package tui

import tea "github.com/charmbracelet/bubbletea"

type ConfirmModel struct {
	action uint8
}

func (cm ConfirmModel) Init() tea.Cmd {
	return nil
}

type ConfirmRequestedMsg struct {
	Action uint8
}

type ConfirmResultMsg struct {
	confirmed bool
}
