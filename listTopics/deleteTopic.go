package forum

import (
	"database/sql"
	"fmt"
	t "forum/users"
	"net/http"
)

func DeleteTopic(r *http.Request, db *sql.DB) {
	if r.FormValue("delete") != "" {
		var owner string
		uuid := r.FormValue("delete")
		query := fmt.Sprintf("SELECT owner FROM topics WHERE uuid = '%s'", uuid)
		row, err := db.Query(query)
		if err != nil {
			fmt.Println(err)
		} else {
			for row.Next() {
				err = row.Scan(&owner)
				if err != nil {
					fmt.Println(err)
				}
			}
			if t.USER.Username == owner || t.USER.Admin == "yes" {
				query = fmt.Sprintf("DELETE FROM topics WHERE uuid = '%s'", uuid)
				db.Exec(query)
				fmt.Println("topic deleted")
			} else {
				fmt.Println("you need to be the owner of the topic to delete it")
			}
		}
	}
}
