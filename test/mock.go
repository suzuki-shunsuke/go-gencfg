package test

import (
	"fmt"

	"github.com/suzuki-shunsuke/go-gencfg/domain"
)

type (
	// NoopFileExist implements domain.FileExist .
	NoopFileExist struct{ Data bool }
	// NoopFileReader implements domain.FileReader .
	NoopFileReader struct {
		Content string
		Err     error
	}
	// NoopFileWriter implements domain.FileWriter .
	NoopFileWriter struct {
		Err error
	}
	// NoopDirMaker implements domain.DirMaker .
	NoopDirMaker struct{}
	// ErrDirMaker implements domain.DirMaker .
	ErrDirMaker struct{}
	// NoopTemplateRenderer implements domain.TemplateRenderer .
	NoopTemplateRenderer struct{}
	// ErrTemplateRenderer implements domain.TemplateRenderer .
	ErrTemplateRenderer struct{}
	// NoopCodeGenerater implements domain.CodeGenerater .
	NoopCodeGenerater struct {
		Err error
	}
	// NoopCmdExecuter implements domain.CmdExecuter .
	NoopCmdExecuter struct{}
	// NoopCfgReader implements domain.CfgReader .
	NoopCfgReader struct {
		Cfg domain.Cfg
		Err error
	}

	// NoopStrFormatter implements domain.StrFormatter .
	NoopStrFormatter struct {
		Str string
		Err error
	}
)

// Exec implements domain.CodeGenerater interface.
func (cg NoopCodeGenerater) Exec(dest, tmpl string, config domain.TemplateData) error {
	return cg.Err
}

// Read implements domain.CfgReader interface.
func (reader NoopCfgReader) Read(src string) (domain.Cfg, error) {
	return reader.Cfg, reader.Err
}

// Exist implements domain.FileExist interface.
func (exist NoopFileExist) Exist(src string) bool {
	return exist.Data
}

// Read implements domain.FileReader interface.
func (reader NoopFileReader) Read(src string) (string, error) {
	return reader.Content, reader.Err
}

// Write implements domain.FileWriter interface.
func (writer NoopFileWriter) Write(dest string, data []byte) error {
	return writer.Err
}

// Make implements domain.DirMaker interface.
func (maker NoopDirMaker) Make(dest string) error {
	return nil
}

// Make implements domain.DirMaker interface.
func (maker ErrDirMaker) Make(dest string) error {
	return fmt.Errorf("failed to mkdir: %s", dest)
}

// Exec implements domain.CmdExecuter interface.
func (executer NoopCmdExecuter) Exec(cmd, dest string) error {
	return nil
}

// Render implements domain.TemplateRenderer interface.
func (render NoopTemplateRenderer) Render(txt string, data interface{}) (string, error) {
	return "", nil
}

// Render implements domain.TemplateRenderer interface.
func (render ErrTemplateRenderer) Render(txt string, data interface{}) (string, error) {
	return "", fmt.Errorf("failed to render")
}

// Format implements domain.StrFormatter interface.
func (fmtr NoopStrFormatter) Format(dest, text string, formatters []string, reader domain.FileReader, executer domain.CmdExecuter) (string, error) {
	return fmtr.Str, fmtr.Err
}
