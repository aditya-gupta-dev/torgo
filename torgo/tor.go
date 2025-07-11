package torgo

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

func IsTorRunning() (bool, error) {

	switch runtime.GOOS {
	case "windows":
		return CheckTorStatusForWindows()
	case "darwin":
		return CheckTorStatusForMac()
	case "linux":
		return CheckTorStatusForLinux()
	default:
		return false, fmt.Errorf("unsupported operating system found while checking tor process status")
	}
}

func CheckTorStatusForLinux() (bool, error) {
	cmd := exec.Command("systemctl", "is-active", "tor")
	output, err := cmd.Output()
	if err != nil {
		return false, err
	}

	status := strings.TrimSpace(string(output))

	return status == "active", nil
}

func CheckTorStatusForWindows() (bool, error) {
	cmd := exec.Command("tasklist", "/FI", "IMAGENAME eq tor.exe")
	output, err := cmd.Output()
	if err != nil {
		return false, err
	}

	outputStr := string(output)
	return strings.Contains(outputStr, "tor"), nil
}

func CheckTorStatusForMac() (bool, error) {
	cmd := exec.Command("pgrep", "tor")
	output, err := cmd.Output()
	if err != nil {
		return false, err
	}
	outputStr := string(output)
	return strings.TrimSpace(outputStr) != "", nil
}
