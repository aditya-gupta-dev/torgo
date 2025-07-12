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
