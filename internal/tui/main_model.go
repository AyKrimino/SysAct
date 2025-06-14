package tui

import (
	"github.com/AyKrimino/SysAct/internal/config"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type MainModel struct {
	list      list.Model
	styles    Styles
	listIndex uint8
	width     int
	height    int
	config    *config.Config
}

func NewMainModel(actions []list.Item, title string, config *config.Config) MainModel {
	m := MainModel{
		list:   list.New(actions, list.NewDefaultDelegate(), 0, 0),
		config: config,
	}
	m.list.Title = title
	m.styles = NewMainModelStyles(config)
	m.list.SetShowHelp(m.config.ShowHelp)
	m.list.SetFilteringEnabled(false)

	m.list.KeyMap.CursorUp = key.NewBinding(
		key.WithKeys(config.KeyUp),
		key.WithHelp(config.KeyUp, "move up"),
	)

	m.list.KeyMap.CursorDown = key.NewBinding(
		key.WithKeys(config.KeyDown),
		key.WithHelp(config.KeyDown, "move down"),
	)

	m.list.KeyMap.Quit = key.NewBinding(
		key.WithKeys(config.KeyQuit),
		key.WithHelp(config.KeyQuit, "quit"),
	)

	return m
}

func (m MainModel) Init() tea.Cmd {
	return nil
}
