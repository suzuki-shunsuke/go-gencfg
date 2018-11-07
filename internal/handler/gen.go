package handler

import (
	"github.com/urfave/cli"

	"github.com/suzuki-shunsuke/go-gencfg/internal/config"
	"github.com/suzuki-shunsuke/go-gencfg/internal/domain"
	"github.com/suzuki-shunsuke/go-gencfg/internal/registry"
	"github.com/suzuki-shunsuke/go-gencfg/internal/usecase"
)

// CmdGen is sub command "gen".
func CmdGen(c *cli.Context) error {
	err := usecase.GenCfgFile(
		domain.GenCfgFileArgs{
			Src:       config.GetCfgPath(c),
			Dest:      config.GetDestPath(c),
			TmplPath:  config.GetTemplatePath(c),
			Reader:    registry.NewFileReader(),
			CfgReader: registry.NewCfgReader(),
			Generator: registry.NewCodeGenerator(),
			Executer:  registry.NewCmdExecuter(),
			CfgUC:     registry.NewCfgUsecase(),
			EnvUC:     registry.NewEnvUsecase(),
			FlagUC:    registry.NewFlagUsecase(),
			ParamUC:   registry.NewParamUsecase(),
		})
	if err == nil {
		return nil
	}
	return cli.NewExitError(err.Error(), 1)
}
