package forum

import (
	"database/sql"
	t "forum/mp"
	t2 "forum/users"
	"html/template"
	"net/http"
)

func Handler_Mp(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../static/html/mp.html"))

	databaseForum, _ := sql.Open("sqlite3", "../forum.db")

	t2.GetCookieHandler(w, r)

	if r.FormValue("mpMessage") != "" {
		t.AddMp(r, databaseForum)
	}

	t.DisplayMp(r, databaseForum)

	t2.GetCookieHandler(w, r)

	tmpl.Execute(w, t.MPSANDTOWHO)
}
