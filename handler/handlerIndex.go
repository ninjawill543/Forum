package forum

import (
	"database/sql"
	t "forum/listTopics"
	t2 "forum/users"
	"html/template"
	"net/http"
)

func Handler_index(w http.ResponseWriter, r *http.Request) {
	databaseUsers, _ := sql.Open("sqlite3", "../users.db")
	databaseTopics, _ := sql.Open("sqlite3", "../topic.db")
	t.DisplayTopic(databaseTopics)

	tmpl1 := template.Must(template.ParseFiles("../static/html/index.html"))

	//register on specifig button
	// if button (register)...
	t2.Login(r, databaseUsers)
	t.AddTopic(r, databaseTopics)
	t2.Register(r, databaseUsers)

	tmpl1.Execute(w, t.TOPICS)
}
