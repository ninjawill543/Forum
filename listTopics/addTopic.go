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
		fmt.Println("New POST: (topic) ")
		topicName := r.FormValue("topic_name")

		if len(topicName) < 5 {
			fmt.Println("Not enough char")
		} else if t.USER.Username == "" {
			fmt.Println("you need to be login to post a topic")
		} else {
			creationDate := time.Now()
			topicInfo := `INSERT INTO topics(name, creationDate, owner, likes, dislikes, uuid) VALUES (?, ?, ?, ?, ?, ?)`
			uuid := uuid.New()
			query, err := database.Prepare(topicInfo)
			if err != nil {
				log.Fatal(err)
			}

			_, err = query.Exec(topicName, creationDate, t.USER.Username, 0, 0, uuid)
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println("adding new topic :", topicName, "in TOPICS")
			}
		}
	}
}
