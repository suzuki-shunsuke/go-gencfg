# go-gencfg

[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/suzuki-shunsuke/go-gencfg)
[![Build Status](https://travis-ci.org/suzuki-shunsuke/go-gencfg.svg?branch=master)](https://travis-ci.org/suzuki-shunsuke/go-gencfg)
[![codecov](https://codecov.io/gh/suzuki-shunsuke/go-gencfg/branch/master/graph/badge.svg)](https://codecov.io/gh/suzuki-shunsuke/go-gencfg)
[![Go Report Card](https://goreportcard.com/badge/github.com/suzuki-shunsuke/go-gencfg)](https://goreportcard.com/report/github.com/suzuki-shunsuke/go-gencfg)
[![GitHub last commit](https://img.shields.io/github/last-commit/suzuki-shunsuke/go-gencfg.svg)](https://github.com/suzuki-shunsuke/go-gencfg)
[![GitHub tag](https://img.shields.io/github/tag/suzuki-shunsuke/go-gencfg.svg)](https://github.com/suzuki-shunsuke/go-gencfg/releases)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/suzuki-shunsuke/go-gencfg/master/LICENSE)

cli tool to generate code for the application specific configuration.

## Getting Started

Intsall

```
$ go get -d github.com/suzuki-shunsuke/go-gencfg
```

or Download a binary from the [release page](https://github.com/suzuki-shunsuke/go-gencfg/releases).

Check whether gencfg is installed.

```
$ gencfg -v
gencfg version 0.1.0
```

Assume you are in the application's root directory and the application uses [github.com/spf13/viper](https://github.com/spf13/viper).
Generate the gencfg's configuration file voilerplate.

```
$ gencfg init port
create .gencfg.yml
create .gencfg_config.tmpl
```

Note that the argument `port` is your application's parameter. It is optional and you can pass the multiple parameters.

The following `.gencfg.yml` and `.gencfg_config.tmpl` are generated.

```yaml
---
# https://github.com/suzuki-shunsuke/go-gencfg
# https://github.com/suzuki-shunsuke/go-gencfg/blob/master/docs/CONFIGURATION.md
dest: config/config.go
package_name:
package:
template: .gencfg_config.tmpl
formatters:
- gofmt -l -s -w
default:
  env:
    bind: false
  flag:
    bind: false
params:
- name: port
  type: string
  default:
  env:
    name:
    bind: false
  flag:
    name:
    description:
    short:
    bind: false
```

```
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
```

Edit the files as you like.
Then generate the code which wraps viper for your application.

```
$ gencfg gen
create config
create config/config.go
gofmt -l -s -w config/config.go
```

See `config/config.go`.

```golang
// Package config wraps viper for the application
package config

import (
	"github.com/spf13/viper"
)

const (
	portKey = "port"
)

func init() {
}

// GetPort returns a port.
func GetPort() string {
	return viper.GetString(portKey)
}

// SetPort sets a port.
func SetPort(value string) {
	viper.Set(portKey, value)
}
```

Then in the application you can get the port number as the following.

```golang
config.GetPort()
```

You can change the generated file path and the template of the generated file by the configuration file and command line parameters.

## Usage

See [USAGE.md](docs/USAGE.md) .

### validate for ci

`gencfg compare -f` is useful to validate the consistency of configuration file (`.gencfg.yml`) and the generated code.
`gencfg compare -f` returns `0` if the result of `gencfg gen` and existing code is equal and otherwise returns not `0`.
It is useful for CI.

## Configuration

See [CONFIGURATION.md](docs/CONFIGURATION.md) .

## License

[MIT](LICENSE)
