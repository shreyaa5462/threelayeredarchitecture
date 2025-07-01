package userds

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func InitUserDB() (*sql.DB, error) {
	dsn := "root:shreya123@tcp(localhost:3306)/shreya"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100) NOT NULL
	)`
	_, err = db.Exec(createTableQuery)

	return db, err
}
