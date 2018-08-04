package main

import (
	"net/http"
	"log"
	"fmt"
	"html"
	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	appengine.Main()
	log.Fatal(http.ListenAndServe(":8080", nil))
}