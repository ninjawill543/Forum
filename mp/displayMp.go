package forum

import (
	"database/sql"
	"fmt"
	t "forum/users"
	"net/http"
	"strings"
)

type MpAndToWho struct {
	ToWho string
	Mps   []Mp `Mp`
}

type Mp struct {
	PrivateMessage string
	User1          string
	CreationDate   string
}

var MPSANDTOWHO MpAndToWho

func DisplayMp(r *http.Request, db *sql.DB) {
	fmt.Println(t.USER.Username)
	var message string
	var creationDate string
	var user1 string
	user2 := strings.Split(r.URL.Path, "/")
	MPSANDTOWHO.ToWho = user2[2]

	query := fmt.Sprintf("SELECT user1, message, creationDate FROM mp WHERE user1 = '%s' AND user2 = '%s' OR user1 = '%s' AND user2 = '%s' ORDER BY creationDate ASC", t.USER.Username, user2[2], user2[2], t.USER.Username)
	row, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	} else {
		MPSANDTOWHO.Mps = nil
		for row.Next() {
			defer row.Close()
			err := row.Scan(&user1, &message, &creationDate)
			if err != nil {
				fmt.Println(err)
			} else {
				mpIndex := len(MPSANDTOWHO.Mps)
				MPSANDTOWHO.Mps = append(MPSANDTOWHO.Mps, Mp{})
				MPSANDTOWHO.Mps[mpIndex].PrivateMessage = message
				MPSANDTOWHO.Mps[mpIndex].CreationDate = creationDate
				MPSANDTOWHO.Mps[mpIndex].User1 = user1
			}
		}
	}
}
