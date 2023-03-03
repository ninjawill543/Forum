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
	Messages     []Message `Message`
}

type Message struct {
	Message      string
	CreationDate string
	Owner        string
	Report       int
	Uuid         string
	Id           int
	Like         int
}

var TOPIC Topic

func TopicPageDisplay(db *sql.DB, r *http.Request) {
	var creationDate string
	var owner string
	var report int
	var uuid string
	var message string
	var id int
	var like int
	uuidPath := strings.Split(r.URL.Path, "/")
	for i := 0; i < len(t.TOPICS); i++ {
		TOPIC.CreationDate = t.TOPICS[i].CreationDate
		TOPIC.Name = t.TOPICS[i].Name
		TOPIC.Owner = t.TOPICS[i].Owner
		TOPIC.Likes = t.TOPICS[i].Likes
		TOPIC.Id = t.TOPICS[i].Id
	}
	query := fmt.Sprintf("SELECT id, message, creationDate, owner, report, like, uuid from messages WHERE uuidPath = '%s'", uuidPath[2])
	row, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	} else {
		TOPIC.Messages = nil
		for row.Next() {
			err = row.Scan(&id, &message, &creationDate, &owner, &report, &like, &uuid)
			fmt.Println(message, like, "HERETEST")
			if err != nil {
				fmt.Println(err)
			} else {
				messageIndex := len(TOPIC.Messages)

				TOPIC.Messages = append(TOPIC.Messages, Message{})
				TOPIC.Messages[messageIndex].Id = id
				TOPIC.Messages[messageIndex].Message = message
				TOPIC.Messages[messageIndex].CreationDate = creationDate
				TOPIC.Messages[messageIndex].Owner = owner
				TOPIC.Messages[messageIndex].Report = report
				TOPIC.Messages[messageIndex].Uuid = uuid
				TOPIC.Messages[messageIndex].Like = like
			}
		}
		row.Close()
	}
}
