package handler

import (
	"os"

	"github.com/urfave/cli"

	"github.com/suzuki-shunsuke/go-gencfg/internal/domain"
)

// Main is called cmd/gencfg/main.go .
func Main() {
	app := cli.NewApp()
	app.Name = "gencfg"
	app.Version = domain.Version
	app.Author = "suzuki-shunsuke https://github.com/suzuki-shunsuke"
	app.Usage = "generate code for the application specific configuration"

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	app.Run(os.Args)
}
