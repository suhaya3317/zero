package main

import (
	"net/http"
	"fmt"
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

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func PostIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Post Index!")
}

func PostShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId := vars["postId"]
	fmt.Fprintln(w, "Post show:", postId)
}