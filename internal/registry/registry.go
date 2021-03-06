package registry

import (
	"github.com/suzuki-shunsuke/go-gencfg/internal/domain"
	"github.com/suzuki-shunsuke/go-gencfg/internal/infra"
	"github.com/suzuki-shunsuke/go-gencfg/internal/usecase"
)

// NewCfgReader returns a new domain.CfgReader.
func NewCfgReader() domain.CfgReader {
	return infra.YAMLCfgReader{}
}

// NewCodeGenerator returns a new domain.CodeGenerator.
func NewCodeGenerator() domain.CodeGenerator {
	return infra.CodeGenerator{
		Renderer:   NewTemplateRenderer(),
		DirMaker:   infra.DirMaker{},
		FileWriter: NewFileWriter(),
	}
}

// NewTemplateRenderer returns a new domain.TemplateRenderer.
func NewTemplateRenderer() domain.TemplateRenderer {
	return infra.TextTemplateRenderer{}
}

// NewFileReader returns a new domain.FileReader.
func NewFileReader() domain.FileReader {
	return infra.FileReader{}
}

// NewFileExist returns a new domain.FileExist.
func NewFileExist() domain.FileExist {
	return infra.FileExist{}
}

// NewFileWriter returns a new domain.FileWriter.
func NewFileWriter() domain.FileWriter {
	return infra.FileWriter{}
}

// NewCmdExecuter returns a new domain.CmdExecuter.
func NewCmdExecuter() domain.CmdExecuter {
	return infra.OSExecExecuter{}
}

// NewStrFormatter returns a new domain.StrFormatter.
func NewStrFormatter() domain.StrFormatter {
	return infra.StrFormatter{}
}

// NewCfgUsecase returns a new domain.CfgUsecase.
func NewCfgUsecase() domain.CfgUsecase {
	return usecase.CfgUsecase{}
}

// NewEnvUsecase returns a new domain.EnvUsecase.
func NewEnvUsecase() domain.EnvUsecase {
	return usecase.EnvUsecase{}
}

// NewFlagUsecase returns a new domain.FlagUsecase.
func NewFlagUsecase() domain.FlagUsecase {
	return usecase.FlagUsecase{}
}

// NewParamUsecase returns a new domain.ParamUsecase.
func NewParamUsecase() domain.ParamUsecase {
	return usecase.ParamUsecase{}
}
