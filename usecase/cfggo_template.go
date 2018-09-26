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
  {{- if .Cfg.CfgFileParam.Name }}
	"github.com/pkg/errors"
  {{- end}}
  {{- if .CfgUC.HasFlag $flagUC .Cfg}}
	"github.com/spf13/pflag"
  {{- end}}
	"github.com/spf13/viper"
)

const (
{{- if .Cfg.CfgFileParam.Name }}
	{{$paramUC.CamelCaseLowerName .Cfg.CfgFileParam}}Key = "{{.Cfg.CfgFileParam.Name}}"
{{- end}}
{{- range .Cfg.Params}}
	{{$paramUC.CamelCaseLowerName .}}Key = "{{.Name}}"
{{- end}}
)

type (
	// Config manages configuration.
	Config interface {
{{- if .Cfg.CfgFileParam.Name }}
	Get{{$paramUC.CamelCaseName .Cfg.CfgFileParam}}() string
	Set{{$paramUC.CamelCaseName .Cfg.CfgFileParam}}(string)
{{- end}}
{{- range .Cfg.Params}}
		Get{{$paramUC.CamelCaseName .}}() {{$paramUC.GetType .}} 
		Set{{$paramUC.CamelCaseName .}}(value {{$paramUC.GetType .}}) 
{{- end}}
	}

	config struct {
		viper *viper.Viper
	}
)

func initViper(v *viper.Viper) error {
{{- if .Cfg.CfgFileParam.Name }}
  {{- if $envUC.IsBind .Cfg.CfgFileParam.Env $default.Env.Bind }}
	v.BindEnv({{$paramUC.CamelCaseLowerName .Cfg.CfgFileParam}}Key, "{{$envUC.GetName .Cfg.CfgFileParam.Env .Cfg.CfgFileParam.Name $default.Env.Prefix}}")
  {{- end}}
  {{- if $paramUC.IsSetDefault .Cfg.CfgFileParam }}
	v.SetDefault({{$paramUC.CamelCaseLowerName .Cfg.CfgFileParam}}Key, {{ $paramUC.GetDefaultStr .Cfg.CfgFileParam}})
  {{- end}}
  {{- if $flagUC.IsBind .Cfg.CfgFileParam.Flag $default.Flag.Bind }}
    {{- if .Cfg.CfgFileParam.Flag.Short}}
	pflag.{{$paramUC.GetPFlagName .Cfg.CfgFileParam}}P("{{$paramUC.GetFlagName .Cfg.CfgFileParam}}", "{{.Cfg.CfgFileParam.Flag.Short}}", {{$paramUC.GetDefaultStr .Cfg.CfgFileParam}}, "{{$paramUC.GetFlagDescription .Cfg.CfgFileParam}}")
    {{- else}}
	pflag.{{$paramUC.GetPFlagName .Cfg.CfgFileParam}}("{{$paramUC.GetFlagName .Cfg.CfgFileParam}}", {{$paramUC.GetDefaultStr .Cfg.CfgFileParam}}, "{{$paramUC.GetFlagDescription .Cfg.CfgFileParam}}")
    {{- end}}
	v.BindPFlag({{$paramUC.CamelCaseLowerName .Cfg.CfgFileParam}}Key, pflag.Lookup("{{$paramUC.GetFlagName .Cfg.CfgFileParam}}"))
  {{- end}}
{{ end }}
{{- if .Cfg.Params}}
  {{- range .Cfg.Params}}
    {{- if $envUC.IsBind .Env $default.Env.Bind }}
	v.BindEnv({{$paramUC.CamelCaseLowerName .}}Key, "{{$envUC.GetName .Env .Name $default.Env.Prefix}}")
    {{- end}}
  {{- end}}
  {{- range .Cfg.Params}}
    {{- if $paramUC.IsSetDefault . }}
	v.SetDefault({{$paramUC.CamelCaseLowerName .}}Key, {{ $paramUC.GetDefaultStr .}})
    {{- end}}
    {{- if $flagUC.IsBind .Flag $default.Flag.Bind }}
		  {{- if .Flag.Short}}
	pflag.{{$paramUC.GetPFlagName .}}P("{{$paramUC.GetFlagName .}}", "{{.Flag.Short}}", {{$paramUC.GetDefaultStr .}}, "{{$paramUC.GetFlagDescription .}}")
		  {{- else}}
	pflag.{{$paramUC.GetPFlagName .}}("{{$paramUC.GetFlagName .}}", {{$paramUC.GetDefaultStr .}}, "{{$paramUC.GetFlagDescription .}}")
		  {{- end}}
	v.BindPFlag({{$paramUC.CamelCaseLowerName .}}Key, pflag.Lookup("{{$paramUC.GetFlagName .}}"))
    {{- end}}
  {{- end}}
  {{- if .CfgUC.HasFlag $flagUC .Cfg}}
	pflag.Parse()
  {{- end}}
{{- end}}
{{- if .Cfg.CfgFileParam.Name }}
	v.SetConfigFile(v.GetString({{$paramUC.CamelCaseLowerName .Cfg.CfgFileParam}}Key))
	if err := v.ReadInConfig(); err != nil {
		return errors.Wrap(err, "failed to read config in file")
	}
	return nil
{{- else }}
	return nil
{{- end}}
}

