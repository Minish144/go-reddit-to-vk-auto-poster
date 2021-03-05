package main

import "reddit-to-vk-auto-poster/src/services/script"

func main() {
	// godotenv.Load()
	// redditclient := reddit.Initialize()
	// posts, _ := redditclient.GetPosts()
	// now := time.Now().Unix() + 60
	// vkclient, _ := vk.Initialize()
	// for index, post := range posts {
	// 	err := vkclient.PostRedditSubmission(post, now+int64(index*3600))
	// 	if err != nil {
	// 		fmt.Print(err.Error())
	// 	}
	// }
	script.Exec()
}
