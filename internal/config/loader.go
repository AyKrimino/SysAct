package config

import (
	"io"
	"os"
	"path/filepath"

	"github.com/AyKrimino/SysAct/internal/logging"
	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	DefaultLanguage string   `toml:"default_language"`
	ConfirmTimeout  int      `toml:"confirm_timeout"`
	KeyUp           []string `toml:"key_up"`
	KeyDown         []string `toml:"key_down"`
	KeyConfirm      []string `toml:"key_confirm"`
	KeyCancel       []string `toml:"key_cancel"`
	ShowHelp        bool     `toml:"show_help"`
	EnableLogging   bool     `toml:"enable_logging"`
	PrimaryColor    string   `toml:"primary_color"`
	SecondaryColor  string   `toml:"secondary_color"`
	HighlightStyle  string   `toml:"highlight_style"`

	Languages map[string]Locale `toml:"lang"`
}

type Locale struct {
	Title         string `toml:"title"`
	LogoutLabel   string `toml:"logout_label"`
	LogoutDesc    string `toml:"logout_desc"`
	SuspendLabel  string `toml:"suspend_label"`
	SuspendDesc   string `toml:"suspend_desc"`
	RebootLabel   string `toml:"reboot_label"`
	RebootDesc    string `toml:"reboot_desc"`
	PoweroffLabel string `toml:"poweroff_label"`
	PoweroffDesc  string `toml:"poweroff_desc"`

	ConfirmTitle string `toml:"confirm_title"`
	ConfirmBody  string `toml:"confirm_body"`
	ConfirmYes   string `toml:"confirm_yes"`
	ConfirmNo    string `toml:"confirm_no"`

	HelpText string `toml:"help_text"`
}

func LoadConfig() (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(homeDir, ".config", "sysact", "config.toml")

	file, err := os.Open(configPath)
	if err != nil && os.IsNotExist(err) {
		file, err = os.Open("./config.toml")
		if err != nil {
			return nil, err
		}
		logging.InfoLogger.Println("Using the default toml file")
	} else if err != nil {
		return nil, err
	}

	defer file.Close()

	var cfg *Config

	b, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = toml.Unmarshal(b, &cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
