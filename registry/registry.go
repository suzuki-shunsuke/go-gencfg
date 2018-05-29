package registry

import (
	"github.com/suzuki-shunsuke/go-gencfg/domain"
	"github.com/suzuki-shunsuke/go-gencfg/infra"
	"github.com/suzuki-shunsuke/go-gencfg/usecase"
)

// NewCfgReader returns a new domain.CfgReader.
func NewCfgReader() domain.CfgReader {
	return infra.YAMLCfgReader{}
}

// NewCodeGenerater returns a new domain.CodeGenerater.
func NewCodeGenerater() domain.CodeGenerater {
	return infra.CodeGenerater{
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
