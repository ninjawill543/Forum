package forum

import (
	"database/sql"
	"fmt"
	t "forum/users"
	"net/http"
)

func LikesDislikes(r *http.Request, db *sql.DB) {
	fmt.Println("try later error")
	if t.USER.Username != "" {
		var likesordislikes int
		var uuid string

		if r.Method == "POST" {
			if r.FormValue("like") != "" {
				likesordislikes = 1
				uuid = r.FormValue("like")

			}
			if r.FormValue("dislike") != "" {
				likesordislikes = -1
				uuid = r.FormValue("dislike")
			}
			query := fmt.Sprintf("UPDATE messages SET like = like + %d WHERE uuid = '%s'", likesordislikes, uuid)
			db.Exec(query)

			query2 := fmt.Sprintf("INSERT INTO likesFromUser(uuidUser, uuidLiked, likeOrDislike) VALUES('%s', '%s', '%d')", t.USER.Uuid, uuid, likesordislikes)
			db.Exec(query2)
		}
	} else {
		fmt.Println("you need to be login to like or dislike a message")
	}
}
