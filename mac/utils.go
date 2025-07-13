package mac

import (
	"fmt"
	"os/exec"
	"path"
	"strings"

	"github.com/aditya-gupta-dev/torgo/utils"
)

func CheckTorInstallation() (string, error) {
	macPaths := []string{
		"/Applications/TorBrowser.app/Contents/MacOS/Tor/tor",
		"/usr/local/bin/tor",
		"/opt/homebrew/bin/tor",
		"/usr/bin/tor",
		"/Applications/TorBrowser.app/Contents/Resources/TorBrowser/Tor/tor",
	}

	for _, path := range macPaths {
		if utils.FileExists(path) {
			return path, nil
		}
	}

	torPath, err := exec.LookPath("tor")
	if err == nil {
		return torPath, nil
	}

	output, err := exec.Command("which", "tor").Output()
	if err == nil {
		return string(output), nil
	}

	return "", fmt.Errorf("tor not found")
}

func CheckTorStatus(torPath string) (bool, error) {
	cmd := exec.Command("pgrep", path.Base(torPath))
	output, err := cmd.Output()
	if err != nil {
		return false, err
	}
	outputStr := string(output)
	return strings.TrimSpace(outputStr) != "", nil
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
