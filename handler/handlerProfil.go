package forum

import (
	"database/sql"
	t "forum/users"
	"html/template"
	"net/http"
)

func Handler_profil(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../static/html/profil.html"))
	databaseUsers, _ := sql.Open("sqlite3", "../users.db")

	if r.FormValue("logOutButton") == "logout" {
		t.Logout(r)
	}

	if r.FormValue("username") != "" || r.FormValue("email") != "" || r.FormValue("password") != "" {
		t.UserEdit(r, databaseUsers)
	}

	tmpl.Execute(w, t.USER)
}
