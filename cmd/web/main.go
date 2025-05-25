package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/ddddami/lenslocked/ui"
	"github.com/go-chi/chi/v5"
)

type application struct {
	templateCache map[string]*template.Template
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.User = User{Name: "Dami"}
	app.render(w, r, "home.gohtml", http.StatusOK, data)
}

func (app *application) contactPage(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	app.render(w, r, "contact.gohtml", http.StatusOK, data)
}

func (app *application) faqPage(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	app.render(w, r, "faq.gohtml", http.StatusOK, data)
}

func (app *application) signup(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	app.render(w, r, "signup.gohtml", http.StatusOK, data)
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	app.render(w, r, "login.gohtml", http.StatusOK, data)
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

	r.Handle("GET /static/*", http.FileServerFS(ui.Files))

	r.Get("/", app.home)
	r.Get("/contact", app.contactPage)
	r.Get("/faq", app.faqPage)

	r.Get("/login", app.login)
	r.Get("/signup", app.signup)

	log.Printf("Starting server at %s", *addr)
	http.ListenAndServe(*addr, r)
}
