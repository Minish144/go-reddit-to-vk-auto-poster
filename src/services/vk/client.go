package vk

import (
	"os"

	"github.com/go-vk-api/vk"
	"github.com/joho/godotenv"
)

// Client structure
type Client struct {
	Token   string
	OwnerID string
	GroupID string
	Client  *vk.Client
}

// Initialize method for initialization vk client variable
func Initialize() (*Client, error) {
	godotenv.Load()
	token := os.Getenv("VK_TOKEN")
	ownerID := os.Getenv("OWNER_ID")
	groupID := ownerID[1:]

	client := Client{
		Token:   token,
		OwnerID: ownerID,
		GroupID: groupID,
	}

	vkClient, err := vk.NewClientWithOptions(
		vk.WithToken(token),
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
