package models

import "database/sql"

type Link struct {
	ID      string
	Title   string
	Address string
	User    *User
	db      *sql.DB
}
