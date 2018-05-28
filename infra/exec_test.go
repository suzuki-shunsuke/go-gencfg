package infra_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-gencfg/infra"
)

func TestOSExecExecuter(t *testing.T) {
	executer := infra.OSExecExecuter{}
	if err := executer.Exec("echo", "hoge"); err != nil {
		t.Fatal(err)
	}
}
