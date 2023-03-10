package forum

import (
	"database/sql"
	"fmt"
	t "forum/structs"
)

var TOPICSANDSESSION t.TopicsAndSession

func GetRandomMessages(db *sql.DB) {
	var name string
	var firstMessage string
	var creationDate string
	var owner string
	var likes int
	// var nmbPosts int
	var category string
	var allTopics []string
	var alreadyUsed bool

	TOPICSANDSESSION.Topics = nil
	allTopics = nil

	for i := 0; i < 10; i++ {
		query := "SELECT name, firstMessage, creationDate, owner, likes, category FROM topics ORDER BY RANDOM() LIMIT 1"
		row, err := db.Query(query)
		if err != nil {
			fmt.Println(err)
		} else {
			for row.Next() {
				defer row.Close()
				row.Scan(&name, &firstMessage, &creationDate, &owner, &likes, &category)
			}
			alreadyUsed = false
			for i := 0; i < len(allTopics); i++ {
				if allTopics[i] == name {
					alreadyUsed = true
				}
			}
			if !alreadyUsed {
				allTopics = append(allTopics, name)
				topicIndex := len(TOPICSANDSESSION.Topics)
				TOPICSANDSESSION.Topics = append(TOPICSANDSESSION.Topics, t.Topic{})
				TOPICSANDSESSION.Topics[topicIndex].Name = name
				TOPICSANDSESSION.Topics[topicIndex].Likes = likes
				TOPICSANDSESSION.Topics[topicIndex].CreationDate = creationDate
				TOPICSANDSESSION.Topics[topicIndex].Owner = owner
				TOPICSANDSESSION.Topics[topicIndex].FirstMessage = firstMessage
				TOPICSANDSESSION.Topics[topicIndex].Category = category
				// TOPICSANDSESSION.Topics[topicIndex].NmbPosts = nmbPosts
			}
		}
	}
}
