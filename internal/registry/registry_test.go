package registry_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-gencfg/internal/registry"
)

func TestNewCfgReader(t *testing.T) {
	registry.NewCfgReader()
}

func TestNewCodeGenerater(t *testing.T) {
	registry.NewCodeGenerater()
}

func TestNewTemplateRenderer(t *testing.T) {
	registry.NewTemplateRenderer()
}

func TestNewFileReader(t *testing.T) {
	registry.NewFileReader()
}

func TestNewFileExist(t *testing.T) {
	registry.NewFileExist()
}

func TestNewFileWriter(t *testing.T) {
	registry.NewFileWriter()
}

func TestNewCmdExecuter(t *testing.T) {
	registry.NewCmdExecuter()
}

func TestNewStrFormatter(t *testing.T) {
	registry.NewStrFormatter()
}

func TestNewCfgUsecase(t *testing.T) {
	registry.NewCfgUsecase()
}

func TestNewEnvUsecase(t *testing.T) {
	registry.NewEnvUsecase()
}

func TestNewFlagUsecase(t *testing.T) {
	registry.NewFlagUsecase()
}

func TestNewParamUsecase(t *testing.T) {
	registry.NewParamUsecase()
}
