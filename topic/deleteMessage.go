package forum

import (
	"database/sql"
	"fmt"
	t "forum/users"
	"net/http"
)

func DeleteMessage(r *http.Request, db *sql.DB) {
	if r.FormValue("delete") != "" {
		var owner string
		query := fmt.Sprintf("SELECT owner FROM messages WHERE uuid = '%s'", r.FormValue("delete"))
		row, err := db.Query(query)
		if err != nil {
			fmt.Println(err)
		} else {
			for row.Next() {
				err = row.Scan(&owner)
				fmt.Println(owner)
				if err != nil {
					fmt.Println(err)
				}
			}
			if owner == t.USER.Username {
				query2 := fmt.Sprintf("DELETE FROM messages WHERE uuid = '%s'", r.FormValue("delete"))
				db.Exec(query2)
			} else {
				fmt.Println("you need to be the owner of the topic to delete it")
			}
		}
	}
}
