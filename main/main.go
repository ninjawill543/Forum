package main

import (
	"fmt"
	t1 "forum/databases"
	t3 "forum/handler"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	t1.CreatingDatabases()

	http.HandleFunc("/", t3.Handler_404)
	http.HandleFunc("/index", t3.Handler_index)
	http.HandleFunc("/topics/", t3.Handler_topicPage)
	http.HandleFunc("/profil/", t3.Handler_profil)
	http.HandleFunc("/edit-topic/", t3.Handler_EditTopic)
	http.HandleFunc("/edit-message/", t3.Handler_EditMessage)

	//url of our funcs
	fs := http.FileServer(http.Dir("../static/css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	fmt.Print("Le Serveur dÃ©marre sur le port 8080\n")
	//listening on port 8080
	http.ListenAndServe(":8080", nil)
}
