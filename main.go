package main

import (
	"github.com/aditya-gupta-dev/torgo/cmd"
	"github.com/aditya-gupta-dev/torgo/model"
	"github.com/alexflint/go-arg"
)

func main() {
	var args model.Args
	arg.MustParse(&args)

	if args.IP {
		cmd.PrintIPAndExit()
	}

	if args.ChangeIP {
		cmd.StartIPChanging(args.Interval, args.Count)
	} else {
		cmd.StartIPChanging(args.Interval, args.Count)
	}
}
