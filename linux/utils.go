package linux

import (
	"os/exec"
	"path"
	"strings"
)

func CheckTorInstallation() (string, error) {
	output, err := exec.Command("which", "tor").Output()
	return string(output), err
}

func CheckTorStatus(torPath string) (bool, error) {
	cmd := exec.Command("systemctl", "is-active", path.Base(torPath))
	output, err := cmd.Output()
	if err != nil {
		return false, err
	}

	status := strings.TrimSpace(string(output))

	return status == "active", nil
}

func KillTor(torPath string) error {
	cmd := exec.Command("pkill", "-f", path.Base(torPath))
	_, err := cmd.Output()
	if err != nil {
		return err
	}
	return nil
}

func RestartTor(torPath string) error {
	status, err := CheckTorStatus(torPath)
	if err != nil {
		return err
	}

	if status {
		if err = KillTor(torPath); err != nil {
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
