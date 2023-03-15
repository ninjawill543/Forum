package forum

import (
	"database/sql"
	"fmt"
	t3 "forum/delete"
	t "forum/listTopics"
	t4 "forum/login"
	t2 "forum/users"
	"html/template"
	"net/http"
)

func Handler_topics(w http.ResponseWriter, r *http.Request) {

	tmpl1 := template.Must(template.ParseFiles("../static/html/topics.html"))

	databaseForum, _ := sql.Open("sqlite3", "../forum.db")

	t2.GetCookieHandler(w, r)

	if r.FormValue("topic_name") != "" {
		t.AddTopic(r, databaseForum)
	}

	if r.FormValue("input_mail") != "" {
		t2.EmailStorage(r)
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

	if r.URL.Path != "/topics/category=blabla" {
		fmt.Println("this category doesnt exist")
	} else {
		t.DisplayTopic(r, databaseForum)
		tmpl1.Execute(w, t.TOPICSANDSESSION)
	}
}
