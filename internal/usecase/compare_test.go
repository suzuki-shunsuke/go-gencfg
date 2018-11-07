package usecase_test

import (
	"testing"

	"github.com/suzuki-shunsuke/gomic/gomic"

	"github.com/suzuki-shunsuke/go-gencfg/internal/domain"
	"github.com/suzuki-shunsuke/go-gencfg/internal/test"
	"github.com/suzuki-shunsuke/go-gencfg/internal/usecase"
)

func TestCompare(t *testing.T) {
	args := domain.CompareArgs{
		Reader:       test.NewFileReaderMock(t, gomic.DoNothing),
		Renderer:     test.NewTemplateRendererMock(t, gomic.DoNothing),
		CfgReader:    test.NewCfgReaderMock(t, gomic.DoNothing),
		Executer:     test.NewCmdExecuterMock(t, gomic.DoNothing),
		StrFormatter: test.NewStrFormatterMock(t, gomic.DoNothing),
		CfgUC:        usecase.CfgUsecase{},
		EnvUC:        usecase.EnvUsecase{},
		FlagUC:       usecase.FlagUsecase{},
		ParamUC:      usecase.ParamUsecase{},
	}
	if err := usecase.Compare(args); err != nil {
		t.Fatal(err)
	}
}
