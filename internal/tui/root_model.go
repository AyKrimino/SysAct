package tui

import (
	"github.com/AyKrimino/SysAct/internal/logging"
	"github.com/AyKrimino/SysAct/internal/system"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
)

type Screen int

const (
	MainModelScreen Screen = iota
	ConfirmModelScreen
)

type RootModel struct {
	screen  Screen
	main    MainModel
	confirm ConfirmModel
}

func NewRootModel(mainModel MainModel, confirmModel ConfirmModel) RootModel {
	return RootModel{
		screen:  MainModelScreen,
		main:    mainModel,
		confirm: confirmModel,
	}
}

func (rm RootModel) Init() tea.Cmd {
	return rm.main.Init()
}

func (rm RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case ConfirmRequestedMsg:
		logging.InfoLogger.Printf("action %d requested", msg.Action)
		rm.confirm = NewConfirmModel()
		rm.confirm.action = msg.Action
		rm.screen = ConfirmModelScreen
		return rm, rm.confirm.Init()
	case ConfirmResultMsg:
		if msg.confirmed == true {
			command, err := system.PerformAction(rm.confirm.action)
			if err != nil {
				logging.ErrorLogger.Fatalf("Command execution failed: %v", err)
			}
			logging.InfoLogger.Printf("Command %s executed successfully", command)
			return rm, tea.Quit
		}
		rm.screen = MainModelScreen
		return rm, rm.main.Init()
	case timer.TimeoutMsg:
		logging.InfoLogger.Println("Timer timed out")
		rm.screen = MainModelScreen
		return rm, rm.main.Init()
	}

	switch rm.screen {
	case MainModelScreen:
		updated, cmd := rm.main.Update(msg)
		rm.main = updated.(MainModel)
		return rm, cmd
	case ConfirmModelScreen:
		updated, cmd := rm.confirm.Update(msg)
		rm.confirm = updated.(ConfirmModel)
		return rm, cmd
	}
	return rm, nil
}

func (rm RootModel) View() string {
	switch rm.screen {
	case MainModelScreen:
		return rm.main.View()
	case ConfirmModelScreen:
		return rm.confirm.View()
	}
	return ""
}
