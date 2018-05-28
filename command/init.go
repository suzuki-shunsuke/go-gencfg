package command

import (
	"github.com/urfave/cli"

	"github.com/suzuki-shunsuke/go-gencfg/config"
	"github.com/suzuki-shunsuke/go-gencfg/registry"
	"github.com/suzuki-shunsuke/go-gencfg/usecase"
)

// CmdInit is the sub command "init".
func CmdInit(c *cli.Context) error {
	renderer := registry.NewTemplateRenderer()
	exist := registry.NewFileExist()
	writer := registry.NewFileWriter()
	err := usecase.InitGenCfgFile(
		config.GetDestPath(c), config.GetParamNames(c), renderer, writer, exist)
	if err == nil {
		return nil
	}
	return cli.NewExitError(err.Error(), 1)
}
