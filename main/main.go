package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	http.HandleFunc("/", Handler_index)

	database, err := sql.Open("sqlite3", "../forum.db")
	if err != nil {
		fmt.Println(err)
	}
	createTable(database)
	defer database.Close()

	//url of our funcs
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Print("Le Serveur dÃ©marre sur le port 8080\n")
	//listening on port 8080
	http.ListenAndServe(":8080", nil)
}

func Handler_index(w http.ResponseWriter, r *http.Request) {
	database, _ := sql.Open("sqlite3", "../forum.db")
	tmpl1 := template.Must(template.ParseFiles("../static/index.html"))

	if r.Method == "POST" {
		username := r.FormValue("input_username")
		password := r.FormValue("input_password")
		mail := r.FormValue("input_mail")
		theTime := time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local)
		creationDate := theTime.Format("2006-1-2 15:4:5")
		birthDay := r.FormValue("input_birthDay")
		notifications := r.FormValue("input_notifications")

		addUsers(database, username, password, mail, creationDate, birthDay, notifications)
	}

	tmpl1.Execute(w, "")
}

func addUsers(db *sql.DB, username string, password string, email string, creationDate string, birthDate string, notifications string) {
	usersInfo := `INSERT INTO users(username, password, email, creationDate, birthDate, notifications) VALUES (?, ?, ?, ?, ?, ?)`
	query, err := db.Prepare(usersInfo)
	if err != nil {
		log.Fatal(err)
	}
	_, err = query.Exec(username, password, email, creationDate, birthDate, notifications)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("adding new user :", username, "in users")
	}
}

func createTable(db *sql.DB) {
	users_table := `CREATE TABLE users(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"username" TEXT,
		"password" TEXT,
		"email" TEXT,
		"creationDate" TEXT,
		"birthDate" TEXT,
		"notifications" TEXT);`

	query, err := db.Prepare(users_table)

	if err != nil {
		fmt.Println(err)
	} else {
		query.Exec()
		fmt.Println("Table created successfully")
	}
}
