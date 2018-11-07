package handler

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

// GlobalFlags is global flags.
var GlobalFlags = []cli.Flag{}

// Commands is sub commands.
var Commands = []cli.Command{
	{
		Name:   "init",
		Usage:  "generate a configuration file voilerplate",
		Action: CmdInit,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "config, c",
				Value: ".gencfg.yml",
				Usage: "generated configuration file path",
			},
			cli.StringFlag{
				Name:  "template, t",
				Value: ".gencfg_config.tmpl",
				Usage: "generated template file path",
			},
		},
	},
	{
		Name:   "gen",
		Usage:  "generate code",
		Action: CmdGen,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "config, c",
				Value: ".gencfg.yml",
				Usage: "configuration file path",
			},
			cli.StringFlag{
				Name:  "dest, d",
				Usage: `generated file path (default: "config/config.go")`,
			},
			cli.StringFlag{
				Name:  "template, t",
				Usage: "template file path",
			},
		},
	},
	{
		Name:   "compare",
		Usage:  "compare existing code with one which will be generated by gen command",
		Action: CmdCompare,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "config, c",
				Value: ".gencfg.yml",
				Usage: "configuration file path",
			},
			cli.StringFlag{
				Name:  "dest, d",
				Usage: `generated file path (default: "config/config.go")`,
			},
			cli.StringFlag{
				Name:  "template, t",
				Usage: "template file path",
			},
			cli.BoolFlag{
				Name:  "failure, f",
				Usage: "exit with non 0 if diff exists",
			},
			cli.BoolFlag{
				Name:  "quiet, q",
				Usage: "diff is not shown. --failure option is enabled automatically",
			},
		},
	},
}

// CommandNotFound is called when the specified command is not found.
func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
