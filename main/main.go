package forum

import (
	"fmt"
	"http"
	"net/http"
)

func main() {
	http.HandleFunc("/", Handler_index)
	//url of our funcs
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Print("Le Serveur dÃ©marre sur le port 8080\n")
	http.ListenAndServe(":8080", nil)
	//listening on port 8080
}

func Handler_index() {
	//tmpl1 := template.Must(template.ParseFiles("./static/index.html"))
	//tmpl1.Execute()
}
