package forum

import (
	"html/template"
	"net/http"
)

func Handler_topicPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../static/html/topicPage.html"))
	tmpl.Execute(w, "")
}
