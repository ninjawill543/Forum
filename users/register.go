package forum

import (
	"database/sql"
	"fmt"
	t "forum/views"
	"net/http"
	"time"
)

func Register(r *http.Request, database *sql.DB) {
	if r.Method == "POST" {
		fmt.Println("New POST: (register) ")
		var checkAll bool
		username := r.FormValue("input_username")
		password := r.FormValue("input_password")
		mail := r.FormValue("input_mail")
		creationDate := time.Now()
		birthDay := r.FormValue("input_birthDay")

		if len(username) < 5 || len(username) > 14 {
			fmt.Println("invalid username")
			checkAll = true
		}

		if t.CheckPassword(password) == false {
			checkAll = true
		}

		if t.CheckMail(mail) == false {
			checkAll = true
		}

		if checkAll == false {
			AddUsers(database, username, t.Hash(password), mail, creationDate, birthDay)
		}
	}
}
