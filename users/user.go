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
	Admin        string
}

var USER User

func UserEdit(r *http.Request, db *sql.DB) {
	if r.Method == "POST" {
		if USER.Username != "" {
			if r.FormValue("username") != "" {
				if len(r.FormValue("username")) < 5 || len(r.FormValue("username")) > 14 {
					fmt.Println("invalid username")
				} else {
					query := fmt.Sprintf("UPDATE users SET username = '%s' WHERE uuid = '%s'", r.FormValue("username"), USER.Uuid)
					db.Exec(query)
				}
			}
			if r.FormValue("password") != "" {
				if t.CheckPassword(r.FormValue("password")) {
					query := fmt.Sprintf("UPDATE users SET password = '%s' WHERE uuid = '%s'", t.Hash(r.FormValue("password")), USER.Uuid)
					db.Exec(query)
				}
			}
			if r.FormValue("email") != "" {
				if t.CheckMail(r.FormValue("email")) {
					query := fmt.Sprintf("UPDATE users SET email = '%s' WHERE uuid = '%s'", r.FormValue("email"), USER.Uuid)
					db.Exec(query)
				}
			}
			Logout(r)
		} else {
			fmt.Println("you need to be login to edit your account")
		}
	}
}
