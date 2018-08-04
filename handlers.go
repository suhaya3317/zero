package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func PostIndex(w http.ResponseWriter, r *http.Request) {
	posts := Posts{
		Post{Text: "First Tweet!", Image: "https://example.com/example.jpg"},
		Post{Text: "Second Tweet!", Image: "https://example.com/example.png"},
	}

	if err := json.NewEncoder(w).Encode(posts); err != nil {
		panic(err)
	}
}

func PostShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId := vars["postId"]
	fmt.Fprintln(w, "Post show:", postId)
}