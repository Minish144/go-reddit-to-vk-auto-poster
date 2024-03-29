package reddit

import (
	"fmt"
	"os"
	"strings"

	"github.com/jzelinskie/geddit"
)

// Post structure for reddit submission
type Post struct {
	Title     string
	ImageURL  string
	ImagePath string
	HasImage  bool
	Source    string
}

// SubmissiontToPost converts geddit.Submission to Post structure
func SubmissiontToPost(submission *geddit.Submission) *Post {
	imageURL := ""
	hasImage := false
	if strings.Contains(submission.URL, ".jpg") ||
		strings.Contains(submission.URL, ".jpeg") ||
		strings.Contains(submission.URL, ".png") {
		imageURL = submission.URL
		hasImage = true
	} else {
		hasImage = false
	}

	return &Post{
		Title:    submission.Title,
		ImageURL: imageURL,
		HasImage: hasImage,
		Source:   "https://reddit.com" + submission.Permalink,
	}
}

// SubmissionsToPosts converts array of geddit.Submissions to array of Posts
func SubmissionsToPosts(submissions []*geddit.Submission) []*Post {
	postsArray := []*Post{}
	for _, subm := range submissions {
		fmt.Println(subm.Permalink, subm.LinkFlairText)
		postsArray = append(postsArray, SubmissiontToPost(subm))
	}
	return postsArray
}

//DeletePostsPhotos removes photo from directory of reddit.Post
func DeletePostsPhotos(posts []*Post) {
	for _, elem := range posts {
		if elem.HasImage {
			os.Remove(elem.ImagePath)
			elem.HasImage = false
			elem.ImagePath = ""
		}
	}
}

//DeletePostPhoto removes photo from directory of reddit.Post
func DeletePostPhoto(post *Post) {
	posts := []*Post{post}
	DeletePostsPhotos(posts)
}
