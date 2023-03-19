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
	//adding new message
	uuidMessage := uuid.New()
	var uuid string
	topicName := strings.Split(r.URL.Path, "/")

	if r.Method == "POST" {

		query := fmt.Sprintf("SELECT uuid FROM topics WHERE name = '%s'", topicName[2])
		row, err := db.Query(query)
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
		}
		uuidPath := uuid

		message := r.FormValue("input_newMessage")

		if len(message) < 2 {
			fmt.Println("not enough char to post a message")
		} else if t.USER.Username == "" {
			fmt.Println("you need to be login to post a message")
		} else {
			creationDate := time.Now().String()

			newMessageQuery := `INSERT INTO messages(message, creationDate, owner, report, uuidPath, like, edited, uuid) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
			queryMessage, err := db.Prepare(newMessageQuery)
			if err != nil {
				fmt.Println(err)
			}

			_, err = queryMessage.Exec(message, creationDate, t.USER.Username, 0, uuidPath, 0, 0, uuidMessage)
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println("new message")
				query2 := fmt.Sprintf("UPDATE topics SET nmbPosts = nmbPosts + 1, lastPost = '%s' WHERE uuid = '%s'", creationDate, uuidPath)
				db.Exec(query2)
			}
		}
	}
}
