package forum

import (
	"database/sql"
	t3 "forum/delete"
	t "forum/listTopics"
	t4 "forum/login"
	t2 "forum/users"
	"html/template"
	"net/http"
	"strings"
)

func Handler_topics(w http.ResponseWriter, r *http.Request) {

	tmpl1 := template.Must(template.ParseFiles("../static/html/topics.html"))
	tmpl2 := template.Must(template.ParseFiles("../static/html/404.html"))

	databaseForum, _ := sql.Open("sqlite3", "../forum.db")

	if r.FormValue("topic_name") != "" {
		t.AddTopic(r, databaseForum)
	}

	if r.FormValue("input_mail") != "" {
		t2.EMAILSTORAGE.Email = r.FormValue("input_mail")
	}

	if r.FormValue("input_loginusername") != "" && r.FormValue("input_loginpassword") != "" {
		t4.Login(r, databaseForum, w)
	}

	if r.FormValue("input_username") != "" && r.FormValue("input_password") != "" && r.FormValue("input_birthDay") != "" {
		t2.Register(r, databaseForum)
	}

	if r.FormValue("like") != "" || r.FormValue("dislike") != "" {
		t.LikesDislikes(r, databaseForum)
	}

	if r.FormValue("delete") != "" {
		t3.DeleteTopic(r, databaseForum)
	}

	if r.URL.Path == "/topics/category=tech" || r.URL.Path == "/topics/category=dinosaurs" || r.URL.Path == "/topics/category=watches" || r.URL.Path == "/topics/category=sneakers" || r.URL.Path == "/topics/category=gardening" || r.URL.Path == "/topics/category=video-games" || r.URL.Path == "/topics/category=climbing" {
		t.DisplayTopic(r, databaseForum)
		t.TOPICSANDSESSION.SessionUser = t2.USER.Username
		category := strings.Split(r.URL.Path, "/")[2]
		t.TOPICSANDSESSION.Category = strings.Split(category, "=")[1]
		t2.GetCookieHandler(w, r)

		tmpl1.Execute(w, t.TOPICSANDSESSION)
	} else {
		tmpl2.Execute(w, t.TOPICSANDSESSION)
	}
}
