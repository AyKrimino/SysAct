package main

import (
	"github.com/AyKrimino/SysAct/internal/config"
	"github.com/AyKrimino/SysAct/internal/logging"
	"github.com/AyKrimino/SysAct/internal/system"
	"github.com/AyKrimino/SysAct/internal/tui"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	logging.InitLogger()

	cfg, err := config.LoadConfig()
	if err != nil {
		logging.FatalLogger.Fatalf("Unable to load config %v", err)
	}

	if err := system.InitActions(); err != nil {
		logging.FatalLogger.Fatalf("Unable to initialize actions %v", err)
	}

	defaultLang := cfg.Languages[cfg.DefaultLanguage]

	l := []list.Item{
		tui.NewAction(defaultLang.LogoutLabel, defaultLang.LogoutDesc),
		tui.NewAction(defaultLang.SuspendLabel, defaultLang.SuspendDesc),
		tui.NewAction(defaultLang.RebootLabel, defaultLang.RebootDesc),
		tui.NewAction(defaultLang.PoweroffLabel, defaultLang.PoweroffDesc),
	}

	m := tui.NewMainModel(l, "System Actions", cfg)

	cm := tui.NewConfirmModel(cfg)

	rm := tui.NewRootModel(m, cm)

	p := tea.NewProgram(rm, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		logging.FatalLogger.Fatalf("Unable to run the app %v", err)
	}
}
