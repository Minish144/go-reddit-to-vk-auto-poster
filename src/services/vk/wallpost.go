package vk

import "reddit-to-vk-auto-poster/src/services/reddit"

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
	publishDate int64) *WallPost {
	wallpost := WallPost{
		OwnerID:     client.OwnerID,
		FromGroup:   fromGroup,
		Message:     post.Title,
		Attachment:  post.ImagePath,
		PublishDate: publishDate,
	}
	return &wallpost
}
