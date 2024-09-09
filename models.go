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

type Post struct {
	ID          int64
	ChannelID   int64
	Title       string
	Description sql.NullString
	Link        string
	PubDate     sql.NullTime
	CreatedAt   sql.NullTime
}
