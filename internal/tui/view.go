package tui

func (m MainModel) View() string {
	return m.styles.docStyle.Render(m.list.View())
}
