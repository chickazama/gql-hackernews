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
	return nil
}
