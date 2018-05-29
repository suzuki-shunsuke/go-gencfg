package domain

type (
	// Global represents a global configuration.
	Global struct {
		Env  Env  `yaml:"env"`
		Flag Flag `yaml:"flag"`
	}

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
		Cfg     Cfg
		CfgUC   CfgUsecase
		EnvUC   EnvUsecase
		FlagUC  FlagUsecase
		ParamUC ParamUsecase
	}

	// Env represents environment variable.
	Env struct {
		Name   *string `yaml:"name"`
		Prefix *string `yaml:"prefix"`
		Bind   *bool   `yaml:"bind"`
	}

	// Flag represents a command line flag.
	Flag struct {
		Description *string `yaml:"description"`
		Name        *string `yaml:"name"`
		Short       *string `yaml:"short"`
		Bind        *bool   `yaml:"bind"`
	}

	// Param represents a parameter.
	Param struct {
		Name        string      `yaml:"name"`
		Type        string      `yaml:"type"`
		Description string      `yaml:"description"`
		Default     interface{} `yaml:"default"`
		Flag        Flag        `yaml:"flag"`
		Env         Env         `yaml:"env"`
	}
)
