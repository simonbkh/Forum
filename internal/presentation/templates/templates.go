package templates

import (
	"html/template"
)
var (
	HomeTemplate  *template.Template
	LoginTemplate *template.Template
	Create_post   *template.Template
	ErrorTemplate *template.Template
	MyPossts      *template.Template
)

func ParseFiles() error {
	var err error

	// Parse all template files, including index.html
	templates, err := template.New("").Funcs(template.FuncMap{
		"add":     add,
		"sub":     sub,
		"iterate": iterate,
	}).ParseFiles(
		"../internal/presentation/templates/layouts/index.html", 
		"../internal/presentation/templates/layouts/nav_bar.html",
		"../internal/presentation/templates/layouts/side_bar.html",
		"../internal/presentation/templates/auth/login.html",
		"../internal/presentation/templates/post/create_post.html",
		"../internal/presentation/templates/errors/error.html",
		"../internal/presentation/templates/post/mypost.html",
	)
	if err != nil {
		return err
	}

	// Assign the parsed templates to the variables
	HomeTemplate = templates.Lookup("index.html")
	LoginTemplate = templates.Lookup("login.html")
	Create_post = templates.Lookup("create_post.html")
	ErrorTemplate = templates.Lookup("error.html")
	MyPossts = templates.Lookup("mypost.html")
	return nil
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func iterate(start, end int) []int {
	var result []int
	for i := start; i <= end; i++ {
		result = append(result, i)
	}
	return result
}

