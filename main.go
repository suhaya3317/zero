package main

import (
	"net/http"

	"google.golang.org/appengine"
)

func init() {
	r := NewRouter()
	http.Handle("/", r)
}

func main() {
	appengine.Main()
}
