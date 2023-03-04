package forum

import (
	"database/sql"
	t "forum/messages"
	"html/template"
	"net/http"
)

func Handler_EditMessage(w http.ResponseWriter, r *http.Request) {
	tmpl1 := template.Must(template.ParseFiles("../static/html/editMessage.html"))
	databaseMessages, _ := sql.Open("sqlite3", "../messages.db")
	t.EditMessage(r, databaseMessages)
	tmpl1.Execute(w, "")
}
