package main

import (
	"net/http"
	"fmt"
	"html"
	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", Index)
	appengine.Main()
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}