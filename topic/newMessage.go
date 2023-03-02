package forum

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func NewMessage(db *sql.DB, r *http.Request) {
	uuid := strings.Split(r.URL.Path, "/")
	if r.Method == "POST" {
		message := r.FormValue("input_newMessage")
		fmt.Println(message)

		if len(message) > 10 {
			creationDate := time.Now()
			newMessage := `INSERT INTO messages(message, creationDate, owner, report, uuid) VALUES (?, ?, ?, ?, ?)`
			query, err := db.Prepare(newMessage)
			if err != nil {
				fmt.Println("test1")
				log.Fatal(err)
			}

			_, err = query.Exec(message, creationDate, "owner", 0, uuid[2])
			if err != nil {
				fmt.Println("test2")
				log.Fatal(err)
			} else {
				fmt.Println("new Post")
			}
		} else {
			fmt.Println("not enough char")
		}
	}
}
