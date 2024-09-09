package main

import (
	"database/sql"
	"time"
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

func toNullTime(t time.Time) sql.NullTime {
	if t.IsZero() {
		return sql.NullTime{
			Valid: false,
		}
	} else {
		return sql.NullTime{
			Valid: true,
			Time:  t,
		}
	}
}
