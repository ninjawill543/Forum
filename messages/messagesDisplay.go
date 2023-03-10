package forum

import (
	"database/sql"
	"fmt"
	t "forum/listTopics"
	t2 "forum/structs"
	"net/http"
	"strings"
)

var Messages t2.Messages

func MessagesPageDisplay(db *sql.DB, r *http.Request) {
	var username string

	cookie, err := r.Cookie("session")
	if err != nil {
		// fmt.Println(err)
		Messages.SessionUser = ""
	} else {
		queryGetName := fmt.Sprintf("SELECT username FROM users WHERE uuid = '%s'", cookie.Value)
		row3, err := db.Query(queryGetName)
		if err != nil {
			fmt.Println(err)
		} else {
			for row3.Next() {
				defer row3.Close()
				row3.Scan(&username)
			}
			Messages.SessionUser = username
		}
	}

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
	row, err := db.Query(queryTopicName)
	if err != nil {
		fmt.Println(err)
	} else {
		for row.Next() {
			defer row.Close()
			err = row.Scan(&uuid)
			if err != nil {
				fmt.Println(err)
			}
		}
		uuidPath = uuid
	}

	for i := 0; i < len(t.TOPICSANDSESSION.Topics); i++ {
		Messages.CreationDate = t.TOPICSANDSESSION.Topics[i].CreationDate
		Messages.Name = t.TOPICSANDSESSION.Topics[i].Name
		Messages.Owner = t.TOPICSANDSESSION.Topics[i].Owner
		Messages.Likes = t.TOPICSANDSESSION.Topics[i].Likes
		Messages.Id = t.TOPICSANDSESSION.Topics[i].Id
		Messages.UuidPath = t.TOPICSANDSESSION.Topics[i].Uuid
	}
	query := fmt.Sprintf("SELECT id, message, creationDate, owner, report, like, edited, uuid FROM messages WHERE uuidPath = '%s' ORDER BY %s DESC", uuidPath, filter)
	row2, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	} else {
		Messages.Messages = nil
		for row2.Next() {
			defer row2.Close()
			err = row2.Scan(&id, &message, &creationDate, &owner, &report, &like, &edited, &uuid)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(creationDate)

				messageIndex := len(Messages.Messages)

				Messages.Messages = append(Messages.Messages, t2.Message{})
				Messages.Messages[messageIndex].Id = id
				Messages.Messages[messageIndex].Message = message
				Messages.Messages[messageIndex].CreationDate = creationDate
				Messages.Messages[messageIndex].Owner = owner
				Messages.Messages[messageIndex].Report = report
				Messages.Messages[messageIndex].Uuid = uuid
				Messages.Messages[messageIndex].Like = like
				Messages.Messages[messageIndex].Edited = edited
			}
		}
	}
}
