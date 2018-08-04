package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"
)

type Post struct {
	Text  string `json:"text"`
	Image string `json:"image"`
}

type Posts []Post

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
	posts := Posts{
		Post{Text: "First Tweet!", Image: "https://example.com/example.jpg"},
		Post{Text: "Second Tweet!", Image: "https://example.com/example.png"},
	}

	json.NewEncoder(w).Encode(posts)
}

func PostShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId := vars["postId"]
	fmt.Fprintln(w, "Post show:", postId)
}
