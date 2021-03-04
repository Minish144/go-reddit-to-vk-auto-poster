package main

import (
	"fmt"
	"reddit-to-vk-auto-poster/src/services/vk"
)

func main() {
	vkclient, _ := vk.Initialize()

	server := vkclient.GetWallUploadServer()
	photoUploadResult, _ := vkclient.UploadPhoto(server.UploadURL, "temp/1614864280180195000.jpg")
	attachment, _ := vkclient.SaveWallPhoto(photoUploadResult)

	fmt.Print(attachment)
}
