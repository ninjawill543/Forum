package forum

import (
	"database/sql"
	"fmt"
	t "forum/views"
	"net/http"
)

type User struct {
	Username     string
	Email        string
	CreationDate string
	BirthDate    string
	Uuid         string
}

var USER User

func Login(r *http.Request, db *sql.DB) {
	var email string
	var username string
	var password string
	var birtDate string
	var uuid string
	var creationDate string

	if r.Method == "POST" {
		fmt.Println("New POSTL (login)")
		//username will also work with email
		usernameInput := r.FormValue("input_loginusername")
		passwordInput := t.Hash(r.FormValue("input_loginpassword"))

		querry := fmt.Sprintf("SELECT username, password, email, creationDate, birthDate, uuid FROM users WHERE password = '%s'", passwordInput)
		row, err := db.Query(querry)
		if err != nil {
			fmt.Println(err)
		} else {
			for row.Next() {
				err = row.Scan(&username, &email, &password, &creationDate, &birtDate, &uuid)
				if err != nil {
					fmt.Println(err)
				} else {
					if 1 == 1 {
						fmt.Println("LOGIN!")
						USER.Username = username
						USER.BirthDate = birtDate
						USER.CreationDate = creationDate
						USER.Email = email
						USER.Uuid = uuid
					} else {
						USER.Username = "pierre"
						fmt.Println("no0b")
						fmt.Println(username, password, email, usernameInput, passwordInput)
					}
				}
			}
		}
	}
}
