package main

import (
	"net/http"
	"fmt"
	"html"
	"github.com/gorilla/mux"
	"google.golang.org/appengine"
)

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/", Index).Methods("GET")
	http.Handle("/", r)
}

func main() {
	appengine.Main()
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}