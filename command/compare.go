package command

import (
	"github.com/urfave/cli"

	"github.com/suzuki-shunsuke/go-gencfg/internal/config"
	"github.com/suzuki-shunsuke/go-gencfg/internal/domain"
	"github.com/suzuki-shunsuke/go-gencfg/internal/registry"
	"github.com/suzuki-shunsuke/go-gencfg/internal/usecase"
)

// CmdCompare is sub command "compare".
func CmdCompare(c *cli.Context) error {
	err := usecase.Compare(
		domain.CompareArgs{
			Src:          config.GetCfgPath(c),
			Dest:         config.GetDestPath(c),
			TmplPath:     config.GetTemplatePath(c),
			IsQuiet:      config.GetQuiet(c),
			IsFailure:    config.GetFailure(c),
			Reader:       registry.NewFileReader(),
			Renderer:     registry.NewTemplateRenderer(),
			CfgReader:    registry.NewCfgReader(),
			Executer:     registry.NewCmdExecuter(),
			StrFormatter: registry.NewStrFormatter(),
			CfgUC:        registry.NewCfgUsecase(),
			EnvUC:        registry.NewEnvUsecase(),
			FlagUC:       registry.NewFlagUsecase(),
			ParamUC:      registry.NewParamUsecase(),
		})
	if err == nil {
		return nil
	}
	return cli.NewExitError(err.Error(), 1)
}
