package config_test

import (
	"testing"

	"github.com/urfave/cli"

	"github.com/suzuki-shunsuke/go-gencfg/internal/config"
)

type (
	CLIContext struct{}
)

func (ctx CLIContext) String(key string) string {
	return key
}

func (ctx CLIContext) Bool(key string) bool {
	return true
}

func (ctx CLIContext) Args() cli.Args {
	return cli.Args{}
}

func TestGetCfgPath(t *testing.T) {
	ctx := CLIContext{}
	exp := "config"
	if act := config.GetCfgPath(ctx); act != exp {
		t.Fatalf(`GetCfgPath() = "%s", wanted "%s"`, act, exp)
	}
}

func TestGetDestPath(t *testing.T) {
	ctx := CLIContext{}
	exp := "dest"
	if act := config.GetDestPath(ctx); act != exp {
		t.Fatalf(`GetDestPath() = "%s", wanted "%s"`, act, exp)
	}
}

func TestGetTemplatePath(t *testing.T) {
	ctx := CLIContext{}
	exp := "template"
	if act := config.GetTemplatePath(ctx); act != exp {
		t.Fatalf(`GetTemplatePath() = "%s", wanted "%s"`, act, exp)
	}
}

func TestGetParamNames(t *testing.T) {
	ctx := CLIContext{}
	if act := config.GetParamNames(ctx); len(act) != 0 {
		t.Fatalf(`GetParamNames() = %v, wanted empty array`, act)
	}
}

func TestGetFailure(t *testing.T) {
	ctx := CLIContext{}
	if !config.GetFailure(ctx) {
		t.Fatal(`GetFailure() = false, wanted true`)
	}
}

func TestGetQuiet(t *testing.T) {
	ctx := CLIContext{}
	if !config.GetQuiet(ctx) {
		t.Fatal(`GetQuiet() = false, wanted true`)
	}
}
