package command

import (
	"github.com/urfave/cli"

	"github.com/suzuki-shunsuke/go-gencfg/config"
	"github.com/suzuki-shunsuke/go-gencfg/domain"
	"github.com/suzuki-shunsuke/go-gencfg/registry"
	"github.com/suzuki-shunsuke/go-gencfg/usecase"
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
			Generater: registry.NewCodeGenerater(),
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
