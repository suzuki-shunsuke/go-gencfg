package usecase

import (
	"path/filepath"
	"sort"
	"strings"

	"github.com/suzuki-shunsuke/go-gencfg/internal/domain"
)

type (
	// CfgUsecase implements domain.CfgUsecase interface.
	CfgUsecase struct{}
)

// HasFlag returns whether there is a bound flag.
func (uc CfgUsecase) HasFlag(flagUC domain.FlagUsecase, cfg domain.Cfg) bool {
	for _, p := range cfg.Params {
		if flagUC.IsBind(p.Flag, cfg.Default.Flag.Bind) {
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
	if cfg.CfgFileParam.Name != "" {
		if cfg.CfgFileParam.Type == "" {
			cfg.CfgFileParam.Type = "string"
		}
		if cfg.Params == nil {
			cfg.Params = []domain.Param{}
		}
		cfg.Params = append(cfg.Params, cfg.CfgFileParam)
	}
	sort.Sort(domain.Params(cfg.Params))
}
