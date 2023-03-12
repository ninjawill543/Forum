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
		var firstMessage string
		var message string

		query := fmt.Sprintf("SELECT owner, uuidPath, message FROM messages WHERE uuid = '%s'", r.FormValue("delete"))

		row, err := db.Query(query)
		if err != nil {
			fmt.Println(err)
		} else {
			for row.Next() {
				defer row.Close()
				err = row.Scan(&owner, &uuidPath, &message)
				if err != nil {
					fmt.Println(err)
				}
			}

			query2 := fmt.Sprintf("SELECT firstMessage FROM topics WHERE uuid = '%s'", uuidPath)
			row, err := db.Query(query2)
			if err != nil {
				fmt.Println(err)
			} else {
				for row.Next() {
					defer row.Close()
					err = row.Scan(&firstMessage)
					if err != nil {
						fmt.Println(err)
					}
				}
			}

			fmt.Println(firstMessage, message)

			if firstMessage == message {
				fmt.Println("you can't delete the first message of a topic")
			} else {
				if owner == t.USER.Username || t.USER.Admin == 1 {
					query2 := fmt.Sprintf("DELETE FROM messages WHERE uuid = '%s'", r.FormValue("delete"))
					db.Exec(query2)
					query = fmt.Sprintf("DELETE FROM likesFromUser WHERE uuidLiked = '%s'", r.FormValue("delete"))
					db.Exec(query)
					fmt.Println("message deleted")
					query3 := fmt.Sprintf("UPDATE topics SET nmbPosts = nmbPosts - 1 WHERE uuid = '%s'", uuidPath)
					db.Exec(query3)

				} else {
					fmt.Println("you need to be the owner of the topic to delete it or cant delete first message of topic, delete topic instead")
				}
			}
		}
	}
}
