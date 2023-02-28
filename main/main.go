package main

import (
	//"database/sql"
	"fmt"
	t3 "forum/handlerIndex"
	t2 "forum/topics"
	t "forum/users"
	"net/http"

	//_ "github.com/mattn/go-sqlite3"
)

func main() {
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
	fs := http.FileServer(http.Dir("../static/css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	fmt.Print("Le Serveur dÃ©marre sur le port 8080\n")
	//listening on port 8080
	http.ListenAndServe(":8080", nil)
}
