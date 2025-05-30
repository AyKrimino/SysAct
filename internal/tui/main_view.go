package tui

import "github.com/charmbracelet/lipgloss"

func (m MainModel) View() string {
	title := m.styles.titleStyle.Render("SysAct — System Actions")

	content := m.styles.docStyle.Render(m.list.View())

	dialog := lipgloss.JoinVertical(lipgloss.Left, title, content)

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center, lipgloss.Center,
		dialog,
	)
}
