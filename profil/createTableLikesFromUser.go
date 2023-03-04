package forum

import (
	"database/sql"
	"fmt"
)

func CreateTableLikesFromUser(db *sql.DB) {
	likesFromUser_table := `CREATE TABLE likesFromUser(
		"uuidUser" TEXT,
		"uuidLiked" TEXT);`

	query, err := db.Prepare(likesFromUser_table)
	if err != nil {
		fmt.Println(err)
	} else {
		query.Exec()
		fmt.Println("table likesFromUser created successfully")
	}
}
