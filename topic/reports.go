package forum

import (
	"database/sql"
	"fmt"
	t "forum/users"
	"net/http"
)

func Reports(r *http.Request, db *sql.DB) {
	if r.Method == "POST" {
		if t.USER.Username != "" {
			uuid := r.FormValue("report")
			report := "report"

			query := fmt.Sprintf("UPDATE messages SET %s = %s + 1 WHERE uuid = '%s'", report, report, uuid)
			db.Exec(query)
		} else {
			fmt.Println("you need to be login to report a message")
		}
	}
}
