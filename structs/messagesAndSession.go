package forum

type Messages struct {
	Id           int
	Name         string
	Likes        int
	Dislikes     int
	CreationDate string
	Owner        string
	Uuid         string
	UuidPath     string
	SessionUser  string
	Category     string
	Messages     []Message `Message`
}

type Message struct {
	Message        string
	CreationDate   string
	Owner          string
	Report         int
	Uuid           string
	Id             int
	Like           int
	Edited         int
	IsLiked        int
	IsDisliked     int
	IsOwnerOrAdmin int
}
