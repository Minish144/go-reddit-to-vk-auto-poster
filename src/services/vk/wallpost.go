package vk

import (
	"reddit-to-vk-auto-poster/src/services/reddit"
	"strconv"

	"github.com/go-vk-api/vk"
)

// WallPost structure for simple group post
type WallPost struct {
	OwnerID     string
	FromGroup   int
	Message     string
	Attachment  string
	PublishDate int64
}

// RedditSubmissionToVkPost convert reddit post struct to vk wallpost struct
func RedditSubmissionToVkPost(
	post *reddit.Post,
	client *Client,
	fromGroup int,
	attachment string,
	publishDate int64) *WallPost {
	wallpost := WallPost{
		OwnerID:     client.OwnerID,
		FromGroup:   fromGroup,
		Message:     post.Title,
		Attachment:  attachment,
		PublishDate: publishDate,
	}
	return &wallpost
}

// Post posts a new post on the wall of your public
func (c *Client) Post(post *WallPost) error {
	params := vk.RequestParams{
		"owner_id":    post.OwnerID,
		"from_group":  post.FromGroup,
		"message":     post.Message,
		"attachments": post.Attachment,
	}
	err := c.Client.CallMethod("wall.post", params, nil)
	if err != nil {
		return err
	}
	return nil
}

// Photo structure
type Photo struct {
	ID      int `json:"id"`
	OwnerID int `json:"owner_id"`
}

// GetWallUploadServer returns address of server for photos uploading
func (c *Client) GetWallUploadServer() (*WallUploadServer, error) {
	var wallUploadServer WallUploadServer
	params := vk.RequestParams{
		"group_id": c.GroupID,
	}
	err := c.Client.CallMethod(
		"photos.getWallUploadServer",
		params,
		&wallUploadServer,
	)
	if err != nil {
		return &WallUploadServer{}, err
	}
	return &wallUploadServer, nil
}

// SaveWallPhoto saves phhoto to VK
func (c *Client) SaveWallPhoto(uploaded *PhotoUploadResult) (string, error) {
	var r []Photo

	params := vk.RequestParams{
		"group_id": c.GroupID,
		"photo":    uploaded.Photo,
		"server":   uploaded.Server,
		"hash":     uploaded.Hash,
	}

	err := c.Client.CallMethod("photos.saveWallPhoto", params, &r)
	if err != nil {
		return "", err
	}
	id := r[0].ID
	photo := "photo" + c.GroupID + "_" + strconv.Itoa(id)
	return photo, nil
}
