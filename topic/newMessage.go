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
			newMessage := `INSERT INTO messages(message, creationDate, owner, report, uuidPath, uuid) VALUES (?, ?, ?, ?, ?, ?)`
			query, err := db.Prepare(newMessage)
			if err != nil {
				log.Fatal(err)
			}

			_, err = query.Exec(message, creationDate, t.USER.Username, 0, uuidPAth[2], uuid)
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println("new Post")
			}
		}
	}
}
