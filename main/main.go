package main

import (
	"crypto/sha1"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"net/mail"
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

	//register
	if r.Method == "POST" {
		fmt.Println("New POST: ")
		var checkAll bool
		username := r.FormValue("input_username")
		password := r.FormValue("input_password")
		mail := r.FormValue("input_mail")
		creationDate := time.Now()
		birthDay := r.FormValue("input_birthDay")
		notifications := r.FormValue("input_notifications")

		if len(username) < 5 || len(username) > 14 {
			fmt.Println("invalid username")
			checkAll = true
		}

		if len(password) < 6 {
			fmt.Println("invalid password")
			checkAll = true
		}

		if checkMail(mail) == false {
			checkAll = true
		}

		if checkAll == false {
			addUsers(database, username, hash(password), mail, creationDate, birthDay, notifications)
		}
	}

	tmpl1.Execute(w, "")
}

func addUsers(db *sql.DB, username string, password string, email string, creationDate time.Time, birthDate string, notifications string) {
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

func checkMail(email string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}
}

func hash(password string) string {
	hash := sha1.New()
	hashInBytes := hash.Sum([]byte(password))[:20]
	return hex.EncodeToString(hashInBytes)
	//encoding passwords in sha1
}
