package forum

import (
	"database/sql"
	t "forum/home"
	t5 "forum/login"
	t4 "forum/users"
	"html/template"
	"net/http"
)

func Handler_Home(w http.ResponseWriter, r *http.Request) {
	tmpl1 := template.Must(template.ParseFiles("../static/html/home.html"))
	databaseForum, _ := sql.Open("sqlite3", "../forum.db")

	t4.GetCookieHandler(w, r)

	if r.FormValue("input_loginusername") != "" && r.FormValue("input_loginpassword") != "" {
		t5.Login(r, databaseForum, w)
	}

	if r.FormValue("input_username") != "" && r.FormValue("input_password") != "" && r.FormValue("input_birthDay") != "" {
		t4.Register(r, databaseForum)
	}

	t.GetRandomMessages(databaseForum, r)

	t.TOPICSANDSESSION.SessionUser = t4.USER.Username

	tmpl1.Execute(w, t.TOPICSANDSESSION)
}
