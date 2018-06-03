package usecase

import (
	"fmt"
	"strings"

	"github.com/suzuki-shunsuke/go-gencfg/domain"
)

// InitGenCfgFile generates the configuration wrapper and test file.
func InitGenCfgFile(args domain.InitGenCfgFileArgs) error {
	dest := args.Dest
	pNames := args.PNames
	exist := args.Exist
	renderer := args.Renderer
	writer := args.Writer

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
