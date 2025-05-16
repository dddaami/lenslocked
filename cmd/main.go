package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
)

func render(w http.ResponseWriter, r *http.Request, path string, status int, data any) {
	files := []string{
		"./ui/pages/base.gohtml",
		path,
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	render(w, r, "./ui/pages/hello.gohtml", http.StatusOK, struct{ Name string }{Name: "Dami"})
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	render(w, r, "ui/pages/contact.gohtml", http.StatusOK, nil)
}

func faqPage(w http.ResponseWriter, r *http.Request) {
	render(w, r, "ui/pages/faq.gohtml", http.StatusOK, nil)
}

func main() {
	addr := flag.String("addr", ":4000", "Server address")
	flag.Parse()

	http.HandleFunc("/{$}", home)
	http.HandleFunc("/contact", contactPage)
	http.HandleFunc("/faq", faqPage)

	log.Printf("Starting server at %s", *addr)
	http.ListenAndServe(*addr, nil)
}
