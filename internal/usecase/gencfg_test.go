package usecase_test

import (
	"fmt"
	"testing"

	"github.com/suzuki-shunsuke/gomic/gomic"

	"github.com/suzuki-shunsuke/go-gencfg/internal/domain"
	"github.com/suzuki-shunsuke/go-gencfg/internal/test"
	"github.com/suzuki-shunsuke/go-gencfg/internal/usecase"
)

func TestGenCfgFile(t *testing.T) {
	args := domain.GenCfgFileArgs{
		Reader:    test.NewFileReaderMock(t, gomic.DoNothing),
		CfgReader: test.NewCfgReaderMock(t, gomic.DoNothing),
		Generator: test.NewCodeGeneratorMock(t, gomic.DoNothing),
		Executer:  test.NewCmdExecuterMock(t, gomic.DoNothing),
		CfgUC:     usecase.CfgUsecase{},
		EnvUC:     usecase.EnvUsecase{},
		FlagUC:    usecase.FlagUsecase{},
		ParamUC:   usecase.ParamUsecase{},
	}
	if err := usecase.GenCfgFile(args); err != nil {
		t.Fatal(err)
	}
	cfgReader := test.NewCfgReaderMock(t, gomic.DoNothing)
	cfgReader.SetReturnRead(domain.Cfg{}, fmt.Errorf("failed to read the configuration file"))
	args.CfgReader = cfgReader
	if err := usecase.GenCfgFile(args); err == nil {
		t.Fatal("it should be failed to read configuration")
	}
	cfgReader.SetFuncRead(nil)
	args.Dest = "hoge"
	if err := usecase.GenCfgFile(args); err == nil {
		t.Fatal("invalid dest")
	}
	args.Dest = ""
	cfgReader.SetReturnRead(domain.Cfg{Dest: "hoge.go"}, nil)
	if err := usecase.GenCfgFile(args); err != nil {
		t.Fatal(err)
	}
	args.TmplPath = "foo.go"
	if err := usecase.GenCfgFile(args); err != nil {
		t.Fatal(err)
	}
}
