package forum

import (
	"database/sql"
	t "forum/users"
	"html/template"
	"net/http"
)

func Handler_profil(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../static/html/profil.html"))
	databaseUsers, _ := sql.Open("sqlite3", "../users.db")
	databaseMessages, _ := sql.Open("sqlite3", "../messages.db")
	databaseTopics, _ := sql.Open("sqlite3", "../topics.db")

	if r.FormValue("logOutButton") == "logout" {
		t.Logout(r)
	} else if r.FormValue("delete") != "" {
		t.DeleteAccount(r, databaseUsers, databaseMessages, databaseTopics)
	} else if r.FormValue("username") != "" || r.FormValue("email") != "" || r.FormValue("password") != "" {
		t.UserEdit(r, databaseUsers)
	}
	t.MessagesSendByUser(databaseMessages)
	t.TopicCreatedByUser(databaseTopics)

	tmpl.Execute(w, t.USER)
}
