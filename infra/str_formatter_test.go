package infra_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-gencfg/infra"
)

func TestStrFormatterFormat(t *testing.T) {
	formatter := infra.StrFormatter{}
	reader := NoopFileReader{content: "hello"}
	executer := NoopCmdExecuter{}
	t.Run("positive", func(t *testing.T) {
		if _, err := formatter.Format("/tmp/hello.go", "", nil, reader, executer); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("failed to read", func(t *testing.T) {
	})
}
