package infra_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-gencfg/domain"
	"github.com/suzuki-shunsuke/go-gencfg/infra"
)

func TestCodeGeneraterExec(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		cg := infra.CodeGenerater{
			Renderer:   NoopRenderer{},
			DirMaker:   NoopDirMaker{},
			FileWriter: NoopFileWriter{},
		}
		if err := cg.Exec("/tmp/hello.go", "", domain.Cfg{}); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("failed to mkdir", func(t *testing.T) {
		cg := infra.CodeGenerater{
			Renderer:   NoopRenderer{},
			DirMaker:   ErrDirMaker{},
			FileWriter: NoopFileWriter{},
		}
		if err := cg.Exec("/tmp/hello.go", "", domain.Cfg{}); err == nil {
			t.Fatal("it should be failed to make directory")
		}
	})
	t.Run("failed to renderer", func(t *testing.T) {
		cg := infra.CodeGenerater{
			Renderer:   ErrRenderer{},
			DirMaker:   NoopDirMaker{},
			FileWriter: NoopFileWriter{},
		}
		if err := cg.Exec("/tmp/hello.go", "", domain.Cfg{}); err == nil {
			t.Fatal("it should be failed to render template")
		}
	})
}
