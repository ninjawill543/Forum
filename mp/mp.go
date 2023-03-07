package forum

import (
	"database/sql"
	"fmt"
	t "forum/users"
	"net/http"
	"strings"
	"time"
)

func AddMp(r *http.Request, db *sql.DB) {
	if r.FormValue("mpMessage") != "" {
		user2 := strings.Split(r.URL.Path, "/")
		if t.USER.Username == user2[2] {
			fmt.Println("you cant mp yourself")
		} else {
			query := fmt.Sprintf("INSERT INTO mp(user1, user2, creationDate, message) VALUES('%s', '%s', '%s', '%s')", t.USER.Username, user2[2], time.Now(), r.FormValue("mpMessage"))
			db.Exec(query)
		}
	}
}
