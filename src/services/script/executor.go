package script

import (
	"fmt"
	"os"
	"reddit-to-vk-auto-poster/src/services/reddit"
	"reddit-to-vk-auto-poster/src/services/vk"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Exec function executes script
func Exec() {
	godotenv.Load()
	redditclient := reddit.Initialize()
	vkclient, _ := vk.Initialize()
	ticker := time.NewTicker(24 * time.Hour)
	osFreq, _ := strconv.Atoi(os.Getenv("FREQUENCY"))
	freq := int(float64(24) / float64(osFreq))
	post(redditclient, vkclient, freq)
	for range ticker.C {
		go func() {
			post(redditclient, vkclient, freq)
		}()
	}
}

func post(redditClient *reddit.Client, vkClient *vk.Client, freq int) {
	posts, _ := redditClient.GetPosts()
	now := int(time.Now().Unix()) + 60
	for index, post := range posts {
		fmt.Print(freq*3600, now, post, "\n\n")
		err := vkClient.PostRedditSubmission(post, now+freq*index*3600)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
}
