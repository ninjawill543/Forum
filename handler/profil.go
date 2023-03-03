package forum

import (
	t "forum/users"
	"html/template"
	"net/http"
)

func Handler_profil(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../static/html/profil.html"))

	if r.FormValue("logOutButton") == "logout" {
		t.Logout(r)
	}

	tmpl.Execute(w, t.USER)
}
