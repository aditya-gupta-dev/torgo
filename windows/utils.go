package windows

import (
	"os/exec"
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

func CheckTorStatus() (bool, error) {
	cmd := exec.Command("tasklist", "/FI", "IMAGENAME eq tor.exe")
	output, err := cmd.Output()
	if err != nil {
		return false, err
	}

	outputStr := string(output)
	return strings.Contains(outputStr, "tor"), nil
}
