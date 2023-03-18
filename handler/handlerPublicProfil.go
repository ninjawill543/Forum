package forum

import (
	"database/sql"
	"fmt"
	t3 "forum/mp"
	t "forum/profil"
	t2 "forum/report"
	"html/template"
	"net/http"
	"strings"
)

func Handler_publicProfil(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../static/html/publicProfil.html"))
	tmpl2 := template.Must(template.ParseFiles("../static/html/404.html"))
	databaseForum, _ := sql.Open("sqlite3", "../forum.db")

	if r.FormValue("ban") != "" {
		t2.Ban(r, databaseForum)
	} else if r.FormValue("report") != "" {
		t2.ReportUser(r, databaseForum)
	}

	if r.FormValue("mpMessage") != "" {
		t3.AddMp(r, databaseForum)

	}

	query := `SELECT username FROM users`

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
		t.PublicProfil(r, databaseForum)
		tmpl.Execute(w, t.PUBLICUSER)
	} else {
		tmpl2.Execute(w, nil)
	}

}
