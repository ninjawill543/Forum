package forum

import (
	"database/sql"
	"fmt"
	t "forum/structs"
)

func MessagesSendByUser(db *sql.DB) {
	var message string
	var uuidPath string
	var name string
	USER.MessagesSend = nil
	query := fmt.Sprintf("SELECT message, uuidPath FROM messages WHERE owner = '%s'", USER.Username)
	row, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	} else {
		for row.Next() {
			defer row.Close()
			err = row.Scan(&message, &uuidPath)
			if err != nil {
				fmt.Println(err)
			} else {
				indexMessagesSend := len(USER.MessagesSend)
				USER.MessagesSend = append(USER.MessagesSend, t.MessageSend{})
				USER.MessagesSend[indexMessagesSend].MessageSendByUser = message

				query2 := fmt.Sprintf("SELECT name FROM topics WHERE uuid = '%s'", uuidPath)
				row2, err := db.Query(query2)
				if err != nil {
					fmt.Println(err)
				} else {
					for row2.Next() {
						defer row2.Close()
						err = row2.Scan(&name)
						if err != nil {
							fmt.Println(err)
						}
					}
					USER.MessagesSend[indexMessagesSend].TopicSentInName = name
				}
			}
		}
	}
}
