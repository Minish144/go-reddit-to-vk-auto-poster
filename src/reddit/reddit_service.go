package redditservice

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
