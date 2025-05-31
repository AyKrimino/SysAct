package tui

import "github.com/charmbracelet/lipgloss"

func (m MainModel) View() string {
	defaultLang := m.config.Languages[m.config.DefaultLanguage]

	title := m.styles.titleStyle.Render(defaultLang.Title)

	content := m.styles.docStyle.Render(m.list.View())

	dialog := lipgloss.JoinVertical(lipgloss.Left, title, content)

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center, lipgloss.Center,
		dialog,
	)
}
