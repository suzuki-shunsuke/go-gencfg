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
  {{- if .Cfg.CfgFilePath }}
	"github.com/pkg/errors"
  {{- end}}
  {{- if .CfgUC.HasFlag $flagUC .Cfg}}
	"github.com/spf13/pflag"
  {{- end}}
	"github.com/spf13/viper"
)

{{if .Cfg.Params -}}
const (
{{- range .Cfg.Params}}
	{{$paramUC.CamelCaseLowerName .}}Key = "{{.Name}}"
{{- end}}
)
{{- end}}

type (
	// Config manages configuration.
	Config interface {
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
{{- if .Cfg.CfgFilePath }}
	v.SetConfigFile("{{.Cfg.CfgFilePath}}")
	if err := v.ReadInConfig(); err != nil {
		return errors.Wrap(err, "failed to read config in file")
	}
	return nil
{{ else }}
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

func InitGlobalConfig() error {
	return initViper(viper.GetViper())
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
