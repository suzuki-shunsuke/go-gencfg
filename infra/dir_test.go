package infra_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-gencfg/infra"
)

func TestDirMakerMake(t *testing.T) {
	maker := infra.DirMaker{}
	if err := maker.Make("/root/hoge"); err == nil {
		t.Fatalf("mkdir /root/hoge should not be permitted")
	}
}
