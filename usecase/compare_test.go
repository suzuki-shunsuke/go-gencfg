package usecase_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-gencfg/domain"
	"github.com/suzuki-shunsuke/go-gencfg/test"
	"github.com/suzuki-shunsuke/go-gencfg/usecase"
)

func TestCompare(t *testing.T) {
	args := domain.CompareArgs{
		Reader:       test.NoopFileReader{},
		Renderer:     test.NoopTemplateRenderer{},
		CfgReader:    test.NoopCfgReader{},
		Executer:     test.NoopCmdExecuter{},
		StrFormatter: test.NoopStrFormatter{},
		CfgUC:        usecase.CfgUsecase{},
		EnvUC:        usecase.EnvUsecase{},
		FlagUC:       usecase.FlagUsecase{},
		ParamUC:      usecase.ParamUsecase{},
	}
	if err := usecase.Compare(args); err != nil {
		t.Fatal(err)
	}
}
