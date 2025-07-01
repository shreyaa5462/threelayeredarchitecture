package taskds

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (*sql.DB, error) {
	dsn := "root:shreya123@tcp(localhost:3306)/shreya"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS mytask (
		id INT AUTO_INCREMENT PRIMARY KEY,
		description TEXT NOT NULL,
		completed BOOLEAN DEFAULT FALSE
		)`
	_, err = db.Exec(createTableQuery)

	return db, err
}
