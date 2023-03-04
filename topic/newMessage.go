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
	uuidPAth := strings.Split(r.URL.Path, "/")
	uuid := uuid.New()

	if r.Method == "POST" {
		message := r.FormValue("input_newMessage")

		if len(message) < 10 {
			fmt.Println("not enough char to post a message")
		} else if t.USER.Username == "" {
			fmt.Println("you need to be login to post a message")
		} else {
			creationDate := time.Now()
			newMessage := `INSERT INTO messages(message, creationDate, owner, report, uuidPath, like, uuid) VALUES (?, ?, ?, ?, ?, ?, ?)`
			query, err := db.Prepare(newMessage)
			if err != nil {
				log.Fatal(err)
			}

			_, err = query.Exec(message, creationDate, t.USER.Username, 0, uuidPAth[2], 0, uuid)
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println("new message")
				databaseTopics, _ := sql.Open("sqlite3", "../topics.db")
				query2 := fmt.Sprintf("UPDATE topics SET nmbPosts = nmbPosts + 1, lastPost = '%s' WHERE uuid = '%s'", creationDate, uuidPAth[2])
				databaseTopics.Exec(query2)
			}
		}
	}
}
