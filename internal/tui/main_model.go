package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type MainModel struct {
	list      list.Model
	styles    Styles
	listIndex uint8
	width     int
	height    int
}

func NewMainModel(actions []list.Item, title string) MainModel {
	m := MainModel{
		list: list.New(actions, list.NewDefaultDelegate(), 0, 0),
	}
	m.list.Title = title
	m.styles = NewMainModelStyles()
	m.list.SetShowHelp(false)
	m.list.SetFilteringEnabled(false)

	return m
}

func (m MainModel) Init() tea.Cmd {
	return nil
}
