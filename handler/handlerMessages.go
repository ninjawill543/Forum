package forum

import (
	"database/sql"
	"fmt"
	t3 "forum/delete"
	t5 "forum/login"
	t "forum/messages"
	t2 "forum/report"
	t4 "forum/users"
	"html/template"
	"net/http"
	"strings"
)

func Handler_Messages(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../static/html/messages.html"))
	tmpl2 := template.Must(template.ParseFiles("../static/html/404.html"))

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

	query := `SELECT name FROM topics`

	var name string
	var exists bool
	urlName := strings.Split(r.URL.Path, "/")
	newUrlName := strings.TrimSpace(urlName[2])

	row, err := databaseForum.Query(query)
	if err != nil {
		fmt.Println(err)
	} else {
		for row.Next() {
			defer row.Close()
			err = row.Scan(&name)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(name)
				fmt.Println(newUrlName)
				if name == string(newUrlName) {
					exists = true
				}
			}
		}
	}

	if exists {
		t.Messages.SessionUser = t4.USER.Username
		tmpl.Execute(w, t.Messages)
	} else {
		tmpl2.Execute(w, nil)
	}
}
