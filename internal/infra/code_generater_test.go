package infra_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suzuki-shunsuke/gomic/gomic"

	"github.com/suzuki-shunsuke/go-gencfg/internal/domain"
	"github.com/suzuki-shunsuke/go-gencfg/internal/infra"
	"github.com/suzuki-shunsuke/go-gencfg/internal/test"
)

func TestCodeGeneratorExec(t *testing.T) {
	failedToMkdirMock := test.NewDirMakerMock(t, gomic.DoNothing)
	failedToMkdirMock.SetFakeMake(fmt.Errorf("failed to mkdir"))
	failedToRendererMock := test.NewTemplateRendererMock(t, gomic.DoNothing)
	failedToRendererMock.SetFakeRender("", fmt.Errorf("failed to renderer"))
	data := []struct {
		message    string
		renderer   domain.TemplateRenderer
		dirMaker   domain.DirMaker
		fileWriter domain.FileWriter
		checkErr   func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			"positive",
			test.NewTemplateRendererMock(t, gomic.DoNothing),
			test.NewDirMakerMock(t, gomic.DoNothing),
			test.NewFileWriterMock(t, gomic.DoNothing),
			assert.Nil,
		},
		{
			"failed to mkdir",
			test.NewTemplateRendererMock(t, gomic.DoNothing),
			failedToMkdirMock,
			test.NewFileWriterMock(t, gomic.DoNothing),
			assert.NotNil,
		},
		{
			"failed to renderer",
			failedToRendererMock,
			test.NewDirMakerMock(t, gomic.DoNothing),
			test.NewFileWriterMock(t, gomic.DoNothing),
			assert.NotNil,
		},
	}

	for _, d := range data {
		t.Run(d.message, func(t *testing.T) {
			cg := infra.CodeGenerator{
				Renderer:   d.renderer,
				DirMaker:   d.dirMaker,
				FileWriter: d.fileWriter,
			}
			d.checkErr(t, cg.Exec("/tmp/hello.go", "", domain.TemplateData{}))
		})
	}
}
