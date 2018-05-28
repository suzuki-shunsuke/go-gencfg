package domain

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/serenize/snaker"
)

const (
	typeString  = "string"
	typeInt     = "int"
	typeFloat64 = "float64"
)

type (
	// Param represents a parameter.
	Param struct {
		Name        string      `yaml:"name"`
		Type        string      `yaml:"type"`
		Description string      `yaml:"description"`
		Default     interface{} `yaml:"default"`
		Flag        Flag        `yaml:"flag"`
		Env         Env         `yaml:"env"`
	}

	// Global represents a global configuration.
	Global struct {
		Env  Env  `yaml:"env"`
		Flag Flag `yaml:"flag"`
	}
)

// IsSetDefault returns whether the parameter is set the default value.
func (p Param) IsSetDefault() bool {
	return p.Default != nil
}

// GetDefaultStr returns a string embedded the code as default value.
func (p Param) GetDefaultStr() string {
	if p.GetType() == typeString {
		if p.Default == nil {
			return `""`
		}
		return fmt.Sprintf(`"%s"`, p.Default)
	}
	return fmt.Sprintf("%v", p.Default)
}

// GetFlagName returns the flag name.
func (p Param) GetFlagName() string {
	if p.Flag.Name != nil {
		return *p.Flag.Name
	}
	return strings.Replace(p.Name, "_", "-", -1)
}

// GetFlagDescription returns the flag's description.
func (p Param) GetFlagDescription() string {
	if p.Flag.Description != nil {
		return *p.Flag.Description
	}
	return p.Description
}

// GetEnvName returns the environment variable name.
func (p Param) GetEnvName() string {
	name := strings.ToUpper(p.Name)
	if p.Env.Prefix == nil {
		return name
	}
	return fmt.Sprintf("%s%s", strings.ToUpper(*p.Env.Prefix), name)
}

// CamelCaseName returns the camel case parameter name.
func (p Param) CamelCaseName() string {
	return snaker.SnakeToCamel(p.Name)
}

// CamelCaseLowerName returns the camel case lower parameter name.
func (p Param) CamelCaseLowerName() string {
	return snaker.SnakeToCamelLower(p.Name)
}

// GetType returns the name of parameter's data type (ex. "int", "string").
func (p Param) GetType() string {
	if p.Type != "" {
		return p.Type
	}
	if p.Default == nil {
		return "interface{}"
	}
	return reflect.ValueOf(p.Default).Kind().String()
}

// GetPFlagName returns the PFlag name.
func (p Param) GetPFlagName() string {
	switch p.GetType() {
	case typeInt:
		return "Int"
	case typeString:
		return "String"
	case "bool":
		return "Bool"
	case typeFloat64:
		return "Float64"
	default:
		return "String"
	}
}

// GetViperGetterName returns the paramter's getter name.
func (p Param) GetViperGetterName() string {
	switch p.GetType() {
	case typeInt:
		return "GetInt"
	case typeString:
		return "GetString"
	case "bool":
		return "GetBool"
	case typeFloat64:
		return "GetFloat64"
	default:
		return "Get"
	}
}
