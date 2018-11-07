package infra

import (
	"bytes"
	"strings"
	"text/template"
)

type (
	// TextTemplateRenderer renders template.
	TextTemplateRenderer struct{}
)

// Render renders a template with given data.
func (txtTmpl TextTemplateRenderer) Render(txt string, data interface{}) (string, error) {
	t := template.New("dummy")
	tmpl, err := t.Parse(strings.Trim(txt, "\n"))
	if err != nil {
		return "", err
	}
	buf := bytes.Buffer{}
	err = tmpl.Execute(&buf, data)
	return buf.String(), err
}
