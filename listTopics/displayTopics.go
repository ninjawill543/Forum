package forum

import (
	"database/sql"
	"fmt"
)

type Topics struct {
	Id           int
	Name         string
	Likes        int
	Dislikes     int
	CreationDate string
	Owner        string
	Uuid         string
}

var TOPICS []Topics

func DisplayTopic(db *sql.DB) {
	var name string
	var likes int
	var dislikes int
	var creationDate string
	var owner string
	var id int
	var uuid string

	row, err := db.Query("SELECT id, name, creationDate, owner, likes, dislikes, uuid from topics;")
	if err != nil {
		fmt.Println(err)
	} else {
		TOPICS = nil
		for row.Next() {
			err = row.Scan(&id, &name, &creationDate, &owner, &likes, &dislikes, &uuid)
			if err != nil {
				fmt.Println(nil)
			} else {
				topicIndex := len(TOPICS)

				TOPICS = append(TOPICS, Topics{})
				TOPICS[topicIndex].Id = id
				TOPICS[topicIndex].Name = name
				TOPICS[topicIndex].Likes = likes
				TOPICS[topicIndex].Dislikes = dislikes
				TOPICS[topicIndex].CreationDate = creationDate
				TOPICS[topicIndex].Owner = owner
				TOPICS[topicIndex].Uuid = uuid
			}
		}
		row.Close()
	}
}
