package forum

import (
	"database/sql"
	"fmt"
	"net/http"
)

func DeleteAccount(r *http.Request, databaseUsers *sql.DB, databaseMessages *sql.DB, databaseTopics *sql.DB, w http.ResponseWriter) {
	if r.Method == "POST" {
		if USER.Username != "" {
			uuid := r.FormValue("delete")
			fmt.Println("account deleted", USER.Username)
			query2 := fmt.Sprintf("DELETE FROM messages WHERE owner = '%s'", USER.Username)
			databaseMessages.Exec(query2)
			query3 := fmt.Sprintf("DELETE FROM topics WHERE owner = '%s'", USER.Username)
			databaseTopics.Exec(query3)

			query := fmt.Sprintf("DELETE FROM users WHERE uuid = '%s'", uuid)
			databaseUsers.Exec(query)
			Logout(r)
			LogOutCookie(r, w)
			log := fmt.Sprintf("Account %s deleted same for all your messages and topics", USER.Username)
			fmt.Println(log)

		} else {
			fmt.Println("you need to be login to delete your account")
		}
	}
}
