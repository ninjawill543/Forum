package forum

import (
	"database/sql"
	"fmt"
	"net/http"
)

func DeleteAccount(r *http.Request, db *sql.DB, w http.ResponseWriter) {
	if r.Method == "POST" {
		if USER.Username == r.FormValue("delete") {
			fmt.Println("account deleted", USER.Username)
			query2 := fmt.Sprintf("DELETE FROM messages WHERE owner = '%s'", USER.Username)
			query3 := fmt.Sprintf("DELETE FROM topics WHERE owner = '%s'", USER.Username)

			query := fmt.Sprintf("DELETE FROM users WHERE username = '%s'", USER.Username)
			db.Exec(query)
			query4 := fmt.Sprintf("DELETE FROM mp WHERE user1 ='%s' OR user2 = '%s'", USER.Username, USER.Username)
			db.Exec(query2)
			db.Exec(query3)
			db.Exec(query4)
			Logout(r)
			LogOutCookie(r, w)
			log := fmt.Sprintf("Account %s deleted same for all your messages and topics", USER.Username)
			fmt.Println(log)

		} else {
			fmt.Println("you need to be login to delete your account")
		}
	}
}
