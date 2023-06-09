package forum

import (
	"database/sql"
	"fmt"
	t2 "forum/structs"
	t "forum/views"
	"net/http"
	"time"
)

var EMAILSTORAGE t2.EmailStorage

func Register(r *http.Request, database *sql.DB) {
	if r.Method == "POST" {
		fmt.Println("start register")
		fmt.Println("New POST: (register) ")
		var checkAll bool
		username := r.FormValue("input_username")
		password := r.FormValue("input_password")
		password2 := r.FormValue("input_password2")
		creationDate := time.Now().String()
		birthDay := r.FormValue("input_birthDay")
		mail := EMAILSTORAGE.Email

		if password != password2 {
			fmt.Println("passowrds dont match")
		}

		if len(username) < 5 || len(username) > 14 {
			fmt.Println("invalid username")
			checkAll = true
		}

		if !t.CheckPassword(password) {
			checkAll = true
		}

		if !t.CheckMail(mail) {
			checkAll = true
		}

		if !checkAll {
			AddUsers(database, username, t.Hash(password), mail, creationDate, birthDay)
		}
	}
}
