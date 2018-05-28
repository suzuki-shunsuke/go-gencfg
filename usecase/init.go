package usecase

import (
	"fmt"
	"strings"

	"github.com/suzuki-shunsuke/go-gencfg/domain"
)

// InitGenCfgFile generates the configuration wrapper and test file.
func InitGenCfgFile(dest string, pNames []string, renderer domain.TemplateRenderer, writer domain.FileWriter, exist domain.FileExist) error {
	if exist.Exist(dest) {
		return nil
	}
	txt, err := renderer.Render(strings.Trim(GenCfgTmpl, "\n"), pNames)
	if err != nil {
		return err
	}
	fmt.Printf("create %s\n", dest)
	return writer.Write(dest, []byte(txt))
}
