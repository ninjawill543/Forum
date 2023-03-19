package forum

import (
	"database/sql"
	t3 "forum/delete"
	t "forum/home"
	t2 "forum/listTopics"
	t5 "forum/login"
	t4 "forum/users"
	"html/template"
	"net/http"
)

func Handler_Home(w http.ResponseWriter, r *http.Request) {
	tmpl1 := template.Must(template.ParseFiles("../static/html/home.html"))
	databaseForum, _ := sql.Open("sqlite3", "../forum.db")

	if r.FormValue("like") != "" || r.FormValue("dislike") != "" {
		t2.LikesDislikes(r, databaseForum)
	}

	if r.FormValue("delete") != "" {
		t3.DeleteTopic(r, databaseForum)
	}

	if r.FormValue("input_mail") != "" {
		t4.EMAILSTORAGE.Email = r.FormValue("input_mail")
	}

	if r.FormValue("input_loginusername") != "" && r.FormValue("input_loginpassword") != "" {
		t5.Login(r, databaseForum, w)
	}

	if r.FormValue("input_username") != "" && r.FormValue("input_password") != "" && r.FormValue("input_birthDay") != "" {
		t4.Register(r, databaseForum)
	}

	t.GetRandomMessages(databaseForum, r)

	t.TOPICSANDSESSION.SessionUser = t4.USER.Username
	t4.GetCookieHandler(w, r)

	tmpl1.Execute(w, t.TOPICSANDSESSION)
}
