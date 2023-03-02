package forum

import (
	"fmt"
	t "forum/listTopics"
	"net/http"
	"strings"
)

type Topic struct {
	Id           int
	Name         string
	Likes        int
	Dislikes     int
	CreationDate string
	Owner        string
	Uuid         string
	Messages     []string
}

var TOPIC Topic

func TopicPageDisplay(r *http.Request) {
	fmt.Println(TOPIC.Messages)
	uuid := strings.Split(r.URL.Path, "/")
	for i := 0; i < len(t.TOPICS); i++ {
		if t.TOPICS[i].Uuid == uuid[2] {
			TOPIC.CreationDate = t.TOPICS[i].CreationDate
			TOPIC.Name = t.TOPICS[i].Name
			TOPIC.Owner = t.TOPICS[i].Owner
			TOPIC.Likes = t.TOPICS[i].Likes
			TOPIC.Dislikes = t.TOPICS[i].Dislikes
			TOPIC.Id = t.TOPICS[i].Id
		}
	}
}
