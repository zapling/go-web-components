package wcomponent

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"io"
	"strings"
)

var Template *template.Template

//go:embed component-boilerplate.js.tmpl
var ComponentBoilerplate string

type Component struct {
	Selector     string
	HTML         template.HTML
	UseShadowDOM bool
}

func (c *Component) GetJavascript() (template.JS, error) {
	buf := bytes.NewBuffer([]byte{})
	err := c.Execute(buf)
	if err != nil {
		return "", fmt.Errorf("failed to render template: %w", err)
	}
	return template.JS(buf.String()), nil
}

func (c *Component) Execute(wr io.Writer) (err error) {
	if Template == nil {
		Template, err = template.New("WebComponent").Parse(ComponentBoilerplate)
		if err != nil {
			return fmt.Errorf("failed to parse web component js boilerplate template: %w", err)
		}
	}

	return Template.Execute(wr, c)
}

func (c *Component) GetClassName() string {
	words := strings.Split(c.Selector, "-")
	for i := range words {
		words[i] = strings.Title(words[i])
	}
	return strings.Join(words, "")
}
