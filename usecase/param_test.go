package usecase_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-ptr"

	"github.com/suzuki-shunsuke/go-gencfg/domain"
	"github.com/suzuki-shunsuke/go-gencfg/usecase"
)

func TestParamUsecaseIsSetDefault(t *testing.T) {
	p := domain.Param{}
	pUC := usecase.ParamUsecase{}
	if pUC.IsSetDefault(p) {
		t.Fatal("pUC.IsSetDefault(p) = true, wanted false")
	}
}

func TestParamUsecaseGetDefaultStr(t *testing.T) {
	p := domain.Param{Type: "string"}
	pUC := usecase.ParamUsecase{}
	exp := `""`
	act := pUC.GetDefaultStr(p)
	if act != exp {
		t.Fatalf(`pUC.GetDefaultStr(p) = "%s", wanted "%s"`, act, exp)
	}

	p = domain.Param{Type: "string", Default: "hello"}
	exp = `"hello"`
	act = pUC.GetDefaultStr(p)
	if act != exp {
		t.Fatalf(`pUC.GetDefaultStr(p) = "%s", wanted "%s"`, act, exp)
	}

	p = domain.Param{Type: "int", Default: 3}
	exp = "3"
	act = pUC.GetDefaultStr(p)
	if act != exp {
		t.Fatalf(`pUC.GetDefaultStr(p) = "%s", wanted "%s"`, act, exp)
	}
}

func TestParamUsecaseGetFlagName(t *testing.T) {
	p := domain.Param{Flag: domain.Flag{Name: ptr.PStr("hello")}}
	pUC := usecase.ParamUsecase{}
	exp := "hello"
	act := pUC.GetFlagName(p)
	if act != exp {
		t.Fatalf(`pUC.GetFlagName(p) = "%s", wanted "%s"`, act, exp)
	}

	p = domain.Param{Name: "foo_bar"}
	exp = "foo-bar"
	act = pUC.GetFlagName(p)
	if act != exp {
		t.Fatalf(`pUC.GetFlagName(p) = "%s", wanted "%s"`, act, exp)
	}
}

func TestParamUsecaseGetFlagDescription(t *testing.T) {
	p := domain.Param{Flag: domain.Flag{Description: ptr.PStr("hello")}}
	pUC := usecase.ParamUsecase{}
	exp := "hello"
	act := pUC.GetFlagDescription(p)
	if act != exp {
		t.Fatalf(`pUC.GetFlagDescription(p) = "%s", wanted "%s"`, act, exp)
	}

	p = domain.Param{Description: "foo_bar"}
	exp = "foo_bar"
	act = pUC.GetFlagDescription(p)
	if act != exp {
		t.Fatalf(`pUC.GetFlagDescription(p) = "%s", wanted "%s"`, act, exp)
	}
}

func TestParamUsecaseGetEnvName(t *testing.T) {
	p := domain.Param{Name: "hello"}
	pUC := usecase.ParamUsecase{}
	exp := "HELLO"
	act := pUC.GetEnvName(p)
	if act != exp {
		t.Fatalf(`pUC.GetEnvName(p) = "%s", wanted "%s"`, act, exp)
	}

	p.Env.Prefix = ptr.PStr("sample")
	exp = "SAMPLEHELLO"
	act = pUC.GetEnvName(p)
	if act != exp {
		t.Fatalf(`pUC.GetEnvName(p) = "%s", wanted "%s"`, act, exp)
	}
}

func TestParamUsecaseCamelCaseName(t *testing.T) {
	pUC := usecase.ParamUsecase{}
	p := domain.Param{Name: "hello_world"}
	exp := "HelloWorld"
	act := pUC.CamelCaseName(p)
	if act != exp {
		t.Fatalf(`pUC.CamelCaseName(p) = "%s", wanted "%s"`, act, exp)
	}
}

