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
			if t.USER.Username == owner || t.USER.Admin == 1 {
				query = fmt.Sprintf("DELETE FROM topics WHERE uuid = '%s'", uuid)
				db.Exec(query)
				fmt.Println("topic deleted")
				databaseMessages, _ := sql.Open("sqlite3", "../messages.db")
				query = fmt.Sprintf("DELETE FROM messages WHERE uuidPath = '%s'", uuid)
				databaseMessages.Exec(query)
				databaseLikeFromUsers, _ := sql.Open("sqlite3", "../likesFromUser.db")
				query = fmt.Sprintf("DELETE FROM likesFromUser WHERE uuidLiked = '%s'", uuid)
				databaseLikeFromUsers.Exec(query)
			} else {
				fmt.Println("you need to be the owner of the topic to delete it")
			}
		}
	}
}
