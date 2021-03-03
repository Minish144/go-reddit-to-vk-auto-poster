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
}

// Setting new session from client
func (c *Client) setCfg() *geddit.LoginSession {
	session, _ := geddit.NewLoginSession(
		c.username,
		c.password,
		c.userAgent,
	)
	return session
}

// Getting session variable which
// gives you an access to reddit
func getSession() (*geddit.LoginSession, *Client) {
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

	return client.setCfg(), client
}
