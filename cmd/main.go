package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hiya!")
}

func main() {
	addr := flag.String("addr", ":4000", "Server address")
	flag.Parse()

	http.HandleFunc("/", handlerFunc)

	log.Printf("Starting server at %s", *addr)
	http.ListenAndServe(*addr, nil)
}