func TestParamUsecaseCamelCaseLowerName(t *testing.T) {
	pUC := usecase.ParamUsecase{}
	p := domain.Param{Name: "hello_world"}
	exp := "helloWorld"
	act := pUC.CamelCaseLowerName(p)
	if act != exp {
		t.Fatalf(`pUC.CamelCaseLowerName(p) = "%s", wanted "%s"`, act, exp)
	}
}

func TestParamUsecaseGetType(t *testing.T) {
	pUC := usecase.ParamUsecase{}
	p := domain.Param{Type: "string"}
	exp := "string"
	act := pUC.GetType(p)
	if act != exp {
		t.Fatalf(`pUC.GetType(p) = "%s", wanted "%s"`, act, exp)
	}

	p = domain.Param{}
	exp = "interface{}"
	act = pUC.GetType(p)
	if act != exp {
		t.Fatalf(`pUC.GetType(p) = "%s", wanted "%s"`, act, exp)
	}

	p = domain.Param{Default: true}
	exp = "bool"
	act = pUC.GetType(p)
	if act != exp {
		t.Fatalf(`pUC.GetType(p) = "%s", wanted "%s"`, act, exp)
	}
}

func TestParamUsecaseGetPFlagName(t *testing.T) {
	pUC := usecase.ParamUsecase{}

	p := domain.Param{Type: "string"}
	exp := "String"
	act := pUC.GetPFlagName(p)
	if act != exp {
		t.Fatalf(`pUC.GetPFlagName(p) = "%s", wanted "%s"`, act, exp)
	}

	p = domain.Param{Type: "int"}
	exp = "Int"
	act = pUC.GetPFlagName(p)
	if act != exp {
		t.Fatalf(`pUC.GetPFlagName(p) = "%s", wanted "%s"`, act, exp)
	}

	p = domain.Param{Type: "bool"}
	exp = "Bool"
	act = pUC.GetPFlagName(p)
	if act != exp {
		t.Fatalf(`pUC.GetPFlagName(p) = "%s", wanted "%s"`, act, exp)
	}

	p = domain.Param{Type: "float64"}
	exp = "Float64"
	act = pUC.GetPFlagName(p)
	if act != exp {
		t.Fatalf(`pUC.GetPFlagName(p) = "%s", wanted "%s"`, act, exp)
	}

	p = domain.Param{}
	exp = "String"
	act = pUC.GetPFlagName(p)
	if act != exp {
		t.Fatalf(`pUC.GetPFlagName(p) = "%s", wanted "%s"`, act, exp)
	}
}

func TestParamUsecaseGetViperGetterName(t *testing.T) {
	pUC := usecase.ParamUsecase{}

	p := domain.Param{Type: "string"}
	exp := "GetString"
	act := pUC.GetViperGetterName(p)
	if act != exp {
		t.Fatalf(`pUC.GetViperGetterName(p) = "%s", wanted "%s"`, act, exp)
	}

	p = domain.Param{Type: "int"}
	exp = "GetInt"
	act = pUC.GetViperGetterName(p)
	if act != exp {
		t.Fatalf(`pUC.GetViperGetterName(p) = "%s", wanted "%s"`, act, exp)
	}

	p = domain.Param{Type: "bool"}
	exp = "GetBool"
	act = pUC.GetViperGetterName(p)
	if act != exp {
		t.Fatalf(`pUC.GetViperGetterName(p) = "%s", wanted "%s"`, act, exp)
	}

	p = domain.Param{Type: "float64"}
	exp = "GetFloat64"
	act = pUC.GetViperGetterName(p)
	if act != exp {
		t.Fatalf(`pUC.GetViperGetterName(p) = "%s", wanted "%s"`, act, exp)
	}

	p = domain.Param{}
	exp = "Get"
	act = pUC.GetViperGetterName(p)
	if act != exp {
		t.Fatalf(`pUC.GetViperGetterName(p) = "%s", wanted "%s"`, act, exp)
	}
}
