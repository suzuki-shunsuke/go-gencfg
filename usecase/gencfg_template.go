package usecase

// GenCfgTmpl is the template of the configuration file.
const GenCfgTmpl = `
---
# https://github.com/suzuki-shunsuke/go-gencfg
# https://github.com/suzuki-shunsuke/go-gencfg/blob/master/docs/CONFIGURATION.md
dest: config/config.go
package_name:
package:
template: .gencfg_config.tmpl
config_file_param:
  name: config_file_path
formatters:
- gofmt -l -s -w
default:
  env:
    bind: false
  flag:
    bind: false
params:
{{- range .}}
- name: {{.}}
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
{{- else}}
- name:
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
{{end}}
`
