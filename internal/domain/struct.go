package domain

type (
	// Default represents a default configuration.
	Default struct {
		Env  Env                    `yaml:"env"`
		Flag Flag                   `yaml:"flag"`
		Ext  map[string]interface{} `yaml:"ext"`
	}

	// Cfg represents the configuration.
	Cfg struct {
		Formatters   []string               `yaml:"formatters"`
		Dest         string                 `yaml:"dest"`
		Pkg          string                 `yaml:"package"`
		PkgName      string                 `yaml:"package_name"`
		TemplatePath string                 `yaml:"template"`
		CfgFileParam Param                  `yaml:"config_file_param"`
		Default      Default                `yaml:"default"`
		Params       []Param                `yaml:"params"`
		Ext          map[string]interface{} `yaml:"ext"`
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
		Name   *string                `yaml:"name"`
		Prefix *string                `yaml:"prefix"`
		Bind   *bool                  `yaml:"bind"`
		Ext    map[string]interface{} `yaml:"ext"`
	}

	// Flag represents a command line flag.
	Flag struct {
		Description *string                `yaml:"description"`
		Name        *string                `yaml:"name"`
		Short       *string                `yaml:"short"`
		Bind        *bool                  `yaml:"bind"`
		Ext         map[string]interface{} `yaml:"ext"`
	}

	// Param represents a parameter.
	Param struct {
		Name          string                 `yaml:"name"`
		CamelCaseName string                 `yaml:"camel_case_name"`
		Type          string                 `yaml:"type"`
		Description   string                 `yaml:"description"`
		Default       interface{}            `yaml:"default"`
		Flag          Flag                   `yaml:"flag"`
		Env           Env                    `yaml:"env"`
		Ext           map[string]interface{} `yaml:"ext"`
	}
)
