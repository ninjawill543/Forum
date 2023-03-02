package forum

import (
	"database/sql"
	t2 "forum/listTopics"
	t5 "forum/topic"
	t "forum/users"
	"log"
)

func CreatingDatabases() {
	databaseUsers, err := sql.Open("sqlite3", "../users.db")
	if err != nil {
		log.Fatal(err)
	}

	databaseTopics, err := sql.Open("sqlite3", "../topic.db")
	if err != nil {
		log.Fatal(err)
	}

	databaseMessages, err := sql.Open("sqlite3", "../messages.db")
	if err != nil {
		log.Fatal(err)
	}

	t.CreateTableUsers(databaseUsers)
	defer databaseUsers.Close()

	t2.CreateTableTopics(databaseTopics)
	defer databaseTopics.Close()

	t5.CreateTableMessage(databaseMessages)
	defer databaseMessages.Close()
}
