package mac

import (
	"fmt"
	"os/exec"

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
