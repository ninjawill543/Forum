package forum

import (
	"database/sql"
	"fmt"
	"net/http"
)

type TopicsAndSession struct {
	SessionUser string
	Topics      []Topic `Topic`
}

type Topic struct {
	Id           int
	Name         string
	Likes        int
	CreationDate string
	Owner        string
	Uuid         string
	NmbPosts     int
	FirstMessage string
	// LastPost     string
}

var TOPICSANDSESSION TopicsAndSession

func DisplayTopic(r *http.Request, db *sql.DB) {
	var name string
	var likes int
	var creationDate string
	var owner string
	var uuid string
	var nmbPosts int
	var filter string
	var id int
	var searchName string
	var query string
	var firstMessage string
	var username string
	// var lastPost string

	filter = r.FormValue("filter")

	if filter == "" {
		filter = "creationDate"
	}

	query = fmt.Sprintf("SELECT id, name, firstMessage, creationDate, owner, likes, nmbPosts, uuid FROM topics ORDER BY %s DESC", filter)

	if r.FormValue("searchbar") != "" {
		searchName = "%" + r.FormValue("searchbar") + "%"
		query = fmt.Sprintf("SELECT id, name, firstMessage, creationDate, owner, likes, nmbPosts, uuid FROM topics WHERE name LIKE '%s' ORDER BY %s DESC", searchName, filter)
	}

	cookie, err := r.Cookie("session")
	if err != nil {
		// fmt.Println(err)
		TOPICSANDSESSION.SessionUser = ""
	} else {
		queryGetName := fmt.Sprintf("SELECT username FROM users WHERE uuid = '%s'", cookie.Value)
		row, err := db.Query(queryGetName)
		if err != nil {
			fmt.Println(err)
		} else {
			for row.Next() {
				defer row.Close()
				row.Scan(&username)
			}
			TOPICSANDSESSION.SessionUser = username
		}
	}

	row2, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	} else {
		TOPICSANDSESSION.Topics = nil
		for row2.Next() {
			defer row2.Close()
			err = row2.Scan(&id, &name, &firstMessage, &creationDate, &owner, &likes, &nmbPosts, &uuid)
			if err != nil {
				fmt.Println(err)
			} else {
				topicIndex := len(TOPICSANDSESSION.Topics)
				TOPICSANDSESSION.Topics = append(TOPICSANDSESSION.Topics, Topic{})
				TOPICSANDSESSION.Topics[topicIndex].Name = name
				TOPICSANDSESSION.Topics[topicIndex].Likes = likes
				TOPICSANDSESSION.Topics[topicIndex].CreationDate = creationDate
				TOPICSANDSESSION.Topics[topicIndex].Owner = owner
				TOPICSANDSESSION.Topics[topicIndex].NmbPosts = nmbPosts
				TOPICSANDSESSION.Topics[topicIndex].Uuid = uuid
				TOPICSANDSESSION.Topics[topicIndex].Id = id
				TOPICSANDSESSION.Topics[topicIndex].FirstMessage = firstMessage
				// TOPICSANDSESSION.Topics[topicIndex].LastPost = lastPost
			}
		}
	}
}
