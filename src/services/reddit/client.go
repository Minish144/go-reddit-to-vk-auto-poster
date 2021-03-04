package reddit

import (
	"os"
	"reddit-to-vk-auto-poster/src/utils"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/jzelinskie/geddit"
)

// Client 's structure
type Client struct {
	ClientID     string
	ClientSecret string
	Password     string
	UserAgent    string
	Username     string
	Subreddit    string
	Limit        int
	Frequency    int
	Session      *geddit.LoginSession
}

func downloadImagesForPosts(posts []*Post, dir string) {
	var filepath string

	for _, elem := range posts {
		if elem.HasImage {
			filepath = dir + "/" + strconv.FormatInt(time.Now().UnixNano(), 10) + ".jpg"
			elem.ImagePath = filepath
			utils.DownloadFile(elem.ImagePath, elem.ImageURL)
		}
	}

	DeletePostsPhotos(posts)
}

// Setting new session for client
func (c *Client) setConfig() {
	session, _ := geddit.NewLoginSession(
		c.Username,
		c.Password,
		c.UserAgent,
	)
	c.Session = session
}

// Initialize gives you client that allows you to get posts and other stuff
// from reddit
func Initialize() *Client {
	godotenv.Load()
	limit, _ := strconv.Atoi(os.Getenv("LIMIT"))
	freq, _ := strconv.Atoi(os.Getenv("FREQUENCY"))
	client := &Client{
		ClientID:     os.Getenv("CLIENTID"),
		ClientSecret: os.Getenv("CLIENTSECRET"),
		Password:     os.Getenv("PASSWORD"),
		UserAgent:    os.Getenv("USERAGENT"),
		Username:     os.Getenv("USERNAME"),
		Subreddit:    os.Getenv("SUBREDDIT"),
		Limit:        limit,
		Frequency:    freq,
	}
	client.setConfig()
	return client
}

// GetPosts returns array of Posts
func (c *Client) GetPosts() ([]*Post, error) {
	submissions, err := c.Session.SubredditSubmissions(
		c.Subreddit, "hot",
		geddit.ListingOptions{Limit: c.Limit},
	)
	posts := SubmissionsToPosts(submissions)
	downloadImagesForPosts(posts, "temp")
	return posts, err
}
