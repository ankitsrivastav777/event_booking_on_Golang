package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	// Open a connection to the SQLite database
	var err error
	db, err := sql.Open("sqlite", "events.db")
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	db.SetMaxOpenConns(10) // Set the maximum number of open connections to the database
	db.SetMaxIdleConns(5)  // Set the maximum number of idle connections to
	DB = db                // Assign the database connection to the global variable

	CreateTables() // Create the necessary tables if they don't exist
}

func CreateTables() {
	// Create the events table if it doesn't exist
	CreateTables := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		datetime DATETIME NOT NULL,
		user_id INTEGER
	);
	`
	_, err := DB.Exec(CreateTables)
	if err != nil {
		panic("Failed to create tables: " + err.Error())
	}
}
