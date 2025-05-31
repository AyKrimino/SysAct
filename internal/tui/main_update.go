package tui

import (
	"github.com/AyKrimino/SysAct/internal/logging"
	"github.com/AyKrimino/SysAct/internal/system"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *MainModel) Up() {
	if m.listIndex == 0 {
		m.listIndex = system.Poweroff
	} else {
		m.listIndex--
	}
	m.list.Select(int(m.listIndex) + 1)
}

func (m *MainModel) Down() {
	if m.listIndex == system.Poweroff {
		m.listIndex = system.Logout
	} else {
		m.listIndex++
	}
	m.list.Select(int(m.listIndex) - 1)
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		logging.InfoLogger.Println("Key pressed:", msg.String())
		switch msg.String() {
		case m.config.KeyQuit:
			return m, tea.Quit
		case m.config.KeyUp:
			m.Up()
		case m.config.KeyDown:
			m.Down()
		case m.config.KeySelect:
			var cmd tea.Cmd
			cmd = func() tea.Msg { return ConfirmRequestedMsg{Action: m.listIndex} }
			return m, cmd
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		logging.InfoLogger.Printf("Window resized: width=%d, height=%d", m.width, m.height)
		listWidth := int(float64(msg.Width) * 0.5)
		listHeight := int(float64(msg.Height) * 0.6)
		m.list.SetSize(listWidth, listHeight)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	return m, cmd
}
