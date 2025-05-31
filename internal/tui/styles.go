package tui

import (
	"github.com/AyKrimino/SysAct/internal/config"
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	docStyle         lipgloss.Style
	titleStyle       lipgloss.Style
	descriptionStyle lipgloss.Style
	timerStyle       lipgloss.Style
}

func NewMainModelStyles(cfg *config.Config) Styles {
	return Styles{
		titleStyle: lipgloss.NewStyle().
			Align(lipgloss.Center).
			Bold(true).
			Foreground(lipgloss.Color(cfg.PrimaryColor)).
			Margin(0, 4).
			MarginBottom(1),

		docStyle: lipgloss.NewStyle().
			Margin(1, 4).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color(cfg.SecondaryColor)).
			Padding(1, 2),
	}
}

func NewConfirmModelStyles(cfg *config.Config) Styles {
	return Styles{
		titleStyle: lipgloss.NewStyle().
			Align(lipgloss.Center).
			Bold(true).
			Foreground(lipgloss.Color(cfg.SecondaryColor)).
			Background(lipgloss.Color(cfg.PrimaryColor)).
			MarginBottom(1),

		descriptionStyle: lipgloss.NewStyle().
			Align(lipgloss.Center).
			Foreground(lipgloss.Color(cfg.SecondaryColor)).
			Italic(true).
			MarginBottom(1),

		timerStyle: lipgloss.NewStyle().
			Align(lipgloss.Center).
			Foreground(lipgloss.Color("#FF5F5F")).
			Bold(true),

		docStyle: lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(cfg.PrimaryColor)).
			Padding(1, 2).
			Align(lipgloss.Center),
	}
}
