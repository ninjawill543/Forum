package forum

import (
	"database/sql"
	"fmt"
)

func CreateTableUsers(db *sql.DB) {
	users_table := `CREATE TABLE users(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"username" TEXT,
		"password" TEXT,
		"email" TEXT,
		"creationDate" TEXT,
		"birthDate" TEXT,
		"uuid" TEXT);`

	query, err := db.Prepare(users_table)

	if err != nil {
		fmt.Println(err)
	} else {
		query.Exec()
		fmt.Println("Table Users created successfully")
	}
}
