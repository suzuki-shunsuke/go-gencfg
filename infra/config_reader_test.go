package infra_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/suzuki-shunsuke/go-gencfg/infra"
)

const (
	sampleYAML = `
---
# https://github.com/suzuki-shunsuke/go-gencfg
dest: config/config.go
template:
test_template:
default:
  env:
    bind: false
    prefix: SAMPLE_
  flag:
    bind: false
params:
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
`
)

func TestYAMLCfgReaderRead(t *testing.T) {
	reader := infra.YAMLCfgReader{}
	t.Run("positive", func(t *testing.T) {
		tmpFile, err := ioutil.TempFile("", "")
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(tmpFile.Name())
		p := tmpFile.Name()
		if _, err := tmpFile.Write([]byte(sampleYAML)); err != nil {
			t.Fatal(err)
		}
		if err := tmpFile.Close(); err != nil {
			t.Fatal(err)
		}
		cfg, err := reader.Read(p)
		if err != nil {
			t.Fatal(err)
		}
		exp := "config/config.go"
		if cfg.Dest != exp {
			t.Fatalf(`cfg.Dest = "%s", wanted "%s"`, cfg.Dest, exp)
		}
	})
	t.Run("src is directory", func(t *testing.T) {
		if _, err := reader.Read("/tmp"); err == nil {
			t.Fatal("/tmp is directory")
		}
	})
	t.Run("src is not found", func(t *testing.T) {
		if _, err := reader.Read("/fooooo"); err == nil {
			t.Fatal("/fooooo is not found")
		}
	})
}
