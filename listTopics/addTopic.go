package forum

import (
	"database/sql"
	"fmt"
	t "forum/users"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func AddTopic(r *http.Request, database *sql.DB) {
	if r.Method == "POST" {
		var name string
		var topicNameTaken bool
		fmt.Println("New POST: (topic) ")
		topicName := r.FormValue("topic_name")
		firstMessage := r.FormValue("firstMessage")

		if len(topicName) < 3 {
			fmt.Println("Not enough char")
		} else if t.USER.Username == "" {
			fmt.Println("you need to be login to post a topic")
		} else {
			query := `SELECT name FROM topics`
			row, err := database.Query(query)
			if err != nil {
				fmt.Println(err)
			} else {
				for row.Next() {
					row.Scan(&name)
					if name == topicName {
						topicNameTaken = true
					}
				}
				row.Close()
			}

			if !topicNameTaken {
				creationDate := time.Now()
				topicInfo := `INSERT INTO topics(name, firstMessage, creationDate, owner, likes, nmbPosts, lastPost, uuid) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
				uuid := uuid.New()
				query, err := database.Prepare(topicInfo)
				if err != nil {
					fmt.Println(err)
				}

				_, err = query.Exec(topicName, firstMessage, creationDate, t.USER.Username, 0, 0, "0", uuid)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("adding new topic :", topicName, "in TOPICS")
					if len(firstMessage) < 10 {
						fmt.Println("not enough char to post the firstmessage")
					} else {
						if firstMessage != "" {
							AddFirstMessageInMessages(firstMessage, creationDate, t.USER.Username, uuid, database)
						}
					}
				}
			} else {
				fmt.Println("topic name already taken")
			}
		}
	}
}
