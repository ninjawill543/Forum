package forum

import (
	"database/sql"
	"fmt"
	t "forum/users"
	"net/http"
)

func LikesDislikes(r *http.Request, db *sql.DB) {
	if t.USER.Username != "" {
		var newLike int
		var uuid string
		var uuidLiked string
		var alreadyLiked bool
		var previousLike int

		if r.Method == "POST" {
			if r.FormValue("like") != "" {
				newLike = 1
				uuid = r.FormValue("like")

			}
			if r.FormValue("dislike") != "" {
				newLike = -1
				uuid = r.FormValue("dislike")
			}
			query := fmt.Sprintf("SELECT uuidLiked, likeOrDislike FROM likesFromUser WHERE uuidUser = '%s'", t.USER.Uuid)
			row, err := db.Query(query)
			if err != nil {
				fmt.Println(err)
			} else {
				for row.Next() {
					defer row.Close()
					row.Scan(&uuidLiked, &previousLike)
					if uuidLiked == uuid {
						alreadyLiked = true
					}
				}
				if alreadyLiked {
					fmt.Println("already liked")
					fmt.Println(previousLike, "previous like")
					fmt.Println(newLike, "new like")
					// if previousLike == 1 && newLike == 1 {
					// 	query = fmt.Sprintf("UPDATE topics SET likes = likes - 1 WHERE uuid = '%s'", uuid)
					// 	db.Exec(query)
					// } else if previousLike == 1 && newLike == -1 {
					// 	query = fmt.Sprintf("UPDATE topics SET likes = likes - 2 WHERE uuid = '%s'", uuid)
					// 	db.Exec(query)
					// } else if previousLike == -1 && newLike == 1 {
					// 	query = fmt.Sprintf("UPDATE topics SET likes = likes + 2 WHERE uuid = '%s'", uuid)
					// 	db.Exec(query)
					// } else if previousLike == -1 && newLike == -1 {
					// 	query = fmt.Sprintf("UPDATE topics SET likes = likes + 1 WHERE uuid = '%s'", uuid)
					// 	db.Exec(query)
					// }
				} else {
					query = fmt.Sprintf("UPDATE topics SET likes = likes + %d WHERE uuid = '%s'", newLike, uuid)
					db.Exec(query)

					query2 := fmt.Sprintf("INSERT INTO likesFromUser(uuidUser, uuidLiked, likeOrDislike) VALUES('%s', '%s', '%d')", t.USER.Uuid, uuid, newLike)
					db.Exec(query2)
				}
			}
		}
	} else {
		fmt.Println("you need to be login to like or dislike a message")
	}
}
