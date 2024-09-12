package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	CreateEventTable()
}

func CreateEventTable() {

	table := `CREATE TABLE IF NOT EXISTS events (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                name TEXT NOT NULL,
                description TEXT NOT NULL,
                location TEXT NOT NULL,
                datetime DATETIME NOT NULL,
                user_id INTEGER
              )`

	_, err := DB.Exec(table)

	if err != nil {
		panic("Could not create events table.")
	}
}
