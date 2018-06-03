package infra

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/suzuki-shunsuke/go-gencfg/domain"
)

type (
	// StrFormatter formats string by commands.
	StrFormatter struct{}
)

// Format formats a string by given commands.
// Write a given string to a temporary file and execute commands to the file and read the file.
func (formatter StrFormatter) Format(dest, text string, formatters []string, reader domain.FileReader, executer domain.CmdExecuter) (string, error) {
	// generate tempfile
	tmpfile, err := ioutil.TempFile(filepath.Dir(dest), "cfg.go")
	if err != nil {
		return "", err
	}
	defer os.Remove(tmpfile.Name())
	tmpPath := tmpfile.Name()
	// write tempfile
	if _, err := tmpfile.Write([]byte(text)); err != nil {
		return "", err
	}
	if err := tmpfile.Close(); err != nil {
		return "", err
	}
	// format tempfile
	for _, s := range formatters {
		if err := executer.Exec(s, tmpPath); err != nil {
			return "", err
		}
	}
	return reader.Read(tmpPath)
}
