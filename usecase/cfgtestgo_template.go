package usecase

// DefaultCfgTestTmpl is the default template of the generated wrapper test code.
const DefaultCfgTestTmpl = `
{{ $paramUC := .ParamUC -}}
package {{.CfgUC.GetPkgName .Cfg}}

import (
	"testing"
)

{{- range .Cfg.Params}}
func Test{{$paramUC.CamelCaseName .}}(t *testing.T) {
	// write your test
	Get{{$paramUC.CamelCaseName .}}()
}

func TestSet{{$paramUC.CamelCaseName .}}(t *testing.T) {
	// write your test
	Set{{$paramUC.CamelCaseName .}}({{ $paramUC.GetDefaultStr .}})
}
{{- end}}
`
