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
	databaseUsers, _ := sql.Open("sqlite3", "../users.db")
	databaseMessages, _ := sql.Open("sqlite3", "../messages.db")
	databaseTopics, _ := sql.Open("sqlite3", "../topics.db")
	databaseReports, _ := sql.Open("sqlite3", "../reports.db")

	t2.ReportUser(r, databaseReports, databaseUsers)

	if r.FormValue("ban") != "" {
		t2.Ban(r, databaseUsers)
	}

	t.PublicProfil(r, databaseUsers, databaseMessages, databaseTopics)

	tmpl.Execute(w, t.PUBLICUSER)
}
