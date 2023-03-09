package forum

import (
	"database/sql"
	"fmt"
	t "forum/users"
	"net/http"
	"strings"
)

type Mp struct {
	PrivateMessage string
	User1          string
	CreationDate   string
}

var MPS []Mp

func DisplayMp(r *http.Request, db *sql.DB) {
	fmt.Println(t.USER.Username)
	var message string
	var creationDate string
	var user1 string
	user2 := strings.Split(r.URL.Path, "/")
	query := fmt.Sprintf("SELECT user1, message, creationDate FROM mp WHERE user1 = '%s' AND user2 = '%s' OR user1 = '%s' AND user2 = '%s' ORDER BY creationDate DESC", t.USER.Username, user2[2], user2[2], t.USER.Username)
	fmt.Println(query)
	row, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	} else {
		MPS = nil
		for row.Next() {
			defer row.Close()
			err := row.Scan(&user1, &message, &creationDate)
			if err != nil {
				fmt.Println(err)
			} else {
				mpIndex := len(MPS)
				MPS = append(MPS, Mp{})
				MPS[mpIndex].PrivateMessage = message
				MPS[mpIndex].CreationDate = creationDate
				MPS[mpIndex].User1 = user1
			}
		}
	}
}
