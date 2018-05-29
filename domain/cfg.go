package domain

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

	// CfgUsecase represents application logic about Cfg
	CfgUsecase interface {
		HasFlag(Cfg) bool
		Update(*Cfg)
	}

	// TemplateData is the argument of domain.TemplateRenderer.Render .
	TemplateData struct {
		Cfg   Cfg
		CfgUC CfgUsecase
	}
)
