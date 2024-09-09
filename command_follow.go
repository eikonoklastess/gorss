package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/eikonoklastess/gorss/internal/database"
	"github.com/mmcdole/gofeed"
)

func commandFollow(cfg *apiConfig, cliArg *string) error {
	feedParser := gofeed.NewParser()
	feed, err := feedParser.ParseURL(*cliArg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while searching for feed: %s", err.Error())
		return err
	}

	err = cfg.DB.CreateChannel(context.Background(), database.CreateChannelParams{
		Title:       feed.Title,
		Link:        feed.Link,
		Description: toNullString(feed.Description),
		LastUpdated: sql.NullTime{},
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while inserting feed in db: %s", err.Error())
		return err
	}

	return nil
}
