package forum

import (
	"database/sql"
	"fmt"
)

func CreateTableMessage(db *sql.DB) {
	messages_table := `CREATE TABLE messages(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"message" TEXT,
		"creationDate" TEXT,
		"owner" TEXT,
		"report" INTEGER,
		"uuid" TEXT);`

	query, err := db.Prepare(messages_table)

	if err != nil {
		fmt.Println(err)
	} else {
		query.Exec()
		fmt.Println("Table Messages created successfully")
	}
}
