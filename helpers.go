package main

import (
	"database/sql"
)

func toNullString(str string) sql.NullString {
	if str == "" {
		return sql.NullString{
			Valid: false,
		}
	} else {
		return sql.NullString{
			Valid:  true,
			String: str,
		}
	}
}
