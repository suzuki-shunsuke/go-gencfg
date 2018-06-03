package usecase

// DefaultCfgTestTmpl is the default template of the generated wrapper test code.
const DefaultCfgTestTmpl = `
{{ $paramUC := .ParamUC -}}
{{ $pkgName := .CfgUC.GetPkgName .Cfg -}}
package {{$pkgName}}_test

import (
	"testing"
	{{- if .CfgUC.GetPkg .Cfg}}

	"{{.CfgUC.GetPkg .Cfg}}"
	{{- else}}
	// import your cfg package
	{{ end -}}
)

{{- range .Cfg.Params}}
func Test{{$paramUC.CamelCaseName .}}(t *testing.T) {
	// write your test
	{{$pkgName}}.Get{{$paramUC.CamelCaseName .}}()
}

func TestSet{{$paramUC.CamelCaseName .}}(t *testing.T) {
	// write your test
	{{$pkgName}}.Set{{$paramUC.CamelCaseName .}}({{ $paramUC.GetDefaultStr .}})
}
{{- end}}
`
