package forum

import (
	"database/sql"
	t2 "forum/listTopics"
	t5 "forum/messages"
	t3 "forum/profil"
	t "forum/users"
	"log"
)

func CreatingDatabases() {
	databaseUsers, err := sql.Open("sqlite3", "../users.db")
	if err != nil {
		log.Fatal(err)
	}
	databaseTopics, err := sql.Open("sqlite3", "../topics.db")
	if err != nil {
		log.Fatal(err)
	}
	databaseMessages, err := sql.Open("sqlite3", "../messages.db")
	if err != nil {
		log.Fatal(err)
	}
	databaseLikesFromUser, err := sql.Open("sqlite3", "../likesFromUser.db")
	if err != nil {
		log.Fatal(err)
	}
	t.CreateTableUsers(databaseUsers)
	defer databaseUsers.Close()
	t2.CreateTableTopics(databaseTopics)
	defer databaseTopics.Close()
	t5.CreateTableMessage(databaseMessages)
	defer databaseMessages.Close()
	t3.CreateTableLikesFromUser(databaseLikesFromUser)
	defer databaseLikesFromUser.Close()
}
