package torgo

import (
	"fmt"
	"runtime"

	"github.com/aditya-gupta-dev/torgo/linux"
	"github.com/aditya-gupta-dev/torgo/mac"
	"github.com/aditya-gupta-dev/torgo/windows"
)

func IsTorInstalled() (string, error) {
	switch runtime.GOOS {
	case "windows":
		return windows.CheckTorInstallation()
	case "darwin":
		return mac.CheckTorInstallation()
	case "linux":
		return linux.CheckTorInstallation()
	default:
		return "", fmt.Errorf("your operating system is not supported yet")
	}
}

func IsTorRunning(torPath string) (bool, error) {

	switch runtime.GOOS {
	case "windows":
		return windows.CheckTorStatus(torPath)
	case "darwin":
		return mac.CheckTorStatus(torPath)
	case "linux":
		return linux.CheckTorStatus(torPath)
	default:
		return false, fmt.Errorf("unsupported OS found while checking tor process status")
	}
}

func RestartTorService(torPath string) error {
	switch runtime.GOOS {
	case "windows":
		return windows.RestartTor(torPath)
	case "darwin":
		return mac.RestartTor(torPath)
	case "linux":
		return linux.RestartTor(torPath)
	default:
		return fmt.Errorf("unsupported OS found while restarting tor")
	}
}
