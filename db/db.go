package db

import (
	"database/sql"

	_ "modernc.org/sqlite" // means we are importing for side-effects only
)

var DB *sql.DB // capitalized variable is exported

// InitDB initializes the database connection
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "api.db")

	if err != nil {
		panic("Cannot connect to database: " + err.Error()) // panic is used when we cannot recover from an error
	}

	DB.SetMaxOpenConns(10) // set maximum number of open connections to the database
	DB.SetMaxIdleConns(5)  // set maximum number of idle connections in the pool

	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time DATETIME NOT NULL,
		user_id INTEGER
	);`

	_, err := DB.Exec(createEventsTable)

	if err != nil {
		panic("Failed to create events table: " + err.Error())
	}
}
