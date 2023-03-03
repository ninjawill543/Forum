package forum

import (
	"database/sql"
	"fmt"
	t "forum/users"
	"net/http"
)

func LikesDislikes(r *http.Request, db *sql.DB) {
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
			query := fmt.Sprintf("UPDATE topics SET likes = likes + %d WHERE uuid = '%s'", likesordislikes, uuid)
			db.Exec(query)
		}
	} else {
		fmt.Println("you need to be login to like or dislike a message")
	}
}
