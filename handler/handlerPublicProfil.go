package forum

import (
	"database/sql"
	t "forum/profil"
	"html/template"
	"net/http"
)

func Handler_publicProfil(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../static/html/publicProfil.html"))
	databaseUsers, _ := sql.Open("sqlite3", "../users.db")
	databaseMessages, _ := sql.Open("sqlite3", "../messages.db")
	databaseTopics, _ := sql.Open("sqlite3", "../topics.db")

	t.PublicProfil(r, databaseUsers, databaseMessages, databaseTopics)

	tmpl.Execute(w, t.PUBLICUSER)
}
