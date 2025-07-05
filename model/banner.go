package model

import "fmt"

func PrintTitle() {
	banner := `
` + Cyan + Bold + `
████████╗ ██████╗ ██████╗  ██████╗  ██████╗ 
╚══██╔══╝██╔═══██╗██╔══██╗██╔════╝ ██╔═══██╗
   ██║   ██║   ██║██████╔╝██║  ███╗██║   ██║
   ██║   ██║   ██║██╔══██╗██║   ██║██║   ██║
   ██║   ╚██████╔╝██║  ██║╚██████╔╝╚██████╔╝
   ╚═╝    ╚═════╝ ╚═╝  ╚═╝ ╚═════╝  ╚═════╝ 
` + Reset + `
` + BrightGreen + `         🔒 IP Changer using Tor. 🔒` + Reset + `
` + BrightCyan + `              Made with Go` + Reset + `
`
	fmt.Print(banner)
}
