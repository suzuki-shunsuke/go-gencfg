package command

import (
	"github.com/urfave/cli"

	"github.com/suzuki-shunsuke/go-gencfg/config"
	"github.com/suzuki-shunsuke/go-gencfg/registry"
	"github.com/suzuki-shunsuke/go-gencfg/usecase"
)

// CmdGen is sub command "gen".
func CmdGen(c *cli.Context) error {
	generater := registry.NewCodeGenerater()
	cfgReader := registry.NewCfgReader()
	executer := registry.NewCmdExecuter()
	reader := registry.NewFileReader()
	err := usecase.GenCfgFile(
		config.GetCfgPath(c), config.GetDestPath(c),
		config.GetTemplatePath(c), config.GetTestTemplatePath(c),
		reader, cfgReader, generater, executer)
	if err == nil {
		return nil
	}
	return cli.NewExitError(err.Error(), 1)
}
