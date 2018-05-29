package usecase

import (
	"path/filepath"
	"sort"
	"strings"

	"github.com/suzuki-shunsuke/go-gencfg/domain"
)

type (
	// CfgUsecase implements domain.CfgUsecase interface.
	CfgUsecase struct{}
)

// HasFlag returns whether there is a bound flag.
func (uc CfgUsecase) HasFlag(flagUC domain.FlagUsecase, cfg domain.Cfg) bool {
	for _, p := range cfg.Params {
		if flagUC.IsBind(p.Flag, cfg.Global.Flag.Bind) {
			return true
		}
	}
	return false
}

// GetPkgName returns the package name.
func (uc CfgUsecase) GetPkgName(cfg domain.Cfg) string {
	if cfg.PkgName != "" {
		return cfg.PkgName
	}
	if cfg.Pkg != "" {
		a := strings.Split(filepath.Base(cfg.Pkg), "-")
		return strings.Split(a[len(a)-1], ".")[0]
	}
	return "config"
}

// GetPkg returns the package path.
func (uc CfgUsecase) GetPkg(cfg domain.Cfg) string {
	return cfg.Pkg
}

// Update updates cfg before rendering.
func (uc CfgUsecase) Update(cfg *domain.Cfg) {
	if cfg == nil {
		return
	}
	sort.Sort(domain.Params(cfg.Params))
}
