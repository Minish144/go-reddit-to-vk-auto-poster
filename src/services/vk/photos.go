package vk

import "github.com/go-vk-api/vk"

// WallUploadServer structer which you get from GetWallUploadServer method
type WallUploadServer struct {
	UploadURL string `json:"upload_url"`
	AlbumID   int    `json:"album_id"`
	UserID    int    `json:"user_id"`
}

// GetWallUploadServer returns address of server for photos uploading
func (c *Client) GetWallUploadServer() *WallUploadServer {
	var wallUploadServer WallUploadServer
	params := vk.RequestParams{
		"group_id": c.GroupID,
	}
	c.Client.CallMethod(
		"photos.getWallUploadServer",
		params,
		&wallUploadServer,
	)
	return &wallUploadServer
}
