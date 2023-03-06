package forum

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func AddFirstMessageInMessages(firstMessage string, creationDate time.Time, owner string, uuidPath uuid.UUID, db *sql.DB) {
	uuid := uuid.New()
	addInMessages := `INSERT INTO messages(message, creationDate, owner, report, uuidPath, like, edited, uuid) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	query, err := db.Prepare(addInMessages)
	if err != nil {
		fmt.Println(err)
	}

	_, err = query.Exec(firstMessage, creationDate, owner, 0, uuidPath, 0, 0, uuid)
	if err != nil {
		fmt.Println(err)
	}
}
