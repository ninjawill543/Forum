package forum

import (
	"database/sql"
	t "forum/profil"
	t2 "forum/report"
	"html/template"
	"net/http"
)

func Handler_publicProfil(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../static/html/publicProfil.html"))
	databaseForum, _ := sql.Open("sqlite3", "../forum.db")

	t2.ReportUser(r, databaseForum)

	if r.FormValue("ban") != "" {
		t2.Ban(r, databaseForum)
	}

	t.PublicProfil(r, databaseForum)

	tmpl.Execute(w, t.PUBLICUSER)
}
