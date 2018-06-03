package usecase

// DefaultCfgTmpl is the default template of the generated wrapper code.
const DefaultCfgTmpl = `
{{ $default := .Cfg.Default -}}
{{ $envUC := .EnvUC -}}
{{ $flagUC := .FlagUC -}}
{{ $paramUC := .ParamUC -}}
{{ $pkgName := .CfgUC.GetPkgName .Cfg -}}
// Package {{$pkgName}} wraps viper for the application
package {{$pkgName}}

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

{{if .Cfg.Params -}}
const (
{{- range .Cfg.Params}}
	{{$paramUC.CamelCaseLowerName .}}Key = "{{.Name}}"
{{- end}}
)
{{- end}}

func init() {
{{- if .Cfg.Params}}
  {{- range .Cfg.Params}}
    {{- if $envUC.IsBind .Env $default.Env.Bind }}
	viper.BindEnv({{$paramUC.CamelCaseLowerName .}}Key, "{{$envUC.GetName .Env .Name $default.Env.Prefix}}")
    {{- end}}
  {{- end}}
  {{- range .Cfg.Params}}
    {{- if $paramUC.IsSetDefault . }}
	viper.SetDefault({{$paramUC.CamelCaseLowerName .}}Key, {{ $paramUC.GetDefaultStr .}})
    {{- end}}
    {{- if $flagUC.IsBind .Flag $default.Flag.Bind }}
		  {{- if .Flag.Short}}
	pflag.{{$paramUC.GetPFlagName .}}P("{{$paramUC.GetFlagName .}}", "{{.Flag.Short}}", {{$paramUC.GetDefaultStr .}}, "{{$paramUC.GetFlagDescription .}}")
		  {{- else}}
	pflag.{{$paramUC.GetPFlagName .}}("{{$paramUC.GetFlagName .}}", {{$paramUC.GetDefaultStr .}}, "{{$paramUC.GetFlagDescription .}}")
		  {{- end}}
	viper.BindPFlag({{$paramUC.CamelCaseLowerName .}}Key, pflag.Lookup("{{$paramUC.GetFlagName .}}"))
    {{- end}}
  {{- end}}
  {{- if .CfgUC.HasFlag $flagUC .Cfg}}
	pflag.Parse()
  {{- end}}
{{- end}}
}

{{- range .Cfg.Params}}

// Get{{$paramUC.CamelCaseName .}} returns a {{.Name}}.
func Get{{$paramUC.CamelCaseName .}}() {{$paramUC.GetType .}} {
	return viper.{{$paramUC.GetViperGetterName .}}({{$paramUC.CamelCaseLowerName .}}Key)
}

// Set{{$paramUC.CamelCaseName .}} sets a {{.Name}}.
func Set{{$paramUC.CamelCaseName .}}(value {{$paramUC.GetType .}}) {
	viper.Set({{$paramUC.CamelCaseLowerName .}}Key, value)
}
{{- end}}
`
