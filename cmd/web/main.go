package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

type application struct {
	templateCache map[string]*template.Template
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "hello.gohtml", http.StatusOK, struct{ Name string }{Name: "Dami"})
}

func (app *application) contactPage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "contact.gohtml", http.StatusOK, nil)
}

func (app *application) faqPage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "faq.gohtml", http.StatusOK, nil)
}

func main() {
	addr := flag.String("addr", ":4000", "Server address")
	flag.Parse()

	templateCache, err := newTemplateCache()
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	app := &application{
		templateCache: templateCache,
	}

	r := chi.NewRouter()
	r.Get("/", app.home)
	r.Get("/contact", app.contactPage)
	r.Get("/faq", app.faqPage)

	log.Printf("Starting server at %s", *addr)
	http.ListenAndServe(*addr, r)
}
