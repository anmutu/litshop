package genmodel

import (
	"text/template"
)

func parseRepositoryTemplateOrPanic(t string) *template.Template {
	tpl, err := template.New("repository_template").Parse(t)
	if err != nil {
		panic(err)
	}
	return tpl
}
