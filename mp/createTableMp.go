package forum

import (
	"database/sql"
	"fmt"
)

func CreateTableMp(db *sql.DB) {
	mp_table := `CREATE TABLE mp(
		"user1" TEXT,
		"user2" TEXT,
		"creationDate" TEXT,
		"message" TEXT);`

	query, err := db.Prepare(mp_table)
	if err != nil {
		fmt.Println(err)
	} else {
		query.Exec()
		fmt.Println("table mp created successfully")
	}
}
