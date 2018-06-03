package test_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-gencfg/domain"
	"github.com/suzuki-shunsuke/go-gencfg/test"
)

func TestNoopFileExistExist(t *testing.T) {
	exist := test.NoopFileExist{}
	if exist.Exist("hoge") {
		t.Fatal("exist.Exist() = true, wanted false")
	}
}

func TestNoopFileReaderRead(t *testing.T) {
	reader := test.NoopFileReader{}
	if _, err := reader.Read("hoge"); err != nil {
		t.Fatal(err)
	}
}

func TestNoopFileWriterWrite(t *testing.T) {
	writer := test.NoopFileWriter{}
	if err := writer.Write("hoge", nil); err != nil {
		t.Fatal(err)
	}
}

func TestNoopDirMakerMake(t *testing.T) {
	maker := test.NoopDirMaker{}
	if err := maker.Make("hoge"); err != nil {
		t.Fatal(err)
	}
}

func TestErrDirMakerMake(t *testing.T) {
	maker := test.ErrDirMaker{}
	if err := maker.Make("hoge"); err == nil {
		t.Fatal("ErrDirMaker.Make() should return error")
	}
}

func TestNoopCfgReaderRead(t *testing.T) {
	reader := test.NoopCfgReader{}
	if _, err := reader.Read("hoge"); err != nil {
		t.Fatal(err)
	}
}

func TestNoopCmdExecuterExec(t *testing.T) {
	executer := test.NoopCmdExecuter{}
	if err := executer.Exec("hoge", "bar"); err != nil {
		t.Fatal(err)
	}
}

func TestNoopTemplateRendererRender(t *testing.T) {
	renderer := test.NoopTemplateRenderer{}
	if _, err := renderer.Render("hoge", nil); err != nil {
		t.Fatal(err)
	}
}

func TestNoopCodeGeneraterExec(t *testing.T) {
	cg := test.NoopCodeGenerater{}
	if err := cg.Exec("hoge", "bar", domain.TemplateData{}); err != nil {
		t.Fatal(err)
	}
}

func TestErrTemplateRendererRender(t *testing.T) {
	renderer := test.ErrTemplateRenderer{}
	if _, err := renderer.Render("hoge", nil); err == nil {
		t.Fatal("ErrTemplateRenderer.Render() should return error")
	}
}

func TestNoopStrFormatter(t *testing.T) {
	fmtr := test.NoopStrFormatter{}
	reader := test.NoopFileReader{}
	executer := test.NoopCmdExecuter{}
	if _, err := fmtr.Format("hoge", "bar", nil, reader, executer); err != nil {
		t.Fatal(err)
	}
}
