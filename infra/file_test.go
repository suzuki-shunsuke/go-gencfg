package infra_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/suzuki-shunsuke/go-gencfg/infra"
)

func TestFileReaderRead(t *testing.T) {
	reader := infra.FileReader{}
	t.Run("positive", func(t *testing.T) {
		tmpFile, err := ioutil.TempFile("", "")
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(tmpFile.Name())
		p := tmpFile.Name()
		exp := "test hello world"
		if _, err := tmpFile.Write([]byte(exp)); err != nil {
			t.Fatal(err)
		}
		if err := tmpFile.Close(); err != nil {
			t.Fatal(err)
		}
		act, err := reader.Read(p)
		if err != nil {
			t.Fatal(err)
		}
		if act != exp {
			t.Fatalf(`FileReader.Read() = "%s", wanted "%s"`, act, exp)
		}
	})
	t.Run("src is directory", func(t *testing.T) {
		if _, err := reader.Read("/tmp"); err == nil {
			t.Fatal("/tmp is directory")
		}
	})
	t.Run("src is not found", func(t *testing.T) {
		if _, err := reader.Read("/fooooo"); err == nil {
			t.Fatal("/fooooo is not found")
		}
	})
}

func TestFileWriterWrite(t *testing.T) {
	reader := infra.FileReader{}
	writer := infra.FileWriter{}
	t.Run("positive", func(t *testing.T) {
		tmpFile, err := ioutil.TempFile("", "")
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(tmpFile.Name())
		p := tmpFile.Name()
		if err := tmpFile.Close(); err != nil {
			t.Fatal(err)
		}
		exp := "test hello world"
		if err := writer.Write(p, []byte(exp)); err != nil {
			t.Fatal(err)
		}
		act, err := reader.Read(p)
		if err != nil {
			t.Fatal(err)
		}
		if act != exp {
			t.Fatalf(`file content is "%s", wanted "%s"`, act, exp)
		}
		// test not append
		if err := writer.Write(p, []byte(exp)); err != nil {
			t.Fatal(err)
		}
		act, err = reader.Read(p)
		if err != nil {
			t.Fatal(err)
		}
		if act != exp {
			t.Fatalf(`file content is "%s", wanted "%s"`, act, exp)
		}
	})
	t.Run("dest is directory", func(t *testing.T) {
		if err := writer.Write("/tmp", []byte("test")); err == nil {
			t.Fatal("/tmp is directory")
		}
	})
}

func TestFileExistExist(t *testing.T) {
	exist := infra.FileExist{}
	if !exist.Exist("/tmp") {
		t.Fatal("/tmp is not found")
	}
	if exist.Exist("/foooo") {
		t.Fatal("/foooo is found")
	}
}
