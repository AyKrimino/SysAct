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
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			m.Up()
		case "down", "j":
			m.Down()
		case "enter", " ":
			command, err := system.PerformAction(m.listIndex)
			if err != nil {
				logging.ErrorLogger.Fatalf("Command execution failed: %v", err)
			}
			logging.InfoLogger.Printf("Command %s executed successfully", command)
		}
	case tea.WindowSizeMsg:
		logging.InfoLogger.Printf("Window resized: width=%d, height=%d", msg.Width, msg.Height)
		h, v := m.styles.docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	return m, cmd
}
