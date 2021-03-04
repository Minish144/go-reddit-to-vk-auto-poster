package vk

import (
	"fmt"
	"os"

	"github.com/go-vk-api/vk"
)

// Initialize method for initialization vk client variable
func Initialize() (*vk.Client, error) {
	fmt.Print("Token: ", os.Getenv("VK_TOKEN"))
	return vk.NewClientWithOptions(
		vk.WithToken(os.Getenv("VK_TOKEN")),
	)
}
