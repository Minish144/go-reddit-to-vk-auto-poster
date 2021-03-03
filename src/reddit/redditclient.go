package redditservice

import (
	"os"
	"strconv"
)

// RedditClient structure
type RedditClient struct {
	clientID     string
	clientSecret string
	password     string
	userAgent    string
	username     string
	subreddit    string
	limit        int
}

func getRedditClient() *RedditClient {
	clientID := os.Getenv("CLIENTID")
	clientSecret := os.Getenv("CLIENTID")
	password := os.Getenv("PASSWORD")
	userAgent := os.Getenv("USERAGENT")
	username := os.Getenv("USERNAME")
	subreddit := os.Getenv("SUBREDDIT")
	limit, _ := strconv.Atoi(os.Getenv("LIMIT"))

	return &RedditClient{
		clientID:     clientID,
		clientSecret: clientSecret,
		password:     password,
		userAgent:    userAgent,
		username:     username,
		subreddit:    subreddit,
		limit:        limit,
	}
}
