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

	http.HandleFunc("/404", t3.Handler_404)
	http.HandleFunc("/topics/", t3.Handler_topics)
	http.HandleFunc("/messages/", t3.Handler_Messages)
	http.HandleFunc("/your-profil/", t3.Handler_profil)
	http.HandleFunc("/profil/", t3.Handler_publicProfil)
	http.HandleFunc("/edit-topic/", t3.Handler_EditTopic)
	http.HandleFunc("/edit-message/", t3.Handler_EditMessage)
	http.HandleFunc("/private-message/", t3.Handler_Mp)
	http.HandleFunc("/", t3.Handler_Home)

	//url of our funcs
	fs := http.FileServer(http.Dir("../static/css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	fmt.Print("Le Serveur dÃ©marre sur le port 8080\n")
	//listening on port 8080
	http.ListenAndServe(":8080", nil)
}
