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
	UuidPath     string
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
	Edited       int
}

var TOPIC Topic

func MessagesPageDisplay(databaseMessages *sql.DB, databaseTopics *sql.DB, r *http.Request) {
	var creationDate string
	var owner string
	var report int
	var uuid string
	var message string
	var id int
	var like int
	var filter string
	var edited int
	var uuidPath string

	filter = r.FormValue("filter")
	if filter == "" {
		filter = "creationDate"
	}

	topicName := strings.Split(r.URL.Path, "/")
	queryTopicName := fmt.Sprintf("SELECT uuid FROM topics WHERE name = '%s'", topicName[2])
	row, err := databaseTopics.Query(queryTopicName)
	if err != nil {
		fmt.Println(err)
	} else {
		for row.Next() {
			err = row.Scan(&uuid)
			if err != nil {
				fmt.Println(err)
			}
		}
		uuidPath = uuid
	}

	for i := 0; i < len(t.TOPICS); i++ {
		TOPIC.CreationDate = t.TOPICS[i].CreationDate
		TOPIC.Name = t.TOPICS[i].Name
		TOPIC.Owner = t.TOPICS[i].Owner
		TOPIC.Likes = t.TOPICS[i].Likes
		TOPIC.Id = t.TOPICS[i].Id
		TOPIC.UuidPath = t.TOPICS[i].Uuid
	}
	query := fmt.Sprintf("SELECT id, message, creationDate, owner, report, like, edited, uuid FROM messages WHERE uuidPath = '%s' ORDER BY %s DESC", uuidPath, filter)
	row, err = databaseMessages.Query(query)
	if err != nil {
		fmt.Println(err)
	} else {
		TOPIC.Messages = nil
		for row.Next() {
			err = row.Scan(&id, &message, &creationDate, &owner, &report, &like, &edited, &uuid)
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
				TOPIC.Messages[messageIndex].Edited = edited
			}
		}
		row.Close()
	}
}
