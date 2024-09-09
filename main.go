package main

import (
	"database/sql"
	"log"

	"github.com/eikonoklastess/gorss/internal/database"
	_ "modernc.org/sqlite"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	db, err := sql.Open("sqlite", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	dbQueries := database.New(db)
	cfg := apiConfig{
		DB: dbQueries,
	}

	startRepl(&cfg)
}
