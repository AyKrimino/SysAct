package tui

import (
	"github.com/AyKrimino/SysAct/internal/logging"
	"github.com/charmbracelet/lipgloss"
)

func (cm ConfirmModel) View() string {
	defaultLang := cm.config.Languages[cm.config.DefaultLanguage]

	title := cm.styles.titleStyle.Render(defaultLang.ConfirmTitle)

	logging.InfoLogger.Println(defaultLang.ConfirmBody)

	description := cm.styles.descriptionStyle.Render(defaultLang.ConfirmBody)
	timer := cm.styles.timerStyle.Render(defaultLang.CancelBody + "  " + cm.timer.View())

	dialog := lipgloss.JoinVertical(
		lipgloss.Center,
		title,
		description,
		timer,
	)

	return cm.styles.docStyle.Render(
		lipgloss.Place(
			cm.width,
			cm.height,
			lipgloss.Center,
			lipgloss.Center,
			dialog,
		),
	)
}
