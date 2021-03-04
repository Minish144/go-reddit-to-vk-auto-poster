package main

import (
	"fmt"
	"reddit-to-vk-auto-poster/src/services/reddit"
	"reddit-to-vk-auto-poster/src/services/vk"
)

func main() {
	redditclient := reddit.Initialize()
	posts, _ := redditclient.GetPosts()

	post := posts[3]

	vkclient, _ := vk.Initialize()
	err := vkclient.PostRedditSubmission(post)
	if err != nil {
		fmt.Print(err.Error())
	}
}
