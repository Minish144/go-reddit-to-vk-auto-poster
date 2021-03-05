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
	PublishDate int
}

// RedditSubmissionToVkPost convert reddit post struct to vk wallpost struct
func RedditSubmissionToVkPost(
	post *reddit.Post,
	client *Client,
	fromGroup int,
	attachment string,
	publishDate int) *WallPost {
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
		"owner_id":     post.OwnerID,
		"from_group":   post.FromGroup,
		"message":      post.Message,
		"attachments":  post.Attachment,
		"publish_date": post.PublishDate,
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
	ownerID := r[0].OwnerID
	photo := "photo" + strconv.Itoa(ownerID) + "_" + strconv.Itoa(id)
	return photo, nil
}

// PostRedditSubmission posts provided Post to VK
func (c *Client) PostRedditSubmission(post *reddit.Post, time int) error {
	if post.HasImage {
		server, err := c.GetWallUploadServer()
		if err != nil {
			return err
		}

		photoUploadResult, err := c.UploadPhoto(server.UploadURL, post.ImagePath)
		if err != nil {
			return err
		}

		attachment, err := c.SaveWallPhoto(photoUploadResult)
		if err != nil {
			return err
		}

		vkPost := RedditSubmissionToVkPost(post, c, 1, attachment, time)
		c.Post(vkPost)

		reddit.DeletePostPhoto(post)
	}
	return nil
}

// PostRedditSubmissions posts provided Posts to VK
func (c *Client) PostRedditSubmissions(withPhotos bool) {

}
