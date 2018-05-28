package domain

import (
	"sort"
)

type (
	// Cfg represents the configuration.
	Cfg struct {
		Formatters       []string `yaml:"formatters"`
		Dest             string   `yaml:"dest"`
		TemplatePath     string   `yaml:"template"`
		TestTemplatePath string   `yaml:"test_template"`
		Global           Global   `yaml:"global"`
		Params           []Param  `yaml:"params"`
	}
)

// HasFlag returns whether there is a bound flag.
func (cfg Cfg) HasFlag() bool {
	for _, p := range cfg.Params {
		if p.Flag.IsBind(cfg.Global.Flag.Bind) {
			return true
		}
	}
	return false
}

// Update updates cfg before rendering.
func (cfg Cfg) Update() {
	sort.Sort(Params(cfg.Params))
}
