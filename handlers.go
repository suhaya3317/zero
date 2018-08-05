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

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		panic(err)
	}
}

func PostShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId := vars["postId"]
	fmt.Fprintln(w, "Post show:", postId)
}

func UserIndex(w http.ResponseWriter, r *http.Request) {
	users := Users{
		User{Name: "Hayato Suzuki", Description: "「すはや」です。RubyとGoを書きます。"},
		User{Name: "Hinako Sano", Description: "「最高のひなこ」発売中！"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		panic(err)
	}
}

func UserShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	fmt.Fprintln(w, "User show:", userId)
}