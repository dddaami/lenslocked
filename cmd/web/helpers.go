package main

import (
	"bytes"
	"log"
	"net/http"
)

func (app *application) render(w http.ResponseWriter, r *http.Request, page string, status int, data any) {
	ts, ok := app.templateCache[page]
	if !ok {
		log.Printf("the template %s does not exist", page)
		return
	}

	buf := new(bytes.Buffer)
	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	w.WriteHeader(status)
	buf.WriteTo(w)
}
