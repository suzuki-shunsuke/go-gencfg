package infra_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/suzuki-shunsuke/go-gencfg/internal/infra"
)

func TestDirMakerMake(t *testing.T) {
	dir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)
	maker := infra.DirMaker{}
	if err := maker.Make(filepath.Join(dir, "hoge")); err != nil {
		t.Fatal(err)
	}
}
