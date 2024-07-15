package database

import (
	"database/sql"

	"github.com/chickazama/gql-hackernews/internal/models"
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

func GetAllLinks() ([]models.Link, error) {
	var ret []models.Link
	query := "SELECT * FROM LINKS;"
	rows, err := DB.Query(query)
	if err != nil {
		return ret, err
	}
	defer rows.Close()
	for rows.Next() {
		var link models.Link
		err := rows.Scan(&link.ID, &link.Title, &link.Address, &link.UserID)
		if err != nil {
			return ret, err
		}
		ret = append(ret, link)
	}
	return ret, nil
}
