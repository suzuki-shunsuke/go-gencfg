package infra

import (
	"fmt"
	"path/filepath"

	"github.com/suzuki-shunsuke/go-gencfg/domain"
)

type (
	// CodeGenerater generates code.
	CodeGenerater struct {
		Renderer   domain.TemplateRenderer
		DirMaker   domain.DirMaker
		FileWriter domain.FileWriter
	}
)

// Exec generates code with template and configuration.
func (cg CodeGenerater) Exec(dest, cfgTmpl string, config domain.Cfg) error {
	// create directory
	dir := filepath.Dir(dest)
	fmt.Printf("create %s\n", dir)
	if err := cg.DirMaker.Make(dir); err != nil {
		return err
	}
	txt, err := cg.Renderer.Render(cfgTmpl, config)
	if err != nil {
		return err
	}
	fmt.Printf("create %s\n", dest)
	return cg.FileWriter.Write(dest, []byte(txt))
}
