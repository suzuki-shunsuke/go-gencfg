package domain

import (
	"fmt"
	"strings"
)

type (
	// Env represents environment variable.
	Env struct {
		Name   *string `yaml:"name"`
		Prefix *string `yaml:"prefix"`
		Bind   *bool   `yaml:"bind"`
	}
)

// IsBind returns whether the environment variable is bound.
func (e Env) IsBind(def *bool) bool {
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
func (e Env) GetPrefix(gPrefix *string) string {
	if e.Prefix != nil {
		return *e.Prefix
	}
	if gPrefix != nil {
		return *gPrefix
	}
	return ""
}

// GetName returns the environment variable name.
func (e Env) GetName(pName string, gPrefix *string) string {
	if e.Name != nil {
		return *e.Name
	}
	return strings.ToUpper(fmt.Sprintf("%s%s", e.GetPrefix(gPrefix), pName))
}
