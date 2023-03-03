package forum

import (
	"database/sql"
	"fmt"
	t "forum/users"
	"net/http"
)

func LikesDislikes(r *http.Request, db *sql.DB) {
	if t.USER.Username != "" {
		var likesordislikes string
		var uuid string
		fmt.Println(likesordislikes)

		if r.Method == "POST" {
			if r.FormValue("like") != "" {
				likesordislikes = "likes"
				uuid = r.FormValue("like")

			}
			if r.FormValue("dislike") != "" {
				likesordislikes = "dislikes"
				uuid = r.FormValue("dislike")
			}
			query := fmt.Sprintf("UPDATE topics SET %s = %s + 1 WHERE uuid = '%s'", likesordislikes, likesordislikes, uuid)
			db.Exec(query)
		}
	} else {
		fmt.Println("you need to be login to like or dislike a message")
	}
}
