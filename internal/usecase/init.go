package usecase

import (
	"fmt"
	"strings"

	"github.com/suzuki-shunsuke/go-gencfg/internal/domain"
)

// InitGenCfgFile generates the configuration wrapper and test file.
func InitGenCfgFile(args domain.InitGenCfgFileArgs) error {
	dest := args.Dest
	tmplDest := args.TmplDest
	pNames := args.PNames
	exist := args.Exist
	renderer := args.Renderer
	writer := args.Writer

	if !exist.Exist(dest) {
		txt, err := renderer.Render(strings.Trim(GenCfgTmpl, "\n"), pNames)
		if err != nil {
			return err
		}
		fmt.Printf("create %s\n", dest)
		if err := writer.Write(dest, []byte(txt)); err != nil {
			return err
		}
	}
	if !exist.Exist(tmplDest) {
		fmt.Printf("create %s\n", tmplDest)
		return writer.Write(tmplDest, []byte(strings.Trim(DefaultCfgTmpl, "\n")))
	}
	return nil
}
