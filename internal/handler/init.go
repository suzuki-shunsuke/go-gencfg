package handler

import (
	"github.com/suzuki-shunsuke/go-cliutil"
	"github.com/urfave/cli"

	"github.com/suzuki-shunsuke/go-gencfg/internal/config"
	"github.com/suzuki-shunsuke/go-gencfg/internal/domain"
	"github.com/suzuki-shunsuke/go-gencfg/internal/registry"
	"github.com/suzuki-shunsuke/go-gencfg/internal/usecase"
)

// CmdInit is the sub command "init".
func CmdInit(c *cli.Context) error {
	return cliutil.ConvErrToExitError(
		usecase.InitGenCfgFile(
			domain.InitGenCfgFileArgs{
				Dest:     config.GetCfgPath(c),
				TmplDest: config.GetTemplatePath(c),
				PNames:   config.GetParamNames(c),
				Renderer: registry.NewTemplateRenderer(),
				Writer:   registry.NewFileWriter(),
				Exist:    registry.NewFileExist(),
			}))
}
