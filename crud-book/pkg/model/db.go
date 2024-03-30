package model

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectDB() error {
	connStr := "postgres://root:Root@:123@localhost/book_db?sslmode=disable" // Replace with your PostgreSQL connection details
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	fmt.Println("Connected to the database")
	return nil
}
