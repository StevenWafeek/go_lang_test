package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "./api.db")
	if err != nil {
		panic("Failed to connect to DB: " + err.Error())
	}

	CreateEventTables()
	fmt.Println("✅ Database initialized and tables created")
}

func CreateEventTables() {

	CreateUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(CreateUserTable)
	if err != nil {
		panic("Failed to create table: " + err.Error())
	}

	CreateEventTables := `
	CREATE TABLE IF NOT EXISTS events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    location TEXT NOT NULL,
    userID INTEGER,
	FOREIGN KEY (userID) REFERENCES users(id)
)

	`
	_, err = DB.Exec(CreateEventTables)
	if err != nil {
		panic("Failed to create table: " + err.Error())
	}
}
