package system

import (
	"fmt"
	"os/exec"
	"runtime"
)

const (
	Logout uint8 = iota
	Suspend
	Reboot
	Poweroff
)

var actions map[uint8]string

func InitActions() error {
	var (
		err error
		cmd string
	)

	cmd, err = getLogoutCommandBasedOnDesktopEnvironmentOrWindowManager()
	if err != nil {
		return err
	}

	actions = map[uint8]string{
		Logout:   cmd,
		Suspend:  "systemctl suspend",
		Reboot:   "systemctl reboot",
		Poweroff: "systemctl poweroff",
	}
	return nil
}

func execute(command string) error {
	_, err := exec.Command(command).Output()
	if err != nil {
		return fmt.Errorf("Failed to execute command %s, %v", command, err)
	}
	return nil
}

func PerformAction(action uint8) error {
	if runtime.GOOS != "linux" {
		return fmt.Errorf("Can't execute this on a %s machine", runtime.GOOS)
	}

	err := execute(actions[action])
	return err
}
