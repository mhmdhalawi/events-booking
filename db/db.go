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

	CreateTables()
}

func CreateTables() {

	users := `CREATE TABLE IF NOT EXISTS users (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				email TEXT NOT NULL UNIQUE,
				password TEXT NOT NULL
			)`

	_, err := DB.Exec(users)

	if err != nil {
		panic(err)
	}

	events := `CREATE TABLE IF NOT EXISTS events (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                name TEXT NOT NULL,
                description TEXT NOT NULL,
                location TEXT NOT NULL,
                datetime DATETIME NOT NULL,
                user_id INTEGER,
				FOREIGN KEY (user_id) REFERENCES users(id)
              )`

	_, err = DB.Exec(events)

	if err != nil {
		panic(err)
	}

	registrations := `CREATE TABLE IF NOT EXISTS registrations (
						id INTEGER PRIMARY KEY AUTOINCREMENT,
						event_id INTEGER,
						user_id INTEGER,
						FOREIGN KEY (event_id) REFERENCES events(id),
						FOREIGN KEY (user_id) REFERENCES users(id)
					)`
	_, err = DB.Exec(registrations)

	if err != nil {
		panic(err)
	}
}
