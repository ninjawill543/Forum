package forum

import (
	"database/sql"
	"fmt"
)

func CreateTableReports(db *sql.DB) {
	reports_table := `CREATE TABLE reports(
		"uuidUser" TEXT,
		"uuidReported" TEXT);`

	query, err := db.Prepare(reports_table)
	if err != nil {
		fmt.Println(err)
	} else {
		query.Exec()
		fmt.Println("table reports created successfully")
	}
}
