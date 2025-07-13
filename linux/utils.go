package linux

import (
	"os/exec"
	"strings"
)

func CheckTorInstallation() (string, error) {
	output, err := exec.Command("which", "tor").Output()
	return string(output), err
}

func CheckTorStatus() (bool, error) {
	cmd := exec.Command("systemctl", "is-active", "tor")
	output, err := cmd.Output()
	if err != nil {
		return false, err
	}

	status := strings.TrimSpace(string(output))

	return status == "active", nil
}

func KillTor() error {
	cmd := exec.Command("pkill", "-f", "tor")
	_, err := cmd.Output()
	if err != nil {
		return err
	}
	return nil
}

func RestartTor(torPath string) error {
	status, err := CheckTorStatus()
	if err != nil {
		return err
	}

	if status {
		if err = KillTor(); err != nil {
			return err
		}
	}

	cmd := exec.Command(torPath)
	_, err = cmd.Output()

	if err != nil {
		return err
	}
	return nil
}
