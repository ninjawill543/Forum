package forum

import (
	"database/sql"
	"fmt"
	"net/http"
)

func DeleteAccount(r *http.Request, db *sql.DB, w http.ResponseWriter) {
	if r.Method == "POST" {
		if USER.Username != "" {
			fmt.Println("account deleted", USER.Username)
			query2 := fmt.Sprintf("DELETE FROM messages WHERE owner = '%s'", USER.Username)
			db.Exec(query2)
			query3 := fmt.Sprintf("DELETE FROM topics WHERE owner = '%s'", USER.Username)
			db.Exec(query3)

			query := fmt.Sprintf("DELETE FROM users WHERE username = '%s'", USER.Username)
			db.Exec(query)
			Logout(r)
			LogOutCookie(r, w)
			log := fmt.Sprintf("Account %s deleted same for all your messages and topics", USER.Username)
			fmt.Println(log)

		} else {
			fmt.Println("you need to be login to delete your account")
		}
	}
}
