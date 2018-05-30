package usecase_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-gencfg/domain"
	"github.com/suzuki-shunsuke/go-gencfg/test"
	"github.com/suzuki-shunsuke/go-gencfg/usecase"
)

func TestGenCfgFile(t *testing.T) {
	args := domain.GenCfgFileArgs{
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
}
