package usecase

// DefaultCfgTestTmpl is the default template of the generated wrapper test code.
const DefaultCfgTestTmpl = `
package cfg_test

import (
	"testing"
	// import your cfg package
)

{{- range .Params}}
func Test{{.CamelCaseName}}(t *testing.T) {
	// write your test
	cfg.{{.CamelCaseName}}()
}

func TestSet{{.CamelCaseName}}(t *testing.T) {
	// write your test
	cfg.Set{{.CamelCaseName}}({{.GetDefaultStr}})
}
{{- end}}
`
