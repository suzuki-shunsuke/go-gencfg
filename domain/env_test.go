package domain_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-ptr"

	"github.com/suzuki-shunsuke/go-gencfg/domain"
)

func TestEnvIsBind(t *testing.T) {
	env := domain.Env{}
	if env.IsBind(nil) {
		t.Fatal("env.IsBind(nil) = true, wanted false")
	}
	if !env.IsBind(ptr.PBool(true)) {
		t.Fatal("env.IsBind(nil) = false, wanted true")
	}
	env.Bind = ptr.PBool(true)
	if !env.IsBind(nil) {
		t.Fatal("env.IsBind(nil) = false, wanted true")
	}
}

func TestEnvGetPrefix(t *testing.T) {
	env := domain.Env{}
	exp := ""
	act := env.GetPrefix(nil)
	if act != exp {
		t.Fatalf(`env.GetPrefix(nil) = "%s", wanted "%s"`, act, exp)
	}
	exp = "hello"
	act = env.GetPrefix(ptr.PStr(exp))
	if act != exp {
		t.Fatalf(`env.GetPrefix(*"hello") = "%s", wanted "%s"`, act, exp)
	}
}
