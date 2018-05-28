package usecase

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"

	"github.com/suzuki-shunsuke/go-gencfg/domain"
)

// ValidateDest validates the generated file's extension is ".go".
func ValidateDest(dest string) error {
	ext := filepath.Ext(dest)
	if ext != ".go" {
		return fmt.Errorf("dest's extension is not '.go': %s", ext)
	}
	return nil
}

// GenCfgFile generates the configuration wrapper and test file.
func GenCfgFile(src, dest, tmplPath, testTmplPath string, reader domain.FileReader, cfgReader domain.CfgReader, generater domain.CodeGenerater, executer domain.CmdExecuter) error {
	// read the config from file
	// parse the config
	// generate the config file
	// generate the config test file
	config, err := cfgReader.Read(src)
	if err != nil {
		return errors.Wrap(err, "failed to read configuration file")
	}
	config.Update()
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
	if err := ValidateDest(dest); err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to validate dest: %s", dest))
	}
	if err := generater.Exec(dest, strings.Trim(cfgTmpl, "\n"), config); err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to generate code: %s", dest))
	}
	if len(config.Formatters) != 0 {
		for _, formatter := range config.Formatters {
			if err := executer.Exec(formatter, dest); err != nil {
				return errors.Wrap(
					err, fmt.Sprintf(
						"failed to format code by command: %s %s", formatter, dest))
			}
		}
	}
	testPath := fmt.Sprintf("%s_test.go", dest[:len(dest)-3])
	if err := generater.Exec(testPath, strings.Trim(cfgTestTmpl, "\n"), config); err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to generate code: %s", testPath))
	}
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
