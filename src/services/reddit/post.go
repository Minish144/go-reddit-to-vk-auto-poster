package reddit

import "github.com/jzelinskie/geddit"

// Post structure for reddit submission
type Post struct {
	Title     string
	ImageURL  string
	ImagePath string
	HasImage  bool
}

// SubmissiontToPost converts geddit.Submission to Post structure
func SubmissiontToPost(submission *geddit.Submission) *Post {
	imageURL := ""
	hasImage := false
	if submission.URL != "" {
		imageURL = submission.URL
		hasImage = true
	} else {
		hasImage = false
	}
	return &Post{
		Title:    submission.Title,
		ImageURL: imageURL,
		HasImage: hasImage,
	}
}

// SubmissionsToPosts converts array of geddit.Submissions to array of Posts
func SubmissionsToPosts(submissions []*geddit.Submission) []*Post {
	postsArray := []*Post{}
	for _, subm := range submissions {
		postsArray = append(postsArray, SubmissiontToPost(subm))
	}
	return postsArray
}
