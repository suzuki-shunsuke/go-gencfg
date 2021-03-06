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

	// CodeGenerator generates code at files with a template and data.
	CodeGenerator interface {
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

	// CfgUsecase represents application logic about Cfg
	CfgUsecase interface {
		HasFlag(FlagUsecase, Cfg) bool
		GetPkg(Cfg) string
		GetPkgName(Cfg) string
		Update(*Cfg)
	}

	// EnvUsecase represents application logic about Env
	EnvUsecase interface {
		IsBind(e Env, def *bool) bool
		GetPrefix(e Env, gPrefix *string) string
		GetName(e Env, pName string, gPrefix *string) string
	}

	// FlagUsecase represents application logic about Env
	FlagUsecase interface {
		IsBind(f Flag, def *bool) bool
	}

	// ParamUsecase represents application logic about Env
	ParamUsecase interface {
		IsSetDefault(p Param) bool
		GetDefaultStr(p Param) string
		GetFlagName(p Param) string
		GetFlagDescription(p Param) string
		GetEnvName(p Param) string
		CamelCaseName(p Param) string
		CamelCaseLowerName(p Param) string
		GetType(p Param) string
		GetPFlagName(p Param) string
		GetViperGetterName(p Param) string
	}
)
