package usecase

// DefaultCfgTmpl is the default template of the generated wrapper code.
const DefaultCfgTmpl = `
{{ $global := .Cfg.Global -}}
{{ $envUC := .EnvUC -}}
{{ $flagUC := .FlagUC -}}
// Package config wraps viper for the application
package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

{{if .Cfg.Params -}}
const (
{{- range .Cfg.Params}}
	{{.CamelCaseLowerName}}Key = "{{.Name}}"
{{- end}}
)
{{- end}}

func init() {
{{- if .Cfg.Params}}
  {{- range .Cfg.Params}}
    {{- if $envUC.IsBind .Env $global.Env.Bind }}
	viper.BindEnv({{.CamelCaseLowerName}}Key, "{{$envUC.GetName .Env .Name $global.Env.Prefix}}")
    {{- end}}
  {{- end}}
  {{- range .Cfg.Params}}
    {{- if .IsSetDefault }}
	viper.SetDefault({{.CamelCaseLowerName}}Key, {{.GetDefaultStr}})
    {{- end}}
    {{- if $flagUC.IsBind .Flag $global.Flag.Bind }}
		  {{- if .Flag.Short}}
	pflag.{{.GetPFlagName}}P("{{.GetFlagName}}", "{{.Flag.Short}}", {{.GetDefaultStr}}, "{{.GetFlagDescription}}")
		  {{- else}}
	pflag.{{.GetPFlagName}}("{{.GetFlagName}}", {{.GetDefaultStr}}, "{{.GetFlagDescription}}")
		  {{- end}}
	viper.BindPFlag({{.CamelCaseLowerName}}Key, pflag.Lookup("{{.GetFlagName}}"))
    {{- end}}
  {{- end}}
  {{- if .CfgUC.HasFlag $flagUC .Cfg}}
	pflag.Parse()
  {{- end}}
{{- end}}
}

{{- range .Cfg.Params}}

// Get{{.CamelCaseName}} returns a {{.Name}}.
func Get{{.CamelCaseName}}() {{.GetType}} {
	return viper.{{.GetViperGetterName}}({{.CamelCaseLowerName}}Key)
}

// Set{{.CamelCaseName}} sets a {{.Name}}.
func Set{{.CamelCaseName}}(value {{.GetType}}) {
	viper.Set({{.CamelCaseLowerName}}Key, value)
}
{{- end}}
`
