package main

import (
	"html/template"
	"path/filepath"
)

type Template struct{}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("ui/html/pages/*.gohtml")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(
			"ui/html/base.gohtml",
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
