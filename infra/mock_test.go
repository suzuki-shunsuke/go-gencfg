package infra_test

import (
	"fmt"
)

type (
	NoopRenderer    struct{}
	NoopDirMaker    struct{}
	NoopFileWriter  struct{}
	NoopCmdExecuter struct{}
	ErrDirMaker     struct{}
	ErrRenderer     struct{}
	NoopFileReader  struct {
		content string
	}
)

func (executer NoopCmdExecuter) Exec(cmd, dest string) error {
	return nil
}

func (reader NoopFileReader) Read(src string) (string, error) {
	return reader.content, nil
}

func (render NoopRenderer) Render(txt string, data interface{}) (string, error) {
	return "", nil
}

func (render ErrRenderer) Render(txt string, data interface{}) (string, error) {
	return "", fmt.Errorf("failed to render")
}

func (maker NoopDirMaker) Make(dest string) error {
	return nil
}

func (maker ErrDirMaker) Make(dest string) error {
	return fmt.Errorf("failed to mkdir: %s", dest)
}

func (writer NoopFileWriter) Write(dest string, data []byte) error {
	return nil
}
