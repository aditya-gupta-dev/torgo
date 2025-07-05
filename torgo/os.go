package torgo

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func CheckTorInstallation() (string, error) {
	switch runtime.GOOS {
	case "windows":
		return CheckTorForWindows()
	case "darwin":
		return CheckTorForMac()
	case "linux":
		return CheckTorForLinux()
	default:
		return "", fmt.Errorf("your operating system is not supported")
	}
}

// TODO:
func CheckTorForWindows() (string, error) {

	windowsPaths := []string{
		`C:\Program Files\Tor Browser\Browser\TorBrowser\Tor\tor.exe`,
		`C:\Program Files (x86)\Tor Browser\Browser\TorBrowser\Tor\tor.exe`,
		`C:\Users\%USERNAME%\Desktop\Tor Browser\Browser\TorBrowser\Tor\tor.exe`,
		`C:\Users\%USERNAME%\AppData\Local\Tor Browser\Browser\TorBrowser\Tor\tor.exe`,
		`C:\Tor\tor.exe`,
	}

	for _, path := range windowsPaths {
		if fileExists(path) {
			return path, nil
		}
	}

	return exec.LookPath("tor.exe")
}

func CheckTorForMac() (string, error) {
	macPaths := []string{
		"/Applications/Tor Browser.app/Contents/MacOS/Tor/tor",
		"/usr/local/bin/tor",
		"/opt/homebrew/bin/tor",
		"/usr/bin/tor",
		"/Applications/TorBrowser.app/Contents/Resources/TorBrowser/Tor/tor",
	}

	for _, path := range macPaths {
		if fileExists(path) {
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

func CheckTorForLinux() (string, error) {
	output, err := exec.Command("which", "tor").Output()
	return string(output), err
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
