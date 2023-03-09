package forum

import (
	"database/sql"
	"fmt"
	t "forum/structs"
)

func MpSendOrReceivedByUser(db *sql.DB) {
	var message string
	var user2 string
	var user1 string
	USER.PrivateMessages = nil
	query := fmt.Sprintf("SELECT message, user1, user2 FROM mp WHERE user1 = '%s' OR user2 = '%s'", USER.Username, USER.Username)
	row, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	} else {
		for row.Next() {
			defer row.Close()
			err = row.Scan(&message, &user1, &user2)
			if err != nil {
				fmt.Println(err)
			} else {
				msgPrivateIndex := len(USER.PrivateMessages)
				USER.PrivateMessages = append(USER.PrivateMessages, t.PrivateMessage{})
				USER.PrivateMessages[msgPrivateIndex].PrivateMessage = message
				if user2 != USER.Username {
					USER.PrivateMessages[msgPrivateIndex].PrivateMessage2ndUser = user2
				} else {
					USER.PrivateMessages[msgPrivateIndex].PrivateMessage2ndUser = user1
				}
			}
		}
	}
}
