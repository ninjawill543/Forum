package forum

import (
	"database/sql"
	"fmt"
	t "forum/views"
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
	Ban           int
	TopicsCreated []string
	MessagesSend  []MessageSend `MessageSend`
}

type MessageSend struct {
	MessageSendByUser string
	TopicSentInName   string
}

var PUBLICUSER PublicUser

func PublicProfil(r *http.Request, db *sql.DB) {
	var username string
	var creationDate string
	var admin string
	var birthDate string
	var reports int
	var ban int
	var uuidPath string
	var name string
	namePublic := strings.Split(r.URL.Path, "/")
	query := fmt.Sprintf("SELECT username, creationDate, admin, birthDate, reports, ban FROM users WHERE username = '%s'", namePublic[2])
	row, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	} else {
		for row.Next() {
			defer row.Close()
			err = row.Scan(&username, &creationDate, &admin, &birthDate, &reports, &ban)
			if err != nil {
				fmt.Println(err)
			}
		}
		PUBLICUSER.Username = username
		PUBLICUSER.CreationDate = t.DisplayTime(creationDate, " ")
		PUBLICUSER.Admin = admin
		PUBLICUSER.BirthDate = birthDate
		PUBLICUSER.Reports = reports
		PUBLICUSER.Ban = ban
	}
	var message string
	PUBLICUSER.MessagesSend = nil
	query = fmt.Sprintf("SELECT message, uuidPath FROM messages WHERE owner = '%s'", PUBLICUSER.Username)
	fmt.Println("test messagesSent", PUBLICUSER.MessagesSend)
	row2, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	} else {
		for row2.Next() {
			defer row2.Close()
			err = row2.Scan(&message, &uuidPath)
			if err != nil {
				fmt.Println(err)
			} else {
				indexMessagesSend := len(PUBLICUSER.MessagesSend)
				PUBLICUSER.MessagesSend = append(PUBLICUSER.MessagesSend, MessageSend{})
				PUBLICUSER.MessagesSend[indexMessagesSend].MessageSendByUser = message

				query2 := fmt.Sprintf("SELECT name FROM topics WHERE uuid = '%s'", uuidPath)
				row3, err := db.Query(query2)
				if err != nil {
					fmt.Println(err)
				} else {
					defer row3.Close()
					for row3.Next() {
						err = row3.Scan(&name)
						if err != nil {
							fmt.Println(err)
						}
					}
					PUBLICUSER.MessagesSend[indexMessagesSend].TopicSentInName = name
				}
			}
		}
	}

	PUBLICUSER.TopicsCreated = nil
	query = fmt.Sprintf("SELECT name FROM topics WHERE owner = '%s'", PUBLICUSER.Username)
	row4, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	} else {
		for row4.Next() {
			defer row4.Close()
			err = row4.Scan(&name)
			if err != nil {
				fmt.Println(err)
			} else {
				PUBLICUSER.TopicsCreated = append(PUBLICUSER.TopicsCreated, name)
			}
		}
	}
}
