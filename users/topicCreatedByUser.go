package forum

import (
	"database/sql"
	"fmt"
)

func TopicCreatedByUser(db *sql.DB) {
	var name string
	USER.TopicsCreated = nil
	query := fmt.Sprintf("SELECT name FROM topics WHERE owner = '%s'", USER.Username)
	row, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	} else {
		for row.Next() {
			err = row.Scan(&name)
			if err != nil {
				fmt.Println(err)
			} else {
				USER.TopicsCreated = append(USER.TopicsCreated, name)
			}
		}
	}
}
