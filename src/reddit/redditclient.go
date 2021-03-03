package reddit

import (
	"os"
	"strconv"

	"github.com/jzelinskie/geddit"
)

// Client structure
type Client struct {
	clientID     string
	clientSecret string
	password     string
	userAgent    string
	username     string
	subreddit    string
	limit        int
	frequency    int
	session      *geddit.LoginSession
}

// Setting new session for client
func (c *Client) setConfig() {
	session, _ := geddit.NewLoginSession(
		c.username,
		c.password,
		c.userAgent,
	)
	c.session = session
}

// Initialize gives you
// client that allows you
// to get posts and other
// stuff from reddit
func Initialize() *Client {
	limit, _ := strconv.Atoi(os.Getenv("LIMIT"))
	freq, _ := strconv.Atoi(os.Getenv("FREQUENCY"))
	client := &Client{
		clientID:     os.Getenv("CLIENTID"),
		clientSecret: os.Getenv("CLIENTSECRET"),
		password:     os.Getenv("PASSWORD"),
		userAgent:    os.Getenv("USERAGENT"),
		username:     os.Getenv("USERNAME"),
		subreddit:    os.Getenv("SUBREDDIT"),
		limit:        limit,
		frequency:    freq,
	}
	client.setConfig()
	return client
}

// GetPosts returns
func (c *Client) GetPosts() ([]*geddit.Submission, error) {
	posts, err := c.session.SubredditSubmissions(c.subreddit, "hot", geddit.ListingOptions{Limit: c.limit})
	return posts, err
}
