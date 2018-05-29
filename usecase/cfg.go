package usecase

import (
	"sort"

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

// Update updates cfg before rendering.
func (uc CfgUsecase) Update(cfg *domain.Cfg) {
	if cfg == nil {
		return
	}
	sort.Sort(domain.Params(cfg.Params))
}