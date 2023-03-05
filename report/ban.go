package forum

import (
	"database/sql"
	"fmt"
	t2 "forum/profil"
	t "forum/users"
	"net/http"
)

func Ban(r *http.Request, db *sql.DB) {
	if t.USER.Admin == 1 {
		query := fmt.Sprintf("UPDATE users SET ban = 1 WHERE username = '%s'", t2.PUBLICUSER.Username)
		db.Exec(query)
	} else {
		fmt.Println("you got to be admin to ban someone")
	}
}
