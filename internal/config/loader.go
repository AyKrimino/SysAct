package config

import (
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	DefaultLanguage string `toml:"default_language"`
	ConfirmTimeout  int    `toml:"confirm_timeout"`
	KeyUp           string `toml:"key_up"`
	KeyDown         string `toml:"key_down"`
	KeyConfirm      string `toml:"key_confirm"`
	KeyCancel       string `toml:"key_cancel"`
	KeyQuit         string `toml:"key_quit"`
	KeySelect       string `toml:"key_select"`
	ShowHelp        bool   `toml:"show_help"`
	EnableLogging   bool   `toml:"enable_logging"`
	PrimaryColor    string `toml:"primary_color"`
	SecondaryColor  string `toml:"secondary_color"`

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
	CancelBody   string `toml:"cancel_body"`
}

func initDefaultConfig() *Config {
	return &Config{
		DefaultLanguage: "en",
		ConfirmTimeout:  5,
		KeyUp:           "up",
		KeyDown:         "down",
		KeyConfirm:      "enter",
		KeyCancel:       "esc",
		KeyQuit:         "q",
		KeySelect:       "enter",
		ShowHelp:        false,
		EnableLogging:   true,
		PrimaryColor:    "#87CEEB",
		SecondaryColor:  "FFD700",
		Languages: map[string]Locale{
			"en": {
				Title:         "SysAct - System Actions",
				LogoutLabel:   "Logout",
				LogoutDesc:    "End the current user session and return to the login screen.",
				SuspendLabel:  "Suspend",
				SuspendDesc:   "Pause the system, saving the session to memory and entering a low-power state.",
				RebootLabel:   "Reboot",
				RebootDesc:    "Restart the system, closing all applications and reinitializing the operating system.",
				PoweroffLabel: "Poweroff",
				PoweroffDesc:  "Completely shut down the system, turning off all hardware components.",
				ConfirmTitle:  "Are you sure?",
				ConfirmBody:   "Press {{key_confirm}} to confirm",
				CancelBody:    "Canceling in",
			},
			"fr": {
				Title:         "SysAct - Actions Système",
				LogoutLabel:   "Se déconnecter",
				LogoutDesc:    "Terminez la session en cours et revenez à l'écran de connexion.",
				SuspendLabel:  "Suspendre",
				SuspendDesc:   "Mettez le système en veille, en enregistrant la session en mémoire.",
				RebootLabel:   "Redémarrer",
				RebootDesc:    "Redémarrez le système, en fermant toutes les applications.",
				PoweroffLabel: "Arrêter",
				PoweroffDesc:  "Arrêtez complètement le système et éteignez le matériel.",
				ConfirmTitle:  "Êtes-vous sûr ?",
				ConfirmBody:   "Appuyez sur {{key_confirm}} pour confirmer",
				CancelBody:    "Annulation dans",
			},
			"es": {
				Title:         "SysAct - Acciones del Sistema",
				LogoutLabel:   "Cerrar sesión",
				LogoutDesc:    "Finaliza la sesión actual y vuelve a la pantalla de inicio de sesión.",
				SuspendLabel:  "Suspender",
				SuspendDesc:   "Suspende el sistema, guardando la sesión en la memoria.",
				RebootLabel:   "Reiniciar",
				RebootDesc:    "Reinicia el sistema, cerrando todas las aplicaciones.",
				PoweroffLabel: "Apagar",
				PoweroffDesc:  "Apaga completamente el sistema, apagando el hardware.",
				ConfirmTitle:  "¿Estás seguro?",
				ConfirmBody:   "Pulsa {{key_confirm}} para confirmar",
				CancelBody:    "Cancelando en",
			},
			"ar": {
				Title:         "SysAct - إجراءات النظام",
				LogoutLabel:   "تسجيل الخروج",
				LogoutDesc:    "إنهاء الجلسة الحالية والعودة إلى شاشة تسجيل الدخول.",
				SuspendLabel:  "تعليق",
				SuspendDesc:   "تعليق النظام وحفظ الجلسة في الذاكرة.",
				RebootLabel:   "إعادة التشغيل",
				RebootDesc:    "إعادة تشغيل النظام، وإغلاق جميع التطبيقات.",
				PoweroffLabel: "إيقاف التشغيل",
				PoweroffDesc:  "إيقاف النظام بالكامل وإيقاف تشغيل الأجهزة.",
				ConfirmTitle:  "هل أنت متأكد؟",
				ConfirmBody:   "اضغط {{key_confirm}} للتأكيد",
				CancelBody:    "سيتم الإلغاء خلال",
			},
		},
	}
}

