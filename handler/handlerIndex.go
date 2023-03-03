package forum

import (
	"database/sql"
	t "forum/listTopics"
	t2 "forum/users"
	"html/template"
	"net/http"
)

func Handler_index(w http.ResponseWriter, r *http.Request) {
	tmpl1 := template.Must(template.ParseFiles("../static/html/index.html"))
	databaseUsers, _ := sql.Open("sqlite3", "../users.db")
	databaseTopics, _ := sql.Open("sqlite3", "../topics.db")

	if r.FormValue("input_loginusername") != "" && r.FormValue("input_loginpassword") != "" {
		t2.Login(r, databaseUsers)
	}

	if r.FormValue("topic_name") != "" {
		t.AddTopic(r, databaseTopics)
	}

	if r.FormValue("input_username") != "" && r.FormValue("input_password") != "" && r.FormValue("input_mail") != "" && r.FormValue("input_birthDay") != "" {
		t2.Register(r, databaseUsers)
	}

	if r.FormValue("like") != "" || r.FormValue("dislike") != "" {
		t.LikesDislikes(r, databaseTopics)
	}

	t.DisplayTopic(databaseTopics)

	tmpl1.Execute(w, t.TOPICS)
}
