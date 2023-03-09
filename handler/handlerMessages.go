package forum

import (
	"database/sql"
	t3 "forum/delete"
	t5 "forum/login"
	t "forum/messages"
	t2 "forum/report"
	t4 "forum/users"
	"html/template"
	"net/http"
)

func Handler_Messages(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../static/html/messages.html"))

	databaseForum, _ := sql.Open("sqlite3", "../forum.db")

	t4.GetCookieHandler(w, r)

	if r.FormValue("input_mail") != "" {
		t4.EmailStorage(r)
	}

	if r.FormValue("input_loginusername") != "" && r.FormValue("input_loginpassword") != "" {
		t5.Login(r, databaseForum, w)
	}

	if r.FormValue("input_username") != "" && r.FormValue("input_password") != "" && r.FormValue("input_birthDay") != "" {
		t4.Register(r, databaseForum)
	}

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

	tmpl.Execute(w, t.Messages)
}
