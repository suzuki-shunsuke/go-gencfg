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
  {{- if .CfgUC.HasFlag $flagUC .Cfg}}
	"github.com/spf13/pflag"
  {{- end}}
	"github.com/spf13/viper"
)

const (
{{- range .Cfg.Params}}
	{{$paramUC.CamelCaseLowerName .}}Key = "{{.Name}}"
{{- end}}
)

type (
	// Config manages configuration.
	Config interface {
{{- range .Cfg.Params}}
		Get{{$paramUC.CamelCaseName .}}() {{$paramUC.GetType .}} 
		Set{{$paramUC.CamelCaseName .}}(value {{$paramUC.GetType .}}) 
{{- end}}
{{- if .Cfg.CfgFileParam.Name }}
		ReadFile() error
{{- end}}
	}

	config struct {
		viper *viper.Viper
	}
)

func initViper(v *viper.Viper) {
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
{{- end}}
}

{{- if .Cfg.CfgFileParam.Name }}
// ReadFile reads configuration in file.
func ReadFile() error {
	return viper.ReadInConfig()
}

// ReadFile reads configuration in file.
func (cfg *config) ReadFile() error {
	return cfg.viper.ReadInConfig()
}
{{- end}}

// New returns an initialized configuration.
func New() Config {
	cfg := &config{
		viper: viper.New(),
	}
	initViper(cfg.viper)
	return cfg
}

// InitGlobalConfig initializes configuration.
func InitGlobalConfig() {
	initViper(viper.GetViper())
}

{{- range .Cfg.Params}}

// Get{{$paramUC.CamelCaseName .}} returns a {{.Name}}.
func (cfg *config) Get{{$paramUC.CamelCaseName .}}() {{$paramUC.GetType .}} {
	return cfg.viper.{{$paramUC.GetViperGetterName .}}({{$paramUC.CamelCaseLowerName .}}Key)
}

// Set{{$paramUC.CamelCaseName .}} sets a {{.Name}}.
func (cfg *config) Set{{$paramUC.CamelCaseName .}}(value {{$paramUC.GetType .}}) {
	cfg.viper.Set({{$paramUC.CamelCaseLowerName .}}Key, value)
}

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
