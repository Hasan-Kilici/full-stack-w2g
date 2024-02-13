package database

import (
	"database/sql"
	"log"
	"time"

	"github.com/patrickmn/go-cache"
	_ "modernc.org/sqlite"
)

var (
	db *sql.DB
	userCache *cache.Cache
)

func init() {
	var err error
	userCache = cache.New(5*time.Minute, 10*time.Minute)
	
	db, err = sql.Open("sqlite", "database/_database.db")
	if err != nil {
		log.Fatal(err)
	}
}

func createTable(query string) {
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateTables() {
	createTable(`
		CREATE TABLE IF NOT EXISTS rooms (
			id 			INTEGER PRIMARY KEY AUTOINCREMENT,
			name 		TEXT NOT NULL,
			description TEXT,
			public 		BOOLEAN
		);
	`)

	createTable(`
		CREATE TABLE IF NOT EXISTS roomsettings (
			roomID 			INTEGER PRIMARY KEY,
			stopVideo 		BOOLEAN,
			changeVideo 	BOOLEAN,
			videoRequest 	BOOLEAN
		);
	`)

	createTable(`
		CREATE TABLE IF NOT EXISTS users (
			token		TEXT,
			id 			TEXT,
			username 	TEXT NOT NULL
		);
	`)

	createTable(`
		CREATE TABLE IF NOT EXISTS roomMembers (
			roomID 		INTEGER,
			userID 		TEXT,
			username	TEXT,
			perm   		TEXT
		);
	`)
}
