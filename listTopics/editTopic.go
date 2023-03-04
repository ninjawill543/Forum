package forum

import (
	"database/sql"
	"fmt"
	t "forum/users"
	"net/http"
	"strings"
)

func EditTopic(r *http.Request, db *sql.DB) {
	if r.FormValue("edit") != "" {
		newName := r.FormValue("newName")
		if len(newName) < 5 {
			fmt.Println("not enough char")
		} else {
			var owner string
			uuid := strings.Split(r.URL.Path, "/")
			query := fmt.Sprintf("SELECT owner FROM topics WHERE uuid = '%s'", uuid[2])
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
			}
			if t.USER.Username == owner {
				fmt.Println(newName)
				query := fmt.Sprintf("UPDATE topics SET name = '%s' WHERE uuid = '%s'", newName, uuid[2])
				db.Exec(query)
			} else {
				fmt.Println("you need to be the owner of the message to edit it")
			}
		}
	}
}
