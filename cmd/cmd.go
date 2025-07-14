package cmd

import (
	"fmt"
	"os"

	"github.com/aditya-gupta-dev/torgo/torgo"
)

func PrintIPAndExit() {
	ip, err := torgo.ShowIP()
	if err != nil {
		fmt.Println("Failed to get IP :(")
	}
	fmt.Printf("\033[32m%s\033[0m", ip)
	os.Exit(0)
}

func StartIPChanging(interval, count int64) {
}
