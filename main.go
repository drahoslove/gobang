// go game in go lang

package main

import (
	"flag"

	"github.com/Drahoslav7/gobang/logic"
)

// GLOBALS
const (
	VERSION     = "0.1"
	APPNAME     = "Gobang!"
	DESCRIPTION = "Goban in Golang"
	AUTHOR      = "Drahoslav Bednář"
	WEBPAGE     = "go.yo2.cz"
)

var (
	app logic.Application
)

func main() {
	app = logic.NewApp()

	var gui, cli bool
	flag.BoolVar(&gui, "gui", true, "run in gui mode (default)")
	flag.BoolVar(&cli, "cli", false, "run in cli mode")
	// serve := flag.Bool("serve", false, "run as server, without gui")
	// port := flag.Int("port", 60109, "set port of server, only make sence with --server")
	flag.Parse()

	if !gui && !cli {
		gui = true
	}

	if gui {
		startGUI()
	}
	if cli {
		startCLI()
	}

	app.SetPlayer("Yoyo", nil)
	app.NewGame(16, nil)

}
