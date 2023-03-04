package forum

import (
	"database/sql"
	"fmt"
	t "forum/users"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

func NewMessage(db *sql.DB, r *http.Request) {
	topicName := strings.Split(r.URL.Path, "/")
	uuid := uuid.New()

	if r.Method == "POST" {
		databaseTopics, _ := sql.Open("sqlite3", "../topics.db")

		query := fmt.Sprintf("SELECT uuid FROM topics WHERE name = '%s'", topicName[2])
		row, err := databaseTopics.Query(query)
		if err != nil {
			fmt.Println(err)
		} else {
			for row.Next() {
				err = row.Scan(&uuid)
				if err != nil {
					fmt.Println(err)
				}
			}
			row.Close()
		}
		uuidPath := uuid

		message := r.FormValue("input_newMessage")

		if len(message) < 10 {
			fmt.Println("not enough char to post a message")
		} else if t.USER.Username == "" {
			fmt.Println("you need to be login to post a message")
		} else {
			creationDate := time.Now()

			newMessageQuery := `INSERT INTO messages(message, creationDate, owner, report, uuidPath, like, edited, uuid) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
			queryMessage, err := db.Prepare(newMessageQuery)
			if err != nil {
				fmt.Println(err)
			}

			_, err = queryMessage.Exec(message, creationDate, t.USER.Username, 0, uuidPath, 0, 0, uuid)
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println("new message")
				query2 := fmt.Sprintf("UPDATE topics SET nmbPosts = nmbPosts + 1, lastPost = '%s' WHERE uuid = '%s'", creationDate, uuidPath)
				databaseTopics.Exec(query2)
			}
		}
	}
}
