package usecase_test

import (
	"fmt"
	"testing"

	"github.com/suzuki-shunsuke/go-gencfg/internal/domain"
	"github.com/suzuki-shunsuke/go-gencfg/internal/test"
	"github.com/suzuki-shunsuke/go-gencfg/internal/usecase"
)

func TestGenCfgFile(t *testing.T) {
	args := domain.GenCfgFileArgs{
		Reader:    test.NoopFileReader{},
		CfgReader: test.NoopCfgReader{},
		Generater: test.NoopCodeGenerater{},
		Executer:  test.NoopCmdExecuter{},
		CfgUC:     usecase.CfgUsecase{},
		EnvUC:     usecase.EnvUsecase{},
		FlagUC:    usecase.FlagUsecase{},
		ParamUC:   usecase.ParamUsecase{},
	}
	if err := usecase.GenCfgFile(args); err != nil {
		t.Fatal(err)
	}
	args.CfgReader = test.NoopCfgReader{
		Err: fmt.Errorf(""),
	}
	if err := usecase.GenCfgFile(args); err == nil {
		t.Fatal("it should be failed to read configuration")
	}
	args.CfgReader = test.NoopCfgReader{}
	args.Dest = "hoge"
	if err := usecase.GenCfgFile(args); err == nil {
		t.Fatal("invalid dest")
	}
	args.Dest = ""
	args.CfgReader = test.NoopCfgReader{
		Cfg: domain.Cfg{
			Dest: "hoge.go",
		},
	}
	if err := usecase.GenCfgFile(args); err != nil {
		t.Fatal(err)
	}
	args.TmplPath = "foo.go"
	if err := usecase.GenCfgFile(args); err != nil {
		t.Fatal(err)
	}
}
