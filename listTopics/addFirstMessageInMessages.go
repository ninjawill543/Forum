package forum

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

func AddFirstMessageInMessages(firstMessage string, creationDate string, owner string, uuidPath uuid.UUID, db *sql.DB) {
	//this is the firt message posted by owner of topic (optional)
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
