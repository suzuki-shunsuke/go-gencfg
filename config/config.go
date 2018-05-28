package config

import (
	"github.com/suzuki-shunsuke/go-gencfg/domain"
)

const (
	cfgPathKey      = "config"
	destPathKey     = "dest"
	failureKey      = "failure"
	quietKey        = "quiet"
	templateKey     = "template"
	testTemplateKey = "test-template"
)

// GetCfgPath returns the configuration file path.
func GetCfgPath(c domain.CLIContext) string {
	return c.String(cfgPathKey)
}

// GetDestPath returns the created file path.
func GetDestPath(c domain.CLIContext) string {
	return c.String(destPathKey)
}

// GetTemplatePath returns the template file path.
func GetTemplatePath(c domain.CLIContext) string {
	return c.String(templateKey)
}

// GetTestTemplatePath returns the test template file path.
func GetTestTemplatePath(c domain.CLIContext) string {
	return c.String(testTemplateKey)
}

// GetParamNames returns the param names.
func GetParamNames(c domain.CLIContext) []string {
	return []string(c.Args())
}

// GetFailure returns failure option.
func GetFailure(c domain.CLIContext) bool {
	return c.Bool(failureKey)
}

// GetQuiet returns quiet option.
func GetQuiet(c domain.CLIContext) bool {
	return c.Bool(quietKey)
}
