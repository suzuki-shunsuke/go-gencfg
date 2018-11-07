package usecase_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-gencfg/internal/domain"
	"github.com/suzuki-shunsuke/go-gencfg/internal/test"
	"github.com/suzuki-shunsuke/go-gencfg/internal/usecase"
)

func TestInitGenCfgFile(t *testing.T) {
	args := domain.InitGenCfgFileArgs{
		Exist:    test.NoopFileExist{},
		Renderer: test.NoopTemplateRenderer{},
		Writer:   test.NoopFileWriter{},
	}
	if err := usecase.InitGenCfgFile(args); err != nil {
		t.Fatal(err)
	}
}
