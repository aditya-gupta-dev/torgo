package windows

import (
	"fmt"
	"os/exec"
	"path"
	"strings"

	"github.com/aditya-gupta-dev/torgo/utils"
)

func CheckTorInstallation() (string, error) {

	windowsPaths := []string{
		`C:\Program Files\Tor Browser\Browser\TorBrowser\Tor\tor.exe`,
		`C:\Program Files (x86)\Tor Browser\Browser\TorBrowser\Tor\tor.exe`,
		`C:\Users\%USERNAME%\Desktop\Tor Browser\Browser\TorBrowser\Tor\tor.exe`,
		`C:\Users\%USERNAME%\AppData\Local\Tor Browser\Browser\TorBrowser\Tor\tor.exe`,
		`C:\Tor\tor.exe`,
	}

	for _, path := range windowsPaths {
		if utils.FileExists(path) {
			return path, nil
		}
	}

	return exec.LookPath("tor.exe")
}

func CheckTorStatus(torPath string) (bool, error) {
	cmd := exec.Command("tasklist", "/FI", fmt.Sprintf("IMAGENAME eq %s", path.Base(torPath)))
	output, err := cmd.Output()
	if err != nil {
		return false, err
	}

	outputStr := string(output)
	return strings.Contains(outputStr, "tor"), nil
}

func KillTor(torPath string) error {
	cmd := exec.Command("taskkill", "/F", "/IM", path.Base(torPath))
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
