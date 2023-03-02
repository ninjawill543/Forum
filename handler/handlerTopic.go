package forum

import (
	"database/sql"
	t "forum/topic"
	"html/template"
	"net/http"
)

func Handler_topicPage(w http.ResponseWriter, r *http.Request) {
	databaseTopics, _ := sql.Open("sqlite3", "../messages.db")
	t.TopicPageDisplay(databaseTopics, r)
	tmpl := template.Must(template.ParseFiles("../static/html/topicPage.html"))
	t.NewMessage(databaseTopics, r)
	tmpl.Execute(w, t.TOPIC)
}
