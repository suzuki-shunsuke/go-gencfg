package usecase_test

import (
	"testing"

	"github.com/suzuki-shunsuke/gomic/gomic"

	"github.com/suzuki-shunsuke/go-gencfg/internal/domain"
	"github.com/suzuki-shunsuke/go-gencfg/internal/test"
	"github.com/suzuki-shunsuke/go-gencfg/internal/usecase"
)

func TestInitGenCfgFile(t *testing.T) {
	args := domain.InitGenCfgFileArgs{
		Exist:    test.NewFileExistMock(t, gomic.DoNothing),
		Renderer: test.NewTemplateRendererMock(t, gomic.DoNothing),
		Writer:   test.NewFileWriterMock(t, gomic.DoNothing),
	}
	if err := usecase.InitGenCfgFile(args); err != nil {
		t.Fatal(err)
	}
}
