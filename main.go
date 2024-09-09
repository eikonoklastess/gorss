package main

import (
	"database/sql"
	"log"
	"os"

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
	schema, err := os.ReadFile("./sql/schema/001_test.sql")
	if err != nil {
		log.Fatalf("error while reading schema: %s", err.Error())
	}
	_, err = db.Exec(string(schema))
	if err != nil {
		log.Fatalf("error while executing schema: %s", err.Error())
	}
	dbQueries := database.New(db)
	cfg := apiConfig{
		DB: dbQueries,
	}

	startRepl(&cfg)
}
