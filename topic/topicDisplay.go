package forum

import (
	"database/sql"
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
	Users        []string
}

var TOPIC Topic

func TopicPageDisplay(db *sql.DB, r *http.Request) {
	TOPIC.Messages = nil
	var message string
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
	querry := fmt.Sprintf("SELECT message from messages WHERE uuid = '%s'", uuid[2])
	row, err := db.Query(querry)
	if err != nil {
		fmt.Println(err)
	} else {
		for row.Next() {
			err = row.Scan(&message)
			if err != nil {
				fmt.Println(err)
			} else {
				TOPIC.Messages = append(TOPIC.Messages, message)
			}
		}
		row.Close()
	}
}
