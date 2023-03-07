package forum

import (
	"database/sql"
	"fmt"
	"net/http"
)

func Mp(r *http.Request, db *sql.DB) {
	if r.FormValue("mpMessage") != "" {
		fmt.Println(r.FormValue("mpMessage"))
	}
}
