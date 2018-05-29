package domain

import (
	"github.com/urfave/cli"
)

type (
	// FileExist checks whether a file exists.
	FileExist interface {
		Exist(src string) bool
	}

	// FileReader reads a file.
	FileReader interface {
		Read(src string) (string, error)
	}

	// FileWriter writes data to a file.
	FileWriter interface {
		Write(dest string, data []byte) error
	}

	// DirMaker makes a directory.
	DirMaker interface {
		Make(dest string) error
	}

	// TemplateRenderer renders a template with data.
	TemplateRenderer interface {
		Render(txt string, data interface{}) (string, error)
	}

	// CodeGenerater generates code at files with a template and data.
	CodeGenerater interface {
		Exec(dest, tmpl string, config TemplateData) error
	}

	// CfgReader represents the command configuration reader.
	CfgReader interface {
		Read(src string) (Cfg, error)
	}

	// CmdExecuter executes a command.
	CmdExecuter interface {
		Exec(cmd, dest string) error
	}

	// StrFormatter formats a string running commands.
	StrFormatter interface {
		Format(dest, text string, formatters []string, reader FileReader, executer CmdExecuter) (string, error)
	}

	// CLIContext represents a urfave.cli interface.
	CLIContext interface {
		String(string) string
		Bool(string) bool
		Args() cli.Args
	}
)
