package usecase_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-ptr"

	"github.com/suzuki-shunsuke/go-gencfg/internal/domain"
	"github.com/suzuki-shunsuke/go-gencfg/internal/usecase"
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

func TestCfgUsecaseGetPkgName(t *testing.T) {
	cfg := domain.Cfg{}
	cfgUC := usecase.CfgUsecase{}
	exp := "config"
	act := cfgUC.GetPkgName(cfg)
	if act != exp {
		t.Fatalf(`cfgUC.GetPkgName(cfg) = "%s", wanted "%s"`, act, exp)
	}
	exp = "hoge"
	cfg.PkgName = exp
	act = cfgUC.GetPkgName(cfg)
	if act != exp {
		t.Fatalf(`cfgUC.GetPkgName(cfg) = "%s", wanted "%s"`, act, exp)
	}
	cfg.PkgName = ""
	cfg.Pkg = "github.com/suzuki-shunsuke/go-hoge.v2"
	act = cfgUC.GetPkgName(cfg)
	if act != exp {
		t.Fatalf(`cfgUC.GetPkgName(cfg) = "%s", wanted "%s"`, act, exp)
	}
}

func TestCfgUsecaseGetPkg(t *testing.T) {
	cfg := domain.Cfg{}
	cfgUC := usecase.CfgUsecase{}
	exp := ""
	act := cfgUC.GetPkg(cfg)
	if act != exp {
		t.Fatalf(`cfgUC.GetPkg(cfg) = "%s", wanted "%s"`, act, exp)
	}
}

func TestCfgUsecaseUpdate(t *testing.T) {
	cfg := domain.Cfg{}
	cfgUC := usecase.CfgUsecase{}
	cfgUC.Update(&cfg)
	cfgUC.Update(nil)
}
