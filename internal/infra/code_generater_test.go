package infra_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-gencfg/internal/domain"
	"github.com/suzuki-shunsuke/go-gencfg/internal/infra"
	"github.com/suzuki-shunsuke/go-gencfg/internal/test"
)

func TestCodeGeneraterExec(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		cg := infra.CodeGenerater{
			Renderer:   test.NoopTemplateRenderer{},
			DirMaker:   test.NoopDirMaker{},
			FileWriter: test.NoopFileWriter{},
		}
		if err := cg.Exec("/tmp/hello.go", "", domain.TemplateData{}); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("failed to mkdir", func(t *testing.T) {
		cg := infra.CodeGenerater{
			Renderer:   test.NoopTemplateRenderer{},
			DirMaker:   test.ErrDirMaker{},
			FileWriter: test.NoopFileWriter{},
		}
		if err := cg.Exec("/tmp/hello.go", "", domain.TemplateData{}); err == nil {
			t.Fatal("it should be failed to make directory")
		}
	})
	t.Run("failed to renderer", func(t *testing.T) {
		cg := infra.CodeGenerater{
			Renderer:   test.ErrTemplateRenderer{},
			DirMaker:   test.NoopDirMaker{},
			FileWriter: test.NoopFileWriter{},
		}
		if err := cg.Exec("/tmp/hello.go", "", domain.TemplateData{}); err == nil {
			t.Fatal("it should be failed to render template")
		}
	})
}
