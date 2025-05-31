package tui

import (
	"time"

	"github.com/AyKrimino/SysAct/internal/config"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
)

const timeout = time.Second * 5

type ConfirmModel struct {
	action uint8
	timer  timer.Model
	styles Styles
	width  int
	height int
	config *config.Config
}

func NewConfirmModel(config *config.Config) ConfirmModel {
	return ConfirmModel{
		timer:  timer.NewWithInterval(timeout, time.Millisecond),
		styles: NewConfirmModelStyles(),
		config: config,
	}
}

func (cm ConfirmModel) Init() tea.Cmd {
	return cm.timer.Init()
}

type ConfirmRequestedMsg struct {
	Action uint8
}

type ConfirmResultMsg struct {
	confirmed bool
}
