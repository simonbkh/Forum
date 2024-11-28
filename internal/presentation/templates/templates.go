package templates

import (
	"html/template"
	"path/filepath"

	"forum/internal/logic/utils"
)

var (
	HomeTemplate     *template.Template
	LoginTemplate    *template.Template
	RegisterTemplate *template.Template
	PostTemplate     *template.Template
)

func ParseFiles() error {
	var err error

	basePath := "../internal/presentation/templates"

	templates, err := template.ParseFiles(
		filepath.Join(basePath, "layouts/index.html"),
		filepath.Join(basePath, "layouts/nav_bar.html"),
		filepath.Join(basePath, "auth/login.html"),
		filepath.Join(basePath, "auth/register.html"),
		filepath.Join(basePath, "layouts/log_nav.html"),
		filepath.Join(basePath, "post/posts.html"),
	)

	if utils.IsErrors(err) {
		return err
	}

	// ربط القوالب بالمتغيرات
	HomeTemplate = templates.Lookup("index.html")
	LoginTemplate = templates.Lookup("login.html")
	RegisterTemplate = templates.Lookup("register.html")
	PostTemplate = templates.Lookup("posts.html")

	return nil
}
