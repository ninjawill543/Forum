package forum

import (
	"database/sql"
	t "forum/topic"
	"html/template"
	"net/http"
)

func Handler_topicPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../static/html/topicPage.html"))
	databaseMessages, _ := sql.Open("sqlite3", "../messages.db")
	databaseTopics, _ := sql.Open("sqlite3", "../topics.db")

	if r.FormValue("delete") != "" {
		t.DeleteMessage(r, databaseMessages)
	}
	if r.FormValue("input_newMessage") != "" {
		t.NewMessage(databaseMessages, r)
	}
	if r.FormValue("report") != "" {
		t.Reports(r, databaseMessages)
	}
	if r.FormValue("like") != "" || r.FormValue("dislike") != "" {
		t.LikesDislikes(r, databaseMessages)
	}
	if r.FormValue("edit") != "" {
		t.EditMessage(r, databaseMessages)
	}

	t.TopicPageDisplay(databaseMessages, databaseTopics, r)

	tmpl.Execute(w, t.TOPIC)
}
