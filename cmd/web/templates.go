package main

import (
	"html/template"
	"io/fs"
	"path/filepath"

	"github.com/ddddami/lenslocked/ui"
)

type Template struct{}

type User struct {
	Name string
}

type templateData struct {
	CurrentYear int
	User        struct {
		Name string
	}
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/*.gohtml")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFS(
			ui.Files,
			"html/base.gohtml",
			page,
		)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}
	return cache, nil
}

var functions = template.FuncMap{}
