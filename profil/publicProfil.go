package forum

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
)

type PublicUser struct {
	Username      string
	Email         string
	CreationDate  string
	BirthDate     string
	Uuid          string
	Admin         string
	Reports       int
	TopicsCreated []string
	MessagesSend  []string
}

var PUBLICUSER PublicUser

func PublicProfil(r *http.Request, dbUsers *sql.DB, dbMessages *sql.DB, dbTopics *sql.DB) {
	var username string
	var creationDate string
	var admin string
	var birthDate string
	var reports int
	namePublic := strings.Split(r.URL.Path, "/")
	query := fmt.Sprintf("SELECT username, creationDate, admin, birthDate, reports FROM users WHERE username = '%s'", namePublic[2])
	row, err := dbUsers.Query(query)
	if err != nil {
		fmt.Println(err)
	} else {
		for row.Next() {
			err = row.Scan(&username, &creationDate, &admin, &birthDate, &reports)
			if err != nil {
				fmt.Println(err)
			}
		}
		PUBLICUSER.Username = username
		PUBLICUSER.CreationDate = creationDate
		PUBLICUSER.Admin = admin
		PUBLICUSER.BirthDate = birthDate
		PUBLICUSER.Reports = reports
	}
	var message string
	PUBLICUSER.MessagesSend = nil
	query = fmt.Sprintf("SELECT message FROM messages WHERE owner = '%s'", PUBLICUSER.Username)
	row, err = dbMessages.Query(query)
	if err != nil {
		fmt.Println(err)
	} else {
		for row.Next() {
			err = row.Scan(&message)
			if err != nil {
				fmt.Println(err)
			} else {
				PUBLICUSER.MessagesSend = append(PUBLICUSER.MessagesSend, message)
			}
		}
	}

	var name string
	PUBLICUSER.TopicsCreated = nil
	query = fmt.Sprintf("SELECT name FROM topics WHERE owner = '%s'", PUBLICUSER.Username)
	row, err = dbTopics.Query(query)
	if err != nil {
		fmt.Println(err)
	} else {
		for row.Next() {
			err = row.Scan(&name)
			if err != nil {
				fmt.Println(err)
			} else {
				PUBLICUSER.TopicsCreated = append(PUBLICUSER.TopicsCreated, name)
			}
		}
	}
}
