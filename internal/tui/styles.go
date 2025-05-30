package tui

import "github.com/charmbracelet/lipgloss"

type Styles struct {
	docStyle         lipgloss.Style
	titleStyle       lipgloss.Style
	descriptionStyle lipgloss.Style
	timerStyle       lipgloss.Style
}

func NewMainModelStyles() Styles {
	return Styles{
		docStyle: lipgloss.NewStyle().
			Margin(2, 4),
	}
}

func NewConfirmModelStyles() Styles {
	return Styles{
		titleStyle: lipgloss.NewStyle().
			Align(lipgloss.Center).
			Bold(true).
			Foreground(lipgloss.Color("#FFD700")).
			Background(lipgloss.Color("#1E1E1E")).
			MarginBottom(1),

		descriptionStyle: lipgloss.NewStyle().
			Align(lipgloss.Center).
			Foreground(lipgloss.Color("#FAFAFA")).
			Italic(true).
			MarginBottom(1),

		timerStyle: lipgloss.NewStyle().
			Align(lipgloss.Center).
			Foreground(lipgloss.Color("#FF5F5F")).
			Bold(true),

		docStyle: lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#7D56F4")).
			Padding(1, 2).
			Align(lipgloss.Center),
	}
}
