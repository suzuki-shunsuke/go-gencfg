package usecase

// GenCfgTmpl is the template of the configuration file.
const GenCfgTmpl = `
---
# https://github.com/suzuki-shunsuke/go-gencfg
dest: config/config.go
package_name:
package:
template:
test_template:
formatters:
- gofmt -l -s -w
global:
  env:
    bind: false
    prefix: SAMPLE_
  flag:
    bind: false
params:
{{- range .}}
- name: {{.}}
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
{{- else}}
- name:
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
{{end}}
`
