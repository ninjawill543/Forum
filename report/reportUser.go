package forum

import (
	"database/sql"
	"fmt"
	t2 "forum/profil"
	t "forum/users"
	"net/http"
)

func ReportUser(r *http.Request, db *sql.DB) {
	if r.Method == "POST" {
		var uuidReported string
		var alreadyReported bool
		uuid := r.FormValue("report")
		if t.USER.Username != "" && t.USER.Username != r.FormValue("report") {
			query := fmt.Sprintf("SELECT uuidReported from reports WHERE uuidUser = '%s'", t.USER.Uuid)
			row, err := db.Query(query)
			if err != nil {
				fmt.Println(err)
			} else {
				for row.Next() {
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
				report := "reports"
				query = fmt.Sprintf("UPDATE users SET %s = %s + 1 WHERE username = '%s'", report, report, uuid)

				db.Exec(query)
			}
		} else {
			fmt.Println("you need to be login to report a message OR CANT REPORT YOURSELF")
		}

		if t2.PUBLICUSER.Reports >= 9 {
			query := fmt.Sprintf("UPDATE users SET ban = 1 WHERE username = '%s'", t2.PUBLICUSER.Username)
			db.Exec(query)
		}
	}
}
