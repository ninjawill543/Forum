package forum

import (
	"database/sql"
	"fmt"
	"net/http"
)

type Topics struct {
	Id           int
	Name         string
	Likes        int
	CreationDate string
	Owner        string
	Uuid         string
	NmbPosts     int
}

var TOPICS []Topics

func DisplayTopic(r *http.Request, db *sql.DB) {
	var name string
	var likes int
	var creationDate string
	var owner string
	var uuid string
	var nmbPosts int
	var filter string
	var id int

	filter = r.FormValue("filter")
	if filter == "" {
		filter = "lastPost"
	}
	query := fmt.Sprintf("SELECT id, name, creationDate, owner, likes, nmbPosts, uuid FROM topics ORDER BY %s DESC", filter)
	row, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	} else {
		TOPICS = nil
		for row.Next() {
			err = row.Scan(&id, &name, &creationDate, &owner, &likes, &nmbPosts, &uuid)
			if err != nil {
				fmt.Println(nil)
			} else {
				topicIndex := len(TOPICS)

				TOPICS = append(TOPICS, Topics{})
				TOPICS[topicIndex].Name = name
				TOPICS[topicIndex].Likes = likes
				TOPICS[topicIndex].CreationDate = creationDate
				TOPICS[topicIndex].Owner = owner
				TOPICS[topicIndex].NmbPosts = nmbPosts
				TOPICS[topicIndex].Uuid = uuid
				TOPICS[topicIndex].Id = id
			}
		}
		row.Close()
	}
}
