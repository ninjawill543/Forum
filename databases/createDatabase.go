package forum

import (
	"database/sql"
	t2 "forum/listTopics"
	t5 "forum/messages"
	t6 "forum/mp"
	t3 "forum/profil"
	t4 "forum/report"
	t "forum/users"
	"log"
)

func CreatingDatabases() {
	databaseForum, err := sql.Open("sqlite3", "../forum.db")
	if err != nil {
		log.Fatal(err)
	} else {
		t.CreateTableUsers(databaseForum)
		t2.CreateTableTopics(databaseForum)
		t5.CreateTableMessage(databaseForum)
		t3.CreateTableLikesFromUser(databaseForum)
		t4.CreateTableReports(databaseForum)
		t6.CreateTableMp(databaseForum)
	}
	// databaseUsers, err := sql.Open("sqlite3", "../users.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// databaseTopics, err := sql.Open("sqlite3", "../topics.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// databaseMessages, err := sql.Open("sqlite3", "../messages.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// databaseLikesFromUser, err := sql.Open("sqlite3", "../likesFromUser.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// databaseReports, err := sql.Open("sqlite3", "../reports.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
