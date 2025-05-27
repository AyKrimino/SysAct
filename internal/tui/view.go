package tui

import "github.com/charmbracelet/lipgloss"

func (m MainModel) View() string {
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Left,
			"",
			m.styles.docStyle.Render(m.list.View()),
		),
	)
}
