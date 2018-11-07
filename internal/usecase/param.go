package usecase

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/serenize/snaker"

	"github.com/suzuki-shunsuke/go-gencfg/internal/domain"
)

const (
	typeString  = "string"
	typeInt     = "int"
	typeFloat64 = "float64"
	typeBool    = "bool"
)

type (
	// ParamUsecase represents a parameter.
	ParamUsecase struct{}
)

// IsSetDefault returns whether the parameter is set the default value.
func (pUC ParamUsecase) IsSetDefault(p domain.Param) bool {
	return p.Default != nil
}

// GetDefaultStr returns a string embedded the code as default value.
func (pUC ParamUsecase) GetDefaultStr(p domain.Param) string {
	if p.Default == nil {
		switch pUC.GetType(p) {
		case typeString:
			return `""`
		case typeInt:
			return "0"
		case typeFloat64:
			return "0.0"
		case typeBool:
			return "false"
		}
	}
	if pUC.GetType(p) == typeString {
		return fmt.Sprintf(`"%s"`, p.Default)
	}
	return fmt.Sprintf("%v", p.Default)
}

// GetFlagName returns the flag name.
func (pUC ParamUsecase) GetFlagName(p domain.Param) string {
	if p.Flag.Name != nil {
		return *p.Flag.Name
	}
	return strings.Replace(p.Name, "_", "-", -1)
}

// GetFlagDescription returns the flag's description.
func (pUC ParamUsecase) GetFlagDescription(p domain.Param) string {
	if p.Flag.Description != nil {
		return *p.Flag.Description
	}
	return p.Description
}

// GetEnvName returns the environment variable name.
func (pUC ParamUsecase) GetEnvName(p domain.Param) string {
	name := strings.ToUpper(p.Name)
	if p.Env.Prefix == nil {
		return name
	}
	return fmt.Sprintf("%s%s", strings.ToUpper(*p.Env.Prefix), name)
}

// CamelCaseName returns the camel case parameter name.
func (pUC ParamUsecase) CamelCaseName(p domain.Param) string {
	if p.CamelCaseName != "" {
		return p.CamelCaseName
	}
	return snaker.SnakeToCamel(p.Name)
}

// CamelCaseLowerName returns the camel case lower parameter name.
func (pUC ParamUsecase) CamelCaseLowerName(p domain.Param) string {
	if p.CamelCaseName != "" {
		return fmt.Sprintf(
			"%s%s", strings.ToLower(p.CamelCaseName[0:1]),
			p.CamelCaseName[1:])
	}
	return snaker.SnakeToCamelLower(p.Name)
}

// GetType returns the name of parameter's data type (ex. "int", "string").
func (pUC ParamUsecase) GetType(p domain.Param) string {
	if p.Type != "" {
		return p.Type
	}
	if p.Default == nil {
		return "interface{}"
	}
	return reflect.ValueOf(p.Default).Kind().String()
}

// GetPFlagName returns the PFlag name.
func (pUC ParamUsecase) GetPFlagName(p domain.Param) string {
	switch pUC.GetType(p) {
	case typeInt:
		return "Int"
	case typeString:
		return "String"
	case typeBool:
		return "Bool"
	case typeFloat64:
		return "Float64"
	default:
		return "String"
	}
}

// GetViperGetterName returns the paramter's getter name.
func (pUC ParamUsecase) GetViperGetterName(p domain.Param) string {
	switch pUC.GetType(p) {
	case typeInt:
		return "GetInt"
	case typeString:
		return "GetString"
	case typeBool:
		return "GetBool"
	case typeFloat64:
		return "GetFloat64"
	default:
		return "Get"
	}
}
