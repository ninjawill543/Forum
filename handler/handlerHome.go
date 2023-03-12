package forum

import (
	"database/sql"
	t "forum/home"
	"html/template"
	"net/http"
)

func Handler_Home(w http.ResponseWriter, r *http.Request) {
	tmpl1 := template.Must(template.ParseFiles("../static/html/home.html"))
	databaseForum, _ := sql.Open("sqlite3", "../forum.db")

	t.GetRandomMessages(databaseForum, r)

	tmpl1.Execute(w, t.TOPICSANDSESSION)
}
