package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MainModel struct {
	list      list.Model
	styles    Styles
	listIndex uint8
}

func NewMainModel(actions []list.Item, title string) MainModel {
	m := MainModel{
		list: list.New(actions, list.NewDefaultDelegate(), 0, 0),
	}
	m.list.Title = title
	m.styles = Styles{
		docStyle: lipgloss.NewStyle().Margin(1, 2),
	}

	return m
}

func (m MainModel) Init() tea.Cmd {
	return nil
}
