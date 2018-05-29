package domain

type (
	// Env represents environment variable.
	Env struct {
		Name   *string `yaml:"name"`
		Prefix *string `yaml:"prefix"`
		Bind   *bool   `yaml:"bind"`
	}
)
