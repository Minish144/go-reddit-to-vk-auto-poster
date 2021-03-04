package main

import (
	"fmt"
	"reddit-to-vk-auto-poster/src/services/vk"
)

func main() {
	// structure := reddit.Initialize()
	// posts, _ := structure.GetPosts()
	// for _, post := range posts {
	// 	fmt.Print(post)
	// 	fmt.Print("\n--------\n")
	// }
	vkClient, _ := vk.Initialize()
	srv := vkClient.GetWallUploadServer()
	fmt.Print(srv)
}
