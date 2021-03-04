package vk

// WallPost structure for simple group post
type WallPost struct {
	OwnerID     string
	FromGroup   int
	Message     string
	Attachments []string
	PublishDate int64
}
