package usecase

// DefaultCfgTestTmpl is the default template of the generated wrapper test code.
const DefaultCfgTestTmpl = `
package config_test

import (
	"testing"
	// import your cfg package
)

{{- range .Cfg.Params}}
func Test{{.CamelCaseName}}(t *testing.T) {
	// write your test
	config.Get{{.CamelCaseName}}()
}

func TestSet{{.CamelCaseName}}(t *testing.T) {
	// write your test
	config.Set{{.CamelCaseName}}({{.GetDefaultStr}})
}
{{- end}}
`
