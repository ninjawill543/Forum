package forum

import (
	"database/sql"
	"fmt"
	t2 "forum/users"
	t "forum/views"
	"net/http"
)

func Login(r *http.Request, db *sql.DB, w http.ResponseWriter) {
	var email string
	var username string
	var password string
	var birtDate string
	var uuid string
	var creationDate string
	var admin int
	var ban int
	fmt.Println("big test")

	if r.Method == "POST" {
		if t2.USER.Username != "" {
			fmt.Println("you're already login")
		} else {
			//username will also work with email
			usernameInput := r.FormValue("input_loginusername")
			passwordInput := t.Hash(r.FormValue("input_loginpassword"))

			querry := fmt.Sprintf("SELECT username, password, email, creationDate, birthDate, admin, uuid, ban FROM users WHERE password = '%s'", passwordInput)
			row, err := db.Query(querry)
			if err != nil {
				fmt.Println(err)
			} else {
				for row.Next() {
					defer row.Close()
					err = row.Scan(&username, &password, &email, &creationDate, &birtDate, &admin, &uuid, &ban)
					if err != nil {
						fmt.Println(err)
					} else if ban == 1 {
						fmt.Println("you have been banned ask admin to get unbanned")
					} else {
						if usernameInput == username || usernameInput == email && passwordInput == password {
							fmt.Println("LOGIN!")
							t2.USER.Username = username
							t2.USER.BirthDate = birtDate
							t2.USER.CreationDate = creationDate
							t2.USER.Email = email
							t2.USER.Uuid = uuid
							t2.USER.Admin = admin

							t2.SetCookieHandler(w, r)
						} else {
							fmt.Println("no0b")
							fmt.Println(username, password, email, usernameInput, passwordInput)
						}
					}
				}
			}
		}
	}
}
