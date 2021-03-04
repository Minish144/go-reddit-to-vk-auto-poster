package main

import (
	"reddit-to-vk-auto-poster/src/services/vk"
)

func main() {
	// structure := reddit.Initialize()
	// posts, _ := structure.GetPosts()
	// for _, post := range posts {
	// 	fmt.Print(post)
	// 	fmt.Print("\n--------\n")
	// }
	vkclient, _ := vk.Initialize()
	srv := vkclient.GetWallUploadServer()
	res, err := vkclient.UploadPhoto(srv.UploadURL, "temp/1614864280180195000.jpg")
}
