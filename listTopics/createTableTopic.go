package forum

import (
	"database/sql"
	"fmt"
)

func CreateTableTopics(db *sql.DB) {
	topic_table := `CREATE TABLE topics(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"name" TEXT,
		"firstMessage" TEXT,
		"creationDate" TEXT,
		"owner" TEXT,
		"likes" INTEGER,
		"nmbPosts" INTEGER,
		"lastPost" TEXT,
		"uuid" TEXT);`

	query, err := db.Prepare(topic_table)

	if err != nil {
		fmt.Println(err)
	} else {
		query.Exec()
		fmt.Println("Table topics created successfully")
	}
}
