package templates

import (
	"html/template"

	"forum/internal/logic/utils"
)

var (
	HomeTemplate     *template.Template
	LoginTemplate    *template.Template
	RegisterTemplate *template.Template
)

func ParseFiles() error {
	var err error
	templates, err := template.ParseFiles(
		"../internal/presentation/templates/layouts/index.html",
		 "../internal/presentation/templates/layouts/nav_bar.html",
		"../internal/presentation/templates/auth/login.html",
		"../internal/presentation/templates/auth/register.html")
	if utils.IsErrors(err) {
		return err
	}
	HomeTemplate = templates.Lookup("index.html")
	LoginTemplate = templates.Lookup("login.html")
	RegisterTemplate = templates.Lookup("register.html")

	return nil
}
