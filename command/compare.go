package command

import (
	"github.com/urfave/cli"

	"github.com/suzuki-shunsuke/go-gencfg/config"
	"github.com/suzuki-shunsuke/go-gencfg/registry"
	"github.com/suzuki-shunsuke/go-gencfg/usecase"
)

// CmdCompare is sub command "compare".
func CmdCompare(c *cli.Context) error {
	cfgReader := registry.NewCfgReader()
	reader := registry.NewFileReader()
	renderer := registry.NewTemplateRenderer()
	executer := registry.NewCmdExecuter()
	strFormatter := registry.NewStrFormatter()
	err := usecase.Compare(
		config.GetCfgPath(c), config.GetDestPath(c),
		config.GetTemplatePath(c), config.GetTestTemplatePath(c),
		config.GetFailure(c), config.GetQuiet(c),
		cfgReader, renderer, reader, executer, strFormatter)
	if err == nil {
		return nil
	}
	return cli.NewExitError(err.Error(), 1)
}
