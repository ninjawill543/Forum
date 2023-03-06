package forum

import (
	"database/sql"
	t3 "forum/delete"
	t "forum/messages"
	t2 "forum/report"
	t4 "forum/users"
	"html/template"
	"net/http"
)

func Handler_topicPage(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template
	if t4.USER.Username == "" {
		tmpl = template.Must(template.ParseFiles("../static/html/topicPagenotlogin.html"))

	} else {
		tmpl = template.Must(template.ParseFiles("../static/html/topicPage.html"))

	}
	databaseForum, _ := sql.Open("sqlite3", "../forum.db")

	t4.GetCookieHandler(w, r)

	if r.FormValue("delete") != "" {
		t3.DeleteMessage(r, databaseForum)
	}
	if r.FormValue("input_newMessage") != "" {
		t.NewMessage(databaseForum, r)
	}
	if r.FormValue("report") != "" {
		t2.ReportMessage(r, databaseForum)
	}
	if r.FormValue("like") != "" || r.FormValue("dislike") != "" {
		t.LikesDislikes(r, databaseForum)
	}
	if r.FormValue("edit") != "" {
		t.EditMessage(r, databaseForum)
	}

	t.MessagesPageDisplay(databaseForum, r)

	tmpl.Execute(w, t.TOPIC)
}
