package usecase

// DefaultCfgTestTmpl is the default template of the generated wrapper test code.
const DefaultCfgTestTmpl = `
{{ $paramUC := .ParamUC -}}
package config_test

import (
	"testing"
	// import your cfg package
)

{{- range .Cfg.Params}}
func Test{{$paramUC.CamelCaseName .}}(t *testing.T) {
	// write your test
	config.Get{{$paramUC.CamelCaseName .}}()
}

func TestSet{{$paramUC.CamelCaseName .}}(t *testing.T) {
	// write your test
	config.Set{{$paramUC.CamelCaseName .}}({{ $paramUC.GetDefaultStr .}})
}
{{- end}}
`
