package forum

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

func AddUsers(db *sql.DB, username string, password string, email string, creationDate time.Time, birthDate string) {
	var testBool bool
	newUsername := username
	newEmail := email
	rows, err := db.Query("SELECT username, email FROM users;")
	if err != nil {
		fmt.Println(err)
	} else {
		for rows.Next() {
			defer rows.Close()
			err = rows.Scan(&username, &email)
			if err != nil {
				fmt.Println(err)
			} else {
				if username == newUsername {
					fmt.Println("username already taken")
					testBool = true
				} else if email == newEmail {
					fmt.Println("email already taken")
					testBool = true
				}
			}
		}
		rows.Close()
		if !testBool {
			usersInfo := `INSERT INTO users(username, password, email, creationDate, birthDate, admin, reports, uuid, ban) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
			uuid := uuid.New()
			fmt.Println(uuid)

			query, err := db.Prepare(usersInfo)
			if err != nil {
				log.Fatal(err)
			}
			_, err = query.Exec(newUsername, password, newEmail, creationDate, birthDate, 0, 0, uuid, 0)
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println("adding new user :", newUsername, "in users")
			}
		}
	}
}
