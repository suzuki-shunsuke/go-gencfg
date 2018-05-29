package usecase

import (
	"fmt"
	"strings"

	"github.com/suzuki-shunsuke/go-gencfg/domain"
)

type (
	// EnvUsecase represents environment variable.
	EnvUsecase struct{}
)

// IsBind returns whether the environment variable is bound.
func (envUC EnvUsecase) IsBind(e domain.Env, def *bool) bool {
	b := e.Bind
	if b == nil {
		if e.Prefix != nil {
			return true
		}
		if def == nil {
			return false
		}
		return *def
	}
	return *b
}

// GetPrefix returns the environment variable name's prefix.
func (envUC EnvUsecase) GetPrefix(e domain.Env, gPrefix *string) string {
	if e.Prefix != nil {
		return *e.Prefix
	}
	if gPrefix != nil {
		return *gPrefix
	}
	return ""
}

// GetName returns the environment variable name.
func (envUC EnvUsecase) GetName(e domain.Env, pName string, gPrefix *string) string {
	if e.Name != nil {
		return *e.Name
	}
	return strings.ToUpper(fmt.Sprintf("%s%s", envUC.GetPrefix(e, gPrefix), pName))
}
