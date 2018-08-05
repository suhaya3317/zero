package main

type Post struct {
	Id int64 `json:"id"`
	Text  string `json:"text"`
	Image string `json:"image"`
}

type Posts []Post
