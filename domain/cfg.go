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

	// TemplateData is the argument of domain.TemplateRenderer.Render .
	TemplateData struct {
		Cfg    Cfg
		CfgUC  CfgUsecase
		EnvUC  EnvUsecase
		FlagUC FlagUsecase
	}

	// Env represents environment variable.
	Env struct {
		Name   *string `yaml:"name"`
		Prefix *string `yaml:"prefix"`
		Bind   *bool   `yaml:"bind"`
	}
)
