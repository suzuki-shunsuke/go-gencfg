package usecase

import (
	"github.com/suzuki-shunsuke/go-gencfg/internal/domain"
)

type (
	// FlagUsecase represents flag.
	FlagUsecase struct{}
)

// IsBind returns whether the flag is bound.
func (fuc FlagUsecase) IsBind(f domain.Flag, def *bool) bool {
	b := f.Bind
	if b == nil {
		if f.Name != nil || f.Description != nil || f.Short != nil {
			return true
		}
		if def == nil {
			return false
		}
		return *def
	}
	return *b
}
