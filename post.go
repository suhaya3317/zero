package main

type Post struct {
	Text  string `json:"text"`
	Image string `json:"image"`
}

type Posts []Post