func LoadConfig() (*Config, error) {
	defaultCfg := initDefaultConfig()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(homeDir, ".config", "sysact", "config.toml")

	file, err := os.Open(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return defaultCfg, nil
		}
		return nil, err
	}
	defer file.Close()

	var userCfg Config

	b, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = toml.Unmarshal(b, &userCfg)
	if err != nil {
		return nil, err
	}

	if loc, ok := userCfg.Languages[userCfg.DefaultLanguage]; ok {
		raw := loc.ConfirmBody
		keyName := userCfg.KeyConfirm
		filled := strings.ReplaceAll(raw, "{{key_confirm}}", keyName)
		loc.ConfirmBody = filled
		userCfg.Languages[userCfg.DefaultLanguage] = loc
	}

	mergeConfigs(&userCfg, defaultCfg)

	return defaultCfg, nil
}

func mergeConfigs(userCfg, defaultCfg *Config) {
	if userCfg.DefaultLanguage != "" {
		if _, ok := defaultCfg.Languages[userCfg.DefaultLanguage]; ok {
			defaultCfg.DefaultLanguage = userCfg.DefaultLanguage
		}
	}

	if userCfg.ConfirmTimeout > 0 && userCfg.ConfirmTimeout <= 60 {
		defaultCfg.ConfirmTimeout = userCfg.ConfirmTimeout
	}

	if userCfg.KeyUp != "" {
		defaultCfg.KeyUp = userCfg.KeyUp
	}
	if userCfg.KeyDown != "" {
		defaultCfg.KeyDown = userCfg.KeyDown
	}
	if userCfg.KeyConfirm != "" {
		defaultCfg.KeyConfirm = userCfg.KeyConfirm
	}
	if userCfg.KeyCancel != "" {
		defaultCfg.KeyCancel = userCfg.KeyCancel
	}
	if userCfg.KeyQuit != "" {
		defaultCfg.KeyQuit = userCfg.KeyQuit
	}
	if userCfg.KeySelect != "" {
		defaultCfg.KeySelect = userCfg.KeySelect
	}

	defaultCfg.ShowHelp = userCfg.ShowHelp
	defaultCfg.EnableLogging = userCfg.EnableLogging

	if userCfg.PrimaryColor != "" {
		if match, _ := regexp.MatchString("^#[0-9A-Fa-f]{6}$", userCfg.PrimaryColor); match {
			defaultCfg.PrimaryColor = userCfg.PrimaryColor
		}
	}
	if userCfg.SecondaryColor != "" {
		if match, _ := regexp.MatchString("^#[0-9A-Fa-f]{6}$", userCfg.PrimaryColor); match {
			defaultCfg.SecondaryColor = userCfg.SecondaryColor
		}
	}

	for langKey, userLocale := range userCfg.Languages {
		if defaultLocale, ok := defaultCfg.Languages[langKey]; ok {
			if userLocale.Title != "" {
				defaultLocale.Title = userLocale.Title
			}
			if userLocale.LogoutLabel != "" {
				defaultLocale.LogoutLabel = userLocale.LogoutLabel
			}
			if userLocale.LogoutDesc != "" {
				defaultLocale.LogoutDesc = userLocale.LogoutDesc
			}
			if userLocale.SuspendLabel != "" {
				defaultLocale.SuspendLabel = userLocale.SuspendLabel
			}
			if userLocale.SuspendDesc != "" {
				defaultLocale.SuspendDesc = userLocale.SuspendDesc
			}
			if userLocale.RebootLabel != "" {
				defaultLocale.RebootLabel = userLocale.RebootLabel
			}
			if userLocale.RebootDesc != "" {
				defaultLocale.RebootDesc = userLocale.RebootDesc
			}
			if userLocale.PoweroffLabel != "" {
				defaultLocale.PoweroffLabel = userLocale.PoweroffLabel
			}
			if userLocale.PoweroffDesc != "" {
				defaultLocale.PoweroffDesc = userLocale.PoweroffDesc
			}
			if userLocale.ConfirmTitle != "" {
				defaultLocale.ConfirmTitle = userLocale.ConfirmTitle
			}
			if userLocale.ConfirmBody != "" {
				defaultLocale.ConfirmBody = userLocale.ConfirmBody
			}
			if userLocale.CancelBody != "" {
				defaultLocale.CancelBody = userLocale.CancelBody
			}

			defaultCfg.Languages[langKey] = defaultLocale
		}
	}
}
