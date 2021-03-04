package vk

// WallUploadServer structer which you get from GetWallUploadServer method
type WallUploadServer struct {
	UploadURL string
	AlbumID   int
	UserID    int
}

// GetWallUploadServer returns address of server for photos uploading
func (c *Client) GetWallUploadServer() {

}
