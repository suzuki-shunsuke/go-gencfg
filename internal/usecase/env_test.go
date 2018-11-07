package usecase_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-ptr"

	"github.com/suzuki-shunsuke/go-gencfg/internal/domain"
	"github.com/suzuki-shunsuke/go-gencfg/internal/usecase"
)

func TestEnvIsBind(t *testing.T) {
	env := domain.Env{}
	envUC := usecase.EnvUsecase{}
	if envUC.IsBind(env, nil) {
		t.Fatal("env.IsBind(nil) = true, wanted false")
	}
	if !envUC.IsBind(env, ptr.PBool(true)) {
		t.Fatal("env.IsBind(nil) = false, wanted true")
	}
	env.Bind = ptr.PBool(true)
	if !envUC.IsBind(env, nil) {
		t.Fatal("env.IsBind(nil) = false, wanted true")
	}
	env.Bind = nil
	env.Prefix = ptr.PStr("hoge")
	if !envUC.IsBind(env, nil) {
		t.Fatal("env.IsBind(nil) = false, wanted true")
	}
}

func TestEnvGetPrefix(t *testing.T) {
	env := domain.Env{}
	envUC := usecase.EnvUsecase{}
	exp := ""
	act := envUC.GetPrefix(env, nil)
	if act != exp {
		t.Fatalf(`envUC.GetPrefix(env, nil) = "%s", wanted "%s"`, act, exp)
	}
	exp = "hello"
	act = envUC.GetPrefix(env, ptr.PStr(exp))
	if act != exp {
		t.Fatalf(`envUC.GetPrefix(env, *"hello") = "%s", wanted "%s"`, act, exp)
	}
	env.Prefix = ptr.PStr(exp)
	act = envUC.GetPrefix(env, nil)
	if act != exp {
		t.Fatalf(`env.GetPrefix(env, nil) = "%s", wanted "%s"`, act, exp)
	}
}

func TestEnvGetName(t *testing.T) {
	env := domain.Env{}
	envUC := usecase.EnvUsecase{}
	exp := ""
	act := envUC.GetName(env, "", nil)
	if act != exp {
		t.Fatalf(`env.GetName(env, "", nil) = "%s", wanted "%s"`, act, exp)
	}
	exp = "hello"
	env.Name = ptr.PStr(exp)
	act = envUC.GetName(env, "", nil)
	if act != exp {
		t.Fatalf(`env.GetName(env, "", nil) = "%s", wanted "%s"`, act, exp)
	}
}
