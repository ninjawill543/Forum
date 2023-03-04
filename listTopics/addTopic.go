package forum

import (
	"database/sql"
	"fmt"
	t "forum/users"
	"log"
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

		if len(topicName) < 5 {
			fmt.Println("Not enough char")
		} else if t.USER.Username == "" {
			fmt.Println("you need to be login to post a topic")
		} else {
			queryTest := `SELECT name FROM topics`
			row, err := database.Query(queryTest)
			if err != nil {
				fmt.Println(err)
			} else {
				for row.Next() {
					row.Scan(&name)
					if name == topicName {
						topicNameTaken = true
					}
				}
			}

			if !topicNameTaken {
				creationDate := time.Now()
				topicInfo := `INSERT INTO topics(name, creationDate, owner, likes, nmbPosts, lastPost, uuid) VALUES (?, ?, ?, ?, ?, ?, ?)`
				uuid := uuid.New()
				query, err := database.Prepare(topicInfo)
				if err != nil {
					log.Fatal(err)
				}

				_, err = query.Exec(topicName, creationDate, t.USER.Username, 0, 0, "0", uuid)
				if err != nil {
					log.Fatal(err)
				} else {
					fmt.Println("adding new topic :", topicName, "in TOPICS")
				}
			} else {
				fmt.Println("topic name already taken")
			}
		}
	}
}