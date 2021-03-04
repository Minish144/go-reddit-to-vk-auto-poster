package main

import (
	"fmt"
	"reddit-to-vk-auto-poster/src/services/vk"
)

func main() {
	vkclient, _ := vk.Initialize()
	srv := vkclient.GetWallUploadServer()
	res, _ := vkclient.UploadPhoto(srv.UploadURL, "temp/1614864280180195000.jpg")
	result := vkclient.SaveWallPhoto(res)
	fmt.Print(result)
}
