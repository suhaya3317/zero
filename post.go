package main

type Post struct {
	Id int64 `json:"id"`
	Text  string `json:"text"`
	Image string `json:"image"`
	UserId int64 `json:"user_id"`
}

type Posts []Post
