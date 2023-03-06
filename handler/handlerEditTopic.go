package forum

import (
	"database/sql"
	t "forum/listTopics"
	"html/template"
	"net/http"
)

func Handler_EditTopic(w http.ResponseWriter, r *http.Request) {
	tmpl1 := template.Must(template.ParseFiles("../static/html/editTopic.html"))
	databaseForum, _ := sql.Open("sqlite3", "../forum.db")
	t.EditTopic(r, databaseForum)
	tmpl1.Execute(w, "")
}
