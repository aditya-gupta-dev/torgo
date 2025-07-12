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

func IsTorRunning() (bool, error) {

	switch runtime.GOOS {
	case "windows":
		return windows.CheckTorStatus()
	case "darwin":
		return mac.CheckTorStatus()
	case "linux":
		return linux.CheckTorStatus()
	default:
		return false, fmt.Errorf("unsupported operating system found while checking tor process status")
	}
}
