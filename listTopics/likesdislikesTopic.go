package forum

import (
	"database/sql"
	"fmt"
	t "forum/users"
	"net/http"
)

func LikesDislikes(r *http.Request, databaseTopics *sql.DB, databaseLikesFromUsers *sql.DB) {
	if t.USER.Username != "" {
		var likesordislikes int
		var uuid string
		var uuidLiked string
		var alreadyLiked bool
		var likeOrDislike int

		if r.Method == "POST" {
			if r.FormValue("like") != "" {
				likesordislikes = 1
				uuid = r.FormValue("like")

			}
			if r.FormValue("dislike") != "" {
				likesordislikes = -1
				uuid = r.FormValue("dislike")
			}
			query := fmt.Sprintf("SELECT uuidLiked, likeOrDislike FROM likesFromUser WHERE uuidUser = '%s'", t.USER.Uuid)
			row, err := databaseLikesFromUsers.Query(query)
			if err != nil {
				fmt.Println(err)
			} else {
				for row.Next() {
					row.Scan(&uuidLiked, &likeOrDislike)
					if uuidLiked == uuid {
						alreadyLiked = true
					}
				}
				if alreadyLiked {
					fmt.Println("already liked")
					// fmt.Println(likeOrDislike, "previous like")
					// fmt.Println(likesordislikes, "new like")
					// if likeOrDislike == 1 && likesordislikes == 1 {
					// 	query = fmt.Sprintf("UPDATE topics SET likes = likes - 1 WHERE uuid = '%s'", uuid)
					// 	databaseTopics.Exec(query)
					// } else if likeOrDislike == 1 && likesordislikes == -1 {
					// 	query = fmt.Sprintf("UPDATE topics SET likes = likes - 2 WHERE uuid = '%s'", uuid)
					// 	databaseTopics.Exec(query)
					// } else if likeOrDislike == -1 && likesordislikes == 1 {
					// 	query = fmt.Sprintf("UPDATE topics SET likes = likes + 2 WHERE uuid = '%s'", uuid)
					// 	databaseTopics.Exec(query)
					// } else if likeOrDislike == -1 && likesordislikes == -1 {
					// 	query = fmt.Sprintf("UPDATE topics SET likes = likes + 1 WHERE uuid = '%s'", uuid)
					// 	databaseTopics.Exec(query)
					// }
				} else {
					query = fmt.Sprintf("UPDATE topics SET likes = likes + %d WHERE uuid = '%s'", likesordislikes, uuid)
					databaseTopics.Exec(query)

					query2 := fmt.Sprintf("INSERT INTO likesFromUser(uuidUser, uuidLiked, likeOrDislike) VALUES('%s', '%s', '%d')", t.USER.Uuid, uuid, likesordislikes)
					databaseLikesFromUsers.Exec(query2)
				}
			}
		}
	} else {
		fmt.Println("you need to be login to like or dislike a message")
	}
}
