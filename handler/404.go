package forum

import (
	"html/template"
	"net/http"
)

func Handler_404(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../static/html/404.html"))
	tmpl.Execute(w, "")
}
