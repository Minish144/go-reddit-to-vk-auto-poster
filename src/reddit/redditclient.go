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

// Getting session variable which
// gives you an access to reddit
func initialize() *Client {
	limit, _ := strconv.Atoi(os.Getenv("LIMIT"))
	client := &Client{
		clientID:     os.Getenv("CLIENTID"),
		clientSecret: os.Getenv("CLIENTID"),
		password:     os.Getenv("PASSWORD"),
		userAgent:    os.Getenv("USERAGENT"),
		username:     os.Getenv("USERNAME"),
		subreddit:    os.Getenv("SUBREDDIT"),
		limit:        limit,
	}
	client.setConfig()
	return client
}
