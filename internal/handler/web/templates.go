package web

import (
	"embed"
	"html/template"
)

//go:embed templates/*.html
//go:embed templates/products/*.html
//go:embed templates/sales/*.html
var templatesFS embed.FS

func LoadTemplates() (*template.Template, error) {
	return template.ParseFS(
		templatesFS,
		"templates/*.html",
		"templates/products/*.html",
		"templates/sales/*.html",
	)
}
