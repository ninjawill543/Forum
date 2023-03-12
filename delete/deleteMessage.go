package forum

import (
	"database/sql"
	"fmt"
	t "forum/users"
	"net/http"
)

func DeleteMessage(r *http.Request, db *sql.DB) {
	if r.FormValue("delete") != "" {
		fmt.Println("test")
		var owner string
		var uuidPath string

		query := fmt.Sprintf("SELECT owner, uuidPath FROM messages WHERE uuid = '%s'", r.FormValue("delete"))

		row, err := db.Query(query)
		if err != nil {
			fmt.Println(err)
		} else {
			for row.Next() {
				defer row.Close()
				err = row.Scan(&owner, &uuidPath)
				if err != nil {
					fmt.Println(err)
				}
			}
			if owner == t.USER.Username || t.USER.Admin == 1 {
				query2 := fmt.Sprintf("DELETE FROM messages WHERE uuid = '%s'", r.FormValue("delete"))
				db.Exec(query2)
				query = fmt.Sprintf("DELETE FROM likesFromUser WHERE uuidLiked = '%s'", r.FormValue("delete"))
				db.Exec(query)
				fmt.Println("message deleted")
				query3 := fmt.Sprintf("UPDATE topics SET nmbPosts = nmbPosts - 1 WHERE uuid = '%s'", uuidPath)
				db.Exec(query3)

			} else {
				fmt.Println("you need to be the owner of the topic to delete it")
			}
		}
	}
}
