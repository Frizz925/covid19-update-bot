package templates

import (
	"bytes"
	"path"
	"text/template"
)

type Generator struct {
	tmpl *template.Template
}

func NewGenerator(templatesDir string) (*Generator, error) {
	tmpl, err := template.ParseGlob(path.Join(templatesDir, "*.tmpl"))
	if err != nil {
		return nil, err
	}
	return &Generator{tmpl}, nil
}

func (g *Generator) Generate(name string, data any) (string, error) {
	var buf bytes.Buffer
	if err := g.tmpl.ExecuteTemplate(&buf, name, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
