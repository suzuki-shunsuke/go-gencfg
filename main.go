package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "gencfg"
	app.Version = "0.1.4"
	app.Author = "suzuki-shunsuke https://github.com/suzuki-shunsuke"
	app.Usage = "generate code for the application specific configuration"

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	app.Run(os.Args)
}
