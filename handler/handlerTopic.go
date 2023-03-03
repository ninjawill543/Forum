package forum

import (
	"database/sql"
	t "forum/topic"
	"html/template"
	"net/http"
)

func Handler_topicPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../static/html/topicPage.html"))
	databaseTopics, _ := sql.Open("sqlite3", "../messages.db")

	if r.FormValue("input_newMessage") != "" {
		t.NewMessage(databaseTopics, r)
	}

	t.TopicPageDisplay(databaseTopics, r)

	tmpl.Execute(w, t.TOPIC)
}
