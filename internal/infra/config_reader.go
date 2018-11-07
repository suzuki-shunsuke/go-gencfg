package infra

import (
	"os"

	"github.com/go-yaml/yaml"

	"github.com/suzuki-shunsuke/go-gencfg/internal/domain"
)

type (
	// YAMLCfgReader is reads the configuration file as YAML.
	YAMLCfgReader struct{}
)

// Read reads the configuration file and parses as YAML and returns domain.Cfg .
func (reader YAMLCfgReader) Read(src string) (domain.Cfg, error) {
	data := domain.Cfg{}
	f, err := os.Open(src)
	if err != nil {
		return data, err
	}
	defer f.Close()
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&data)
	return data, err
}
