package forum

import (
	"database/sql"
	"fmt"
)

func MessagesSendByUser(db *sql.DB) {
	var message string
	var uuidPath string
	var name string
	USER.MessagesSend = nil
	query := fmt.Sprintf("SELECT message, uuidPath FROM messages WHERE owner = '%s'", USER.Username)
	fmt.Println("test messagesSent", USER.MessagesSend)
	row, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	} else {
		for row.Next() {
			err = row.Scan(&message, &uuidPath)
			if err != nil {
				fmt.Println(err)
			} else {
				indexMessagesSend := len(USER.MessagesSend)
				USER.MessagesSend = append(USER.MessagesSend, MessageSend{})
				USER.MessagesSend[indexMessagesSend].MessageSendByUser = message

				query2 := fmt.Sprintf("SELECT name FROM topics WHERE uuid = '%s'", uuidPath)
				row, err = db.Query(query2)
				if err != nil {
					fmt.Println(err)
				} else {
					for row.Next() {
						err = row.Scan(&name)
						if err != nil {
							fmt.Println(err)
						}
					}
					row.Close()
					USER.MessagesSend[indexMessagesSend].TopicSentInName = name
				}
			}
		}
		row.Close()
	}
}
