package usecase

import (
	"fmt"
	"strings"

	"github.com/andreyvit/diff"

	"github.com/suzuki-shunsuke/go-gencfg/domain"
)

// Compare generates the configuration wrapper and test file.
func Compare(args domain.CompareArgs) error {
	// read the config from file
	// parse the config
	cfgUC := args.CfgUC
	dest := args.Dest
	tmplPath := args.TmplPath
	renderer := args.Renderer
	executer := args.Executer
	cfgReader := args.CfgReader
	reader := args.Reader
	isQuiet := args.IsQuiet
	isFailure := args.IsFailure
	strFormatter := args.StrFormatter

	config, err := cfgReader.Read(args.Src)
	if err != nil {
		return err
	}
	cfgUC.Update(&config)
	td := domain.TemplateData{
		Cfg:     config,
		CfgUC:   cfgUC,
		EnvUC:   args.EnvUC,
		FlagUC:  args.FlagUC,
		ParamUC: args.ParamUC}
	if dest == "" {
		if config.Dest != "" {
			dest = config.Dest
		} else {
			dest = "cfg/cfg.go"
		}
	}
	cfgTmpl := DefaultCfgTmpl
	if tmplPath != "" {
		cfgTmpl, err = reader.Read(tmplPath)
		if err != nil {
			return err
		}
	} else {
		if config.TemplatePath != "" {
			cfgTmpl, err = reader.Read(config.TemplatePath)
			if err != nil {
				return err
			}
		}
	}
	if err := validateDest(dest); err != nil {
		return err
	}
	// render the config
	exp, err := renderer.Render(strings.Trim(cfgTmpl, "\n"), td)
	if err != nil {
		return err
	}
	act, err := reader.Read(dest)
	if err != nil {
		return err
	}
	// format exp if formatters exist
	if len(config.Formatters) == 0 {
		// read generated code
		// compare
		if exp == act {
			return nil
		}
		if isQuiet {
			return fmt.Errorf("")
		}
		fmt.Println(diff.LineDiff(exp, act))
		if isFailure {
			return fmt.Errorf("")
		}
		return nil
	}

	exp, err = strFormatter.Format(
		dest, exp, config.Formatters, reader, executer)
	if err != nil {
		return err
	}

	// compare
	if exp == act {
		return nil
	}

	if isQuiet {
		return fmt.Errorf("")
	}
	fmt.Println(diff.LineDiff(exp, act))
	if isFailure {
		return fmt.Errorf("")
	}
	return nil
}
