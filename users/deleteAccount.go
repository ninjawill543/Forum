package forum

import (
	"database/sql"
	"fmt"
	"net/http"
)

func DeleteAccount(r *http.Request, db *sql.DB) {
	if r.Method == "POST" {
		if USER.Username != "" {
			uuid := r.FormValue("delete")
			query := fmt.Sprintf("DELETE FROM users WHERE uuid = '%s'", uuid)
			fmt.Println("account deleted", USER.Username)
			Logout(r)
			db.Exec(query)
		} else {
			fmt.Println("you need to be login to delete your account")
		}
	}
}
