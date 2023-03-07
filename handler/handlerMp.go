package forum

import (
	"database/sql"
	t "forum/mp"
	"html/template"
	"net/http"
)

func Handler_Mp(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../static/html/mp.html"))

	databaseForum, _ := sql.Open("sqlite3", "../forum.db")

	if r.FormValue("mpMessage") != "" {
		t.Mp(r, databaseForum)
	}

	tmpl.Execute(w, "")
}
