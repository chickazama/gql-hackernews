package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const (
	src    = "./data/database.db"
	driver = "sqlite3"
)

var DB *sql.DB

func InitDB() error {
	db, err := sql.Open(driver, src)
	if err != nil {
		return err
	}
	DB = db
	return LinksTableUp(db)
}

func LinksTableUp(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS "LINKS" (
		"ID" TEXT PRIMARY KEY,
		"Title" TEXT NOT NULL,
		"Address" TEXT NOT NULL,
		"UserID" TEXT NOT NULL
		);`
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}
