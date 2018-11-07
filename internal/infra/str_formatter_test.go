package infra_test

import (
	"testing"

	"github.com/suzuki-shunsuke/gomic/gomic"

	"github.com/suzuki-shunsuke/go-gencfg/internal/infra"
	"github.com/suzuki-shunsuke/go-gencfg/internal/test"
)

func TestStrFormatterFormat(t *testing.T) {
	formatter := infra.StrFormatter{}
	reader := test.NewFileReaderMock(t, gomic.DoNothing)
	reader.SetFakeRead("hello", nil)
	executer := test.NewCmdExecuterMock(t, gomic.DoNothing)

	t.Run("positive", func(t *testing.T) {
		if _, err := formatter.Format("/tmp/hello.go", "", nil, reader, executer); err != nil {
			t.Fatal(err)
		}
	})
}
