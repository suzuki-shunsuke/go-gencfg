package config

import (
	"github.com/urfave/cli"
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
func GetCfgPath(c *cli.Context) string {
	return c.String(cfgPathKey)
}

// GetDestPath returns the created file path.
func GetDestPath(c *cli.Context) string {
	return c.String(destPathKey)
}

// GetTemplatePath returns the template file path.
func GetTemplatePath(c *cli.Context) string {
	return c.String(templateKey)
}

// GetTestTemplatePath returns the test template file path.
func GetTestTemplatePath(c *cli.Context) string {
	return c.String(testTemplateKey)
}

// GetParamNames returns the param names.
func GetParamNames(c *cli.Context) []string {
	return []string(c.Args())
}

// GetFailure returns failure option.
func GetFailure(c *cli.Context) bool {
	return c.Bool(failureKey)
}

// GetQuiet returns quiet option.
func GetQuiet(c *cli.Context) bool {
	return c.Bool(quietKey)
}
