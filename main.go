package main

import (
	"fmt"
	"os"

	"github.com/aditya-gupta-dev/torgo/model"
	"github.com/aditya-gupta-dev/torgo/torgo"
	"github.com/alexflint/go-arg"
)

func main() {
	var args model.Args
	arg.MustParse(&args)

	if args.IP {
		ip, err := torgo.ShowIP()

		if err != nil {
			fmt.Println("Failed to get IP :(")
		}

		fmt.Printf("\033[32m%s\033[0m", ip)
		os.Exit(0)
	}

	if args.ChangeIP {
		torgo.ChangeIPRepeatedly(args.Interval, args.Count)
	} else {
		torgo.ChangeIPRepeatedly(10, 0)
	}
}
