package usecase

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"

	"github.com/suzuki-shunsuke/go-gencfg/domain"
)

func validateDest(dest string) error {
	ext := filepath.Ext(dest)
	if ext != ".go" {
		return fmt.Errorf("dest's extension is not '.go': %s", ext)
	}
	return nil
}

// ValidateGenCfgFileArgs validates domain.GenCfgFileArgs .
func ValidateGenCfgFileArgs(args domain.GenCfgFileArgs) error {
	if args.Reader == nil {
		return fmt.Errorf("GenCfgFileArgs.Reader is required")
	}
	if args.CfgReader == nil {
		return fmt.Errorf("GenCfgFileArgs.CfgReader is required")
	}
	if args.Generater == nil {
		return fmt.Errorf("GenCfgFileArgs.Generater is required")
	}
	if args.Executer == nil {
		return fmt.Errorf("GenCfgFileArgs.Executer is required")
	}
	if args.CfgUC == nil {
		return fmt.Errorf("GenCfgFileArgs.CfgUC is required")
	}
	if args.EnvUC == nil {
		return fmt.Errorf("GenCfgFileArgs.EnvUC is required")
	}
	if args.FlagUC == nil {
		return fmt.Errorf("GenCfgFileArgs.FlagUC is required")
	}
	if args.ParamUC == nil {
		return fmt.Errorf("GenCfgFileArgs.ParamUC is required")
	}
	return nil
}

// GenCfgFile generates the configuration wrapper and test file.
func GenCfgFile(args domain.GenCfgFileArgs) error {
	if err := ValidateGenCfgFileArgs(args); err != nil {
		return errors.Wrap(err, "failed to validate arguments")
	}
	reader := args.Reader
	cfgUC := args.CfgUC
	generater := args.Generater
	executer := args.Executer
	dest := args.Dest
	tmplPath := args.TmplPath
	testTmplPath := args.TestTmplPath

	// read and parse the config from file
	config, err := args.CfgReader.Read(args.Src)
	if err != nil {
		return errors.Wrap(err, "failed to read configuration file")
	}
	cfgUC.Update(&config)
	td := domain.TemplateData{
		Cfg:     config,
		CfgUC:   cfgUC,
		EnvUC:   args.EnvUC,
		FlagUC:  args.FlagUC,
		ParamUC: args.ParamUC}

	// set config.Dest
	if dest == "" {
		if config.Dest != "" {
			dest = config.Dest
		} else {
			dest = "cfg/cfg.go"
		}
	}

	if err := validateDest(dest); err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to validate dest: %s", dest))
	}

	// set code template
	cfgTmpl := DefaultCfgTmpl
	if tmplPath != "" {
		cfgTmpl, err = reader.Read(tmplPath)
		if err != nil {
			return errors.Wrap(
				err, fmt.Sprintf("failed to read the template file: %s", tmplPath))
		}
	} else {
		if config.TemplatePath != "" {
			cfgTmpl, err = reader.Read(config.TemplatePath)
			if err != nil {
				return errors.Wrap(
					err, fmt.Sprintf(
						"failed to read the template file: %s", config.TemplatePath))
			}
		}
	}

	// set test code template
	cfgTestTmpl := DefaultCfgTestTmpl
	if testTmplPath != "" {
		cfgTestTmpl, err = reader.Read(testTmplPath)
		if err != nil {
			return errors.Wrap(
				err, fmt.Sprintf("failed to read the template file: %s", testTmplPath))
		}
	} else {
		if config.TestTemplatePath != "" {
			cfgTestTmpl, err = reader.Read(config.TestTemplatePath)
			if err != nil {
				return errors.Wrap(
					err, fmt.Sprintf(
						"failed to read the template file: %s", config.TestTemplatePath))
			}
		}
	}

	// generate code file
	if err := generater.Exec(dest, strings.Trim(cfgTmpl, "\n"), td); err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to generate code: %s", dest))
	}

	// format code file by formatters
	if len(config.Formatters) != 0 {
		for _, formatter := range config.Formatters {
			if err := executer.Exec(formatter, dest); err != nil {
				return errors.Wrap(
					err, fmt.Sprintf(
						"failed to format code by command: %s %s", formatter, dest))
			}
		}
	}

	// generate test code file
	testPath := fmt.Sprintf("%s_test.go", dest[:len(dest)-3])
	if err := generater.Exec(testPath, strings.Trim(cfgTestTmpl, "\n"), td); err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to generate code: %s", testPath))
	}

	// format test code file by formatters
	if len(config.Formatters) != 0 {
		for _, formatter := range config.Formatters {
			if err := executer.Exec(formatter, testPath); err != nil {
				return errors.Wrap(
					err, fmt.Sprintf(
						"failed to format code by command: %s %s", formatter, testPath))
			}
		}
	}
	return nil
}
