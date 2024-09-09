package main

import (
	"database/sql"
)

type Channel struct {
	ID          int64
	Title       string
	Description sql.NullString
	Link        string
	LastUpdated sql.NullTime
	CreatedAt   sql.NullTime
}
