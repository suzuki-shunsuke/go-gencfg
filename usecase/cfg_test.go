package usecase_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-ptr"

	"github.com/suzuki-shunsuke/go-gencfg/domain"
	"github.com/suzuki-shunsuke/go-gencfg/usecase"
)

func TestCfgUsecaseHasFlag(t *testing.T) {
	cfg := domain.Cfg{}
	cfgUC := usecase.CfgUsecase{}
	flagUC := usecase.FlagUsecase{}
	if cfgUC.HasFlag(flagUC, cfg) {
		t.Fatal("cfg.HasFlag() = true, wanted false")
	}
	cfg.Params = []domain.Param{{}}
	if cfgUC.HasFlag(flagUC, cfg) {
		t.Fatal("cfg.HasFlag() = true, wanted false")
	}
	cfg.Params = []domain.Param{{Flag: domain.Flag{Bind: ptr.PBool(true)}}}
	if !cfgUC.HasFlag(flagUC, cfg) {
		t.Fatal("cfg.HasFlag() = false, wanted true")
	}
}

func TestCfgUsecaseUpdate(t *testing.T) {
	cfg := domain.Cfg{}
	cfgUC := usecase.CfgUsecase{}
	cfgUC.Update(&cfg)
}
