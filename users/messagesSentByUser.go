package forum

import (
	"database/sql"
	"fmt"
)

func MessagesSendByUser(db *sql.DB) {
	var message string
	USER.MessagesSend = nil
	query := fmt.Sprintf("SELECT message FROM messages WHERE owner = '%s'", USER.Username)
	row, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	} else {
		for row.Next() {
			err = row.Scan(&message)
			if err != nil {
				fmt.Println(err)
			} else {
				USER.MessagesSend = append(USER.MessagesSend, message)
			}
		}
	}
}
