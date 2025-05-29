package tui

func (cm ConfirmModel) View() string {
	s := "Confirm your action by pressing the 'y' key on you keyboard\n"

	s += "Canceling in " + cm.timer.View()
	return s
}
