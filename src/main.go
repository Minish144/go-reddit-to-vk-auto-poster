package main

import (
	"fmt"
	"reddit-to-vk-auto-poster/src/reddit"
)

func main() {
	structure := reddit.Initialize()
	posts, _ := structure.GetPosts()
	for _, post := range posts {
		fmt.Print(post.URL)
		fmt.Print("\n--------\n")
	}
}
