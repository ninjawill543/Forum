package main

import (
	"database/sql"
	"fmt"
	t4 "forum/404"
	t3 "forum/handlerIndex"
	t2 "forum/topics"
	t "forum/users"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	http.HandleFunc("/", t4.Handler_404)
	http.HandleFunc("/index", t3.Handler_index)

	//creating DB if not exist
	databaseUsers, err := sql.Open("sqlite3", "../users.db")
	if err != nil {
		fmt.Println(err)
	}

	databaseTopics, err := sql.Open("sqlite3", "../topic.db")
	if err != nil {
		fmt.Println(err)
	}

	t.CreateTableUsers(databaseUsers)
	defer databaseUsers.Close()

	t2.CreateTableTopics(databaseTopics)
	defer databaseTopics.Close()

	//url of our funcs
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Print("Le Serveur dÃ©marre sur le port 8080\n")
	//listening on port 8080
	http.ListenAndServe(":8080", nil)
}
