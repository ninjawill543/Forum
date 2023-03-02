package forum

import (
	"database/sql"
	"fmt"
)

func CreateTableTopics(db *sql.DB) {
	topic_table := `CREATE TABLE topics(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"name" TEXT,
		"creationDate" TEXT,
		"owner" TEXT,
		"likes" INTEGER,
		"dislikes" INTEGER,
		"uuid" INTEGER);`

	query, err := db.Prepare(topic_table)

	if err != nil {
		fmt.Println(err)
	} else {
		query.Exec()
		fmt.Println("Table Topic created successfully")
	}
}
