package main

import (
	"fmt"

	fb "github.com/huandu/facebook"
)

// Fetcher for fetch posts from url
type Fetcher struct {
	Path   string
	Fields string
}

// FetchPosts fetch all post for fetcher
func (fetcher Fetcher) FetchPosts() []Post {
	res, error := fb.Get(fetcher.Path+"/posts", fb.Params{
		"fields":       fetcher.Fields,
		"access_token": "EAAEn4f7l8RoBACiOVZBvu7lmgoONOCvAsmQyZAk8DjZBN3jn23NwEtiOpP7fVjGZAYMJc1IaZATlUQsQz6icZArthIcKWXjzsCnZB69RZBowjZC04MfYCtE3tidfV4InWHlYpBLuHejz9OQH5mSxOlSVbnKQYlx5xduEfeYhIN1BoIAZDZD",
	})
	if nil != error {
		fmt.Println(error)
	}
	var items []fb.Result
	res.DecodeField("data", &items)

	var posts []Post

	for _, item := range items {
		var post Post
		item.Decode(&post)
		item.DecodeField("id", &post.PostID)
		posts = append(posts, post)
	}
	return posts
}
