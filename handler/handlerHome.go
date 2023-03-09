package forum

import (
	"html/template"
	"net/http"
)

func Handler_Home(w http.ResponseWriter, r *http.Request) {
	tmpl1 := template.Must(template.ParseFiles("../static/html/home.html"))
	tmpl1.Execute(w, nil)
}
