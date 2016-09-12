package main

import "fmt"

type Summary struct {
	Count int
}

// Post structure for handle facebook's Post
type Post struct {
	PostID string
	Name   string
}

func main() {
	setupFBVersion()
	performFetch()
}

func performFetch() {
	fields := "likes.summary(true),name,description"
	urls := []string{"novoyevremya", "www.adme.ru", "afigenchik.ru"}

	posts := make([]Post, 0)

	for _, url := range urls {
		fetcher := Fetcher{url, fields}
		newPosts := fetcher.FetchPosts()
		posts = append(posts, newPosts...)
	}
	for _, post := range posts {
		fmt.Println(post.Name)
	}
}
