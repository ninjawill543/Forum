package forum

import (
	"database/sql"
	"fmt"
	t "forum/users"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

func AddTopic(r *http.Request, database *sql.DB) {
	if r.Method == "POST" {
		var name string
		var topicNameTaken bool
		category := strings.Split(r.URL.Path, "/")
		category = strings.Split(category[2], "=")

		fmt.Println("New POST: (topic) ")
		topicName := r.FormValue("topic_name")
		firstMessage := r.FormValue("firstMessage")

		if len(topicName) < 4 {
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
					defer row.Close()
					row.Scan(&name)
					if name == topicName {
						topicNameTaken = true
					}
				}
			}

			if !topicNameTaken {
				creationDate := time.Now()
				topicInfo := `INSERT INTO topics(name, firstMessage, creationDate, owner, likes, nmbPosts, lastPost, category, uuid) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
				uuid := uuid.New()
				query, err := database.Prepare(topicInfo)
				if err != nil {
					fmt.Println(err)
				}

				_, err = query.Exec(topicName, firstMessage, creationDate, t.USER.Username, 0, 0, "0", category[1], uuid)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("adding new topic :", topicName, "in TOPICS")
					if len(firstMessage) < 2 {
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
