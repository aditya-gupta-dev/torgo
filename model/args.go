package model

type Args struct {
	Interval int64 `default:"60" help:"time in seconds between IP changes."`
	IP       bool  `help:"display the current IP and quit."`
}
