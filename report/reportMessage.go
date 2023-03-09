package forum

import (
	"database/sql"
	"fmt"
	t "forum/users"
	"net/http"
)

func ReportMessage(r *http.Request, db *sql.DB) {
	if r.Method == "POST" {
		var uuidReported string
		var alreadyReported bool
		uuid := r.FormValue("report")
		if t.USER.Username != "" {

			query := fmt.Sprintf("SELECT uuidReported from reports WHERE uuidUser = '%s'", t.USER.Uuid)
			row, err := db.Query(query)
			if err != nil {
				fmt.Println(err)
			} else {
				for row.Next() {
					defer row.Close()
					err = row.Scan(&uuidReported)
					if err != nil {
						fmt.Println(err)
					} else if uuidReported == uuid {
						alreadyReported = true
					}
				}
			}
			if alreadyReported {
				fmt.Println("arleady reported")
			} else {
				query = fmt.Sprintf("INSERT into reports(uuidUser, uuidReported) VALUES ('%s', '%s')", t.USER.Uuid, uuid)
				db.Exec(query)
				report := "report"
				query = fmt.Sprintf("UPDATE messages SET %s = %s + 1 WHERE uuid = '%s'", report, report, uuid)

				db.Exec(query)
			}
		} else {
			fmt.Println("you need to be login to report a message")
		}
		db.Exec("DELETE FROM messages WHERE report >= 5")
	}
}
