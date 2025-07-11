package model

type Args struct {
	Interval int64 `arg:"-i,--interval" default:"30" help:"time in seconds between IP changes. defaults to 10"`
	Count    int64 `arg:"-c,--count" default:"0" help:"number of times you want to change IP. defaults to infinity"`
	IP       bool  `help:"display the current IP and quit."`
}
