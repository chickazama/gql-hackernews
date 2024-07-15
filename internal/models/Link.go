package models

import (
	"database/sql"

	"github.com/google/uuid"
)

type Link struct {
	ID      string
	Title   string
	Address string
	UserID  string
}

func NewLink(title, address, userID string) *Link {
	id := uuid.NewString()
	return &Link{
		ID:      id,
		Title:   title,
		Address: address,
	}
}

func (l *Link) Save(db *sql.DB) (int64, error) {
	query := `INSERT INTO "LINKS" (
		ID,
		Title,
		Address,
		UserID
	) VALUES (?, ?, ?, ?);`
	stmt, err := db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(l.ID, l.Title, l.Address, l.UserID)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, err
}
