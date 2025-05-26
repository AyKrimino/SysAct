package tui

import (
	"github.com/AyKrimino/SysAct/internal/logging"
	tea "github.com/charmbracelet/bubbletea"
)

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		logging.InfoLogger.Println("Key pressed:", msg.String())
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
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
