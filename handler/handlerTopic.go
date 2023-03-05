package forum

import (
	"database/sql"
	t3 "forum/delete"
	t "forum/messages"
	t2 "forum/report"
	"html/template"
	"net/http"
)

func Handler_topicPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../static/html/topicPage.html"))
	databaseMessages, _ := sql.Open("sqlite3", "../messages.db")
	databaseTopics, _ := sql.Open("sqlite3", "../topics.db")
	databaseLikeFromUsers, _ := sql.Open("sqlite3", "../likesFromUser.db")
	databaseReports, _ := sql.Open("sqlite3", "../reports.db")

	if r.FormValue("delete") != "" {
		t3.DeleteMessage(r, databaseMessages)
	}
	if r.FormValue("input_newMessage") != "" {
		t.NewMessage(databaseMessages, r)
	}
	if r.FormValue("report") != "" {
		t2.ReportMessage(r, databaseMessages, databaseReports)
	}
	if r.FormValue("like") != "" || r.FormValue("dislike") != "" {
		t.LikesDislikes(r, databaseTopics, databaseLikeFromUsers)
	}
	if r.FormValue("edit") != "" {
		t.EditMessage(r, databaseMessages)
	}

	t.MessagesPageDisplay(databaseMessages, databaseTopics, r)

	tmpl.Execute(w, t.TOPIC)
}
