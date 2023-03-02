package forum

import (
	"database/sql"
	"fmt"
	t "forum/users"
	"log"
	"net/http"
	"strings"
	"time"
)

func NewMessage(db *sql.DB, r *http.Request) {
	uuid := strings.Split(r.URL.Path, "/")
	if r.Method == "POST" {
		message := r.FormValue("input_newMessage")

		if len(message) < 10 {
			fmt.Println("not enough char to post a message")
		} else if t.USER.Username == "" {
			fmt.Println("you need to be login to post a message")
		} else {
			creationDate := time.Now()
			newMessage := `INSERT INTO messages(message, creationDate, owner, report, uuid) VALUES (?, ?, ?, ?, ?)`
			query, err := db.Prepare(newMessage)
			if err != nil {
				log.Fatal(err)
			}

			_, err = query.Exec(message, creationDate, t.USER.Username, 0, uuid[2])
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println("new Post")
			}
		}
	}
}
