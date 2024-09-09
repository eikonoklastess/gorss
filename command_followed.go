package main

import (
	"context"
	"fmt"
)

func commandFollowed(cfg *apiConfig, cliArg *string) error {
	channels, err := cfg.DB.GetChannels(context.Background())
	if err != nil {
		fmt.Println("> viewing your followed channels has failed")
		return err
	}
	for _, channel := range channels {
		fmt.Printf("> %s: %s\n", channel.Title, channel.Description.String)
	}
	return nil
}
