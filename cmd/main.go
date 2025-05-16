package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/pages/hello.gohtml")
	if err != nil {
		fmt.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, struct{ Name string }{Name: "Dami"})
	if err != nil {
		fmt.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>Contact Page</h1> You can reach out at <a href="mailto:dami@damilola.dev">dami@damilola.dev</a>`)
}

func faqPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>FAQ Page</h1> <ul><li><b>Is there a free version? </b>yeah</li></ul>`)
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