// New returns an initialized configuration.
func New() (Config, error) {
	cfg := &config{
		viper: viper.New(),
	}
	err := initViper(cfg.viper)
	return cfg, err
}

// InitGlobalConfig initializes configuration.
func InitGlobalConfig() error {
	return initViper(viper.GetViper())
}

{{- if .Cfg.CfgFileParam.Name }}
// Get{{$paramUC.CamelCaseName .Cfg.CfgFileParam}} returns a {{.Cfg.CfgFileParam.Name}}.
func (cfg *config) Get{{$paramUC.CamelCaseName .Cfg.CfgFileParam}}() string {
	return cfg.viper.GetString({{$paramUC.CamelCaseLowerName .Cfg.CfgFileParam}}Key)
}

// Set{{$paramUC.CamelCaseName .Cfg.CfgFileParam}} sets a {{.Cfg.CfgFileParam.Name}}.
func (cfg *config) Set{{$paramUC.CamelCaseName .Cfg.CfgFileParam}}(value string) {
	cfg.viper.Set({{$paramUC.CamelCaseLowerName .Cfg.CfgFileParam}}Key, value)
}
{{- end}}

{{- range .Cfg.Params}}

// Get{{$paramUC.CamelCaseName .}} returns a {{.Name}}.
func (cfg *config) Get{{$paramUC.CamelCaseName .}}() {{$paramUC.GetType .}} {
	return cfg.viper.{{$paramUC.GetViperGetterName .}}({{$paramUC.CamelCaseLowerName .}}Key)
}

// Set{{$paramUC.CamelCaseName .}} sets a {{.Name}}.
func (cfg *config) Set{{$paramUC.CamelCaseName .}}(value {{$paramUC.GetType .}}) {
	cfg.viper.Set({{$paramUC.CamelCaseLowerName .}}Key, value)
}
{{- end}}

{{- if .Cfg.CfgFileParam.Name }}
// Get{{$paramUC.CamelCaseName .Cfg.CfgFileParam}} returns a {{.Cfg.CfgFileParam.Name}}.
func Get{{$paramUC.CamelCaseName .Cfg.CfgFileParam}}() string {
	return viper.GetString({{$paramUC.CamelCaseLowerName .Cfg.CfgFileParam}}Key)
}

// Set{{$paramUC.CamelCaseName .Cfg.CfgFileParam}} sets a {{.Cfg.CfgFileParam.Name}}.
func Set{{$paramUC.CamelCaseName .Cfg.CfgFileParam}}(value string) {
	viper.Set({{$paramUC.CamelCaseLowerName .Cfg.CfgFileParam}}Key, value)
}
{{- end}}

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
