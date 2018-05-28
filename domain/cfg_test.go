package domain_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-ptr"

	"github.com/suzuki-shunsuke/go-gencfg/domain"
)

func TestCfgHasFlag(t *testing.T) {
	cfg := domain.Cfg{}
	if cfg.HasFlag() {
		t.Fatal("cfg.HasFlag() = true, wanted false")
	}
	cfg.Params = []domain.Param{{}}
	if cfg.HasFlag() {
		t.Fatal("cfg.HasFlag() = true, wanted false")
	}
	cfg.Params = []domain.Param{{Flag: domain.Flag{Bind: ptr.PBool(true)}}}
	if !cfg.HasFlag() {
		t.Fatal("cfg.HasFlag() = false, wanted true")
	}
}

func TestCfgUpdate(t *testing.T) {
	cfg := domain.Cfg{}
	cfg.Update()
}
