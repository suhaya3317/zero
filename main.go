package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"
)

func init() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", Index).Methods("GET")
	r.HandleFunc("/posts", PostIndex).Methods("GET")
	r.HandleFunc("/posts/{postId}", PostShow)
	http.Handle("/", r)
}

func main() {
	appengine.Main()
}
