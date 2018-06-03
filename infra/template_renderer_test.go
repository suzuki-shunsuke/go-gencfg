package infra_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-gencfg/domain"
	"github.com/suzuki-shunsuke/go-gencfg/infra"
)

func TestTextTemplateRendererRender(t *testing.T) {
	renderer := infra.TextTemplateRenderer{}
	data := domain.Cfg{Dest: "hello"}
	t.Run("positive", func(t *testing.T) {
		act, err := renderer.Render("{{.Dest}}", data)
		if err != nil {
			t.Fatal(err)
		}
		if act != data.Dest {
			t.Fatalf(`Render() returns "%s", wanted "%s"`, act, data.Dest)
		}
	})
	t.Run("positive", func(t *testing.T) {
		act, err := renderer.Render("", nil)
		if err != nil {
			t.Fatal(err)
		}
		if act != "" {
			t.Fatalf(`renderer.Render("", nil) = "%s", wanted ""`, act)
		}
	})
}
