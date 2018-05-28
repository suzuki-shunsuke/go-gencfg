# go-gencfg

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
create gencfg.yml
```

Notice that the argument `port` is your application's parameter. It is optional and you can pass the multiple parameters.

The following `gencfg.yml` is generated.

```yaml
---
# https://github.com/suzuki-shunsuke/go-gencfg
dest: config/config.go
template:
test_template:
global:
  env:
    bind: false
    prefix: SAMPLE_
  flag:
    bind: false
params:
- name: port
  type: string
	description:
  default:
  env:
    name:
    prefix: SAMPLE_
    bind: false
  flag:
	  name:
    description:
    short:
		bind: true
```

Edit the configuration file as you like.
Then generate the code which wraps viper for your application.

```
$ gencfg gen
create config
create config/config.go
gofmt -l -s -w config/config.go
create config
create config/config_test.go
gofmt -l -s -w config/config_test.go
```

`config` directory and `config/config.go` and `config/config_test.go` are generated.

Show the `config/config.go`.

```golang
// config wraps viper for the application
package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	portKey = "port"
)

func init() {
	viper.BindEnv(portKey, "SAMPLE_PORT")
	viper.SetDefault(portKey, 3000)
	pflag.Int("port", 3000, "app port number")
	viper.BindPFlag(portKey, pflag.Lookup("port"))
	pflag.Parse()
}

// GetPort returns a port.
func GetPort() int {
	return viper.GetInt(portKey)
}

// SetPort sets a port.
func SetPort(value int) {
	viper.Set(portKey, value)
}
```

Then in the application you can get the port number as the following.

```golang
config.GetPort()
```

You can change the generated file path and the template of the generated file by the configuration file and command line parameters.

## Usage

See [USAGE](docs/USAGE.md) .

## Install

```bash
$ go get -d github.com/suzuki-shunsuke/go-gencfg
```

or Download a binary from the [release page](https://github.com/suzuki-shunsuke/go-gencfg/releases).

Check whether gencfg is installed.

```
$ gencfg -v
gencfg version 0.1.0
```

### validate for ci

`gencfg compare -f` is useful to validate the consistency of configuration file (`gencfg.json`) and the generated code.
`gencfg compare -f` returns `0` if the result of `gencfg gen` and existing code is equal and otherwise returns not `0`.
It is useful for CI.

## Congiguration

See [CONFIGURATION.md](docs/CONFIGURATION.md) .

## License

[MIT](LICENSE)
