package forum

import (
	"database/sql"
	"fmt"
)

type Topics struct {
	Id           int
	Name         string
	Likes        string
	Dislikes     string
	CreationDate string
	Owner        string
	Uuid         int
}

var TOPICS []Topics

func DisplayTopic(db *sql.DB) {
	name := ""
	likes := ""
	dislikes := ""
	creationDate := ""
	owner := ""
	id := 0

	row, err := db.Query("SELECT id, name, creationDate, owner, likes, dislikes from topics;")
	if err != nil {
		fmt.Println(err)
	} else {
		for row.Next() {
			err = row.Scan(&id, &name, &creationDate, &owner, &likes, &dislikes)
			if err != nil {
				fmt.Println(nil)
			} else {
				topicIndex := len(TOPICS)

				if topicIndex == 0 || topicIndex == id-1 {
					TOPICS = append(TOPICS, Topics{})
					TOPICS[topicIndex].Id = id
					TOPICS[topicIndex].Name = name
					TOPICS[topicIndex].Likes = likes
					TOPICS[topicIndex].Dislikes = dislikes
					TOPICS[topicIndex].CreationDate = creationDate
					TOPICS[topicIndex].Owner = owner
				}
			}
		}
		row.Close()
	}
}
