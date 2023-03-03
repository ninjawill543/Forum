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
		var uuidPath string
		query := fmt.Sprintf("SELECT owner, uuidPath FROM messages WHERE uuid = '%s'", r.FormValue("delete"))
		row, err := db.Query(query)
		if err != nil {
			fmt.Println(err)
		} else {
			for row.Next() {
				err = row.Scan(&owner, &uuidPath)
				fmt.Println(owner)
				if err != nil {
					fmt.Println(err)
				}
			}
			if owner == t.USER.Username {
				query2 := fmt.Sprintf("DELETE FROM messages WHERE uuid = '%s'", r.FormValue("delete"))
				db.Exec(query2)
				fmt.Println("message deleted")
				query3 := fmt.Sprintf("UPDATE topics SET nmbPosts = nmbPosts - 1 WHERE uuid = '%s'", uuidPath)
				fmt.Println(query3)
				databaseTopics, _ := sql.Open("sqlite3", "../topics.db")
				databaseTopics.Exec(query3)

			} else {
				fmt.Println("you need to be the owner of the topic to delete it")
			}
		}
	}
}
