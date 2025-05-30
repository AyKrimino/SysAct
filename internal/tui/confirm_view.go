package tui

import "github.com/charmbracelet/lipgloss"

func (cm ConfirmModel) View() string {
	title := cm.styles.titleStyle.Render("System Action Confirmation")
	description := cm.styles.descriptionStyle.Render("Confirm your action by pressing the 'y' key.")
	timer := cm.styles.timerStyle.Render("Canceling in " + cm.timer.View())

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
