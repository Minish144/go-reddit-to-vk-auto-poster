package vk

import (
	"os"

	"github.com/go-vk-api/vk"
)

// Client structure
type Client struct {
	Token   string
	OwnerID string
	Client  *vk.Client
}

// Initialize method for initialization vk client variable
func Initialize() (*Client, error) {
	client := Client{
		Token:   os.Getenv("VK_TOKEN"),
		OwnerID: os.Getenv("OWNER_ID"),
	}
	vkClient, err := vk.NewClientWithOptions(
		vk.WithToken(os.Getenv("VK_TOKEN")),
	)
	client.Client = vkClient
	return &client, err
}

// PostRedditSubmission posts provided Post to VK
func (c *Client) PostRedditSubmission() {

}

// PostRedditSubmissions posts provided Posts to VK
func (c *Client) PostRedditSubmissions(withPhotos bool) {

}
