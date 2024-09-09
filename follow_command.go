package main

import (
	"fmt"
	"os"

	"github.com/mmcdole/gofeed"
)

func commandFollow(cfg *apiConfig, cliArg *string) error {
	feedParser := gofeed.NewParser()
	feed, err := feedParser.ParseURL(*cliArg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while searching for feed: %s", err.Error())
		return err
	}

	fmt.Println(feed.Title)
	return nil
}
