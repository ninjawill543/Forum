package forum

type TopicsAndSession struct {
	SessionUser string
	Category    string
	Topics      []Topic `Topic`
}

type Topic struct {
	Id             int
	Name           string
	Likes          int
	CreationDate   string
	Owner          string
	Uuid           string
	FirstMessage   string
	NmbPosts       int
	LastPost       string
	IsLiked        int
	IsDisliked     int
	Category       string
	IsOwnerOrAdmin int
}
