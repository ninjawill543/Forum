package forum

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func AddUsers(db *sql.DB, username string, password string, email string, creationDate time.Time, birthDate string) {
	usersInfo := `INSERT INTO users(username, password, email, creationDate, birthDate) VALUES (?, ?, ?, ?, ?)`
	query, err := db.Prepare(usersInfo)
	if err != nil {
		log.Fatal(err)
	}
	_, err = query.Exec(username, password, email, creationDate, birthDate)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("adding new user :", username, "in users")
	}
}
