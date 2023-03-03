package forum

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/404", Handler_404)
}

func Handler_404(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../static/html/notfound.html"))
	tmpl.Execute(w, "")
}
