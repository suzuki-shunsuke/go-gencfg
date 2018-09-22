# go-gencfg

[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/suzuki-shunsuke/go-gencfg)
[![CircleCI](https://circleci.com/gh/suzuki-shunsuke/go-gencfg.svg?style=svg)](https://circleci.com/gh/suzuki-shunsuke/go-gencfg)
[![codecov](https://codecov.io/gh/suzuki-shunsuke/go-gencfg/branch/master/graph/badge.svg)](https://codecov.io/gh/suzuki-shunsuke/go-gencfg)
[![Go Report Card](https://goreportcard.com/badge/github.com/suzuki-shunsuke/go-gencfg)](https://goreportcard.com/report/github.com/suzuki-shunsuke/go-gencfg)
[![GitHub last commit](https://img.shields.io/github/last-commit/suzuki-shunsuke/go-gencfg.svg)](https://github.com/suzuki-shunsuke/go-gencfg)
[![GitHub tag](https://img.shields.io/github/tag/suzuki-shunsuke/go-gencfg.svg)](https://github.com/suzuki-shunsuke/go-gencfg/releases)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/suzuki-shunsuke/go-gencfg/master/LICENSE)

cli tool to generate code for the application specific configuration.

## Install

```
$ go get -u github.com/suzuki-shunsuke/go-gencfg
```

or Download a binary from the [release page](https://github.com/suzuki-shunsuke/go-gencfg/releases).

Check whether gencfg is installed.

```
$ gencfg -v
gencfg version 0.1.0
```

## Getting Started

1. Generate configuration file.

```
$ gencfg init
create .gencfg.yml
create .gencfg_config.tmpl
```

2. Edit configuration file according to the application.

Here is the example.

```yaml
dest: config/config.go
package_name:
package:
template: .gencfg_config.tmpl
formatters:
- gofmt -l -s -w
default:
  env:
    bind: true
  flag:
    bind: false
params:
- name: user
  type: string
- name: port
  type: int
  default: 6000
- name: access_token
  type: string
```

3. Generate code.

```
$ gencfg gen
```

4. The end. You can manage configuration with generated code.

Use global configuration.

```go
config.InitGlobalConfig()
config.GetUser()
config.SetUser("foo")
```

Or use local configuration.

```go
cfg, _ := config.New()
cfg.GetUser()
cfg.SetUser("foo")
```

When you update the configuration file `.gencfg.yml`, please run `gencfg gen` again.

## Why we use this tool?

1. We can manage the application configuration items with YAML. New developers can understand what the application configuration items is by YAML.
2. We don't have to write obvious code. We can focus on the application logic.
3. We can access configuration with better interfaces. viper specifies the configuration item by string key. This may cause the typo of the key. The typo of key doesn't raise compile error and probably we can't find the typo even with linter. On the other hand, typo of method name raise the compile error so we can find the typo easily.

```go
viper.GetString("user")
```

```go
config.GetUser()
```

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
