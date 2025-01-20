package templates

import (
	"html/template"

	"forum/internal/logic/utils"
)

var (
	HomeTemplate     *template.Template
	LoginTemplate    *template.Template
	// RegisterTemplate *template.Template
	Create_post      *template.Template
	ErrorTemplate    *template.Template
)

func ParseFiles() error {
	var err error
	templates, err := template.ParseFiles(
		"../internal/presentation/templates/layouts/index.html",
		"../internal/presentation/templates/layouts/nav_bar.html",
		"../internal/presentation/templates/auth/login.html",
		// "../internal/presentation/templates/auth/register.html",
		"../internal/presentation/templates/post/create_post.html",
		"../internal/presentation/templates/errors/error.html",
	)
	if utils.IsErrors(err) {
		return err
	}
	HomeTemplate = templates.Lookup("index.html")
	LoginTemplate = templates.Lookup("login.html")
	// RegisterTemplate = templates.Lookup("register.html")
	Create_post = templates.Lookup("create_post.html")
	ErrorTemplate = templates.Lookup("error.html")
	return nil
}
