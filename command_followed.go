package main

import (
	"context"
	"fmt"
)

func commandFollowed(cfg *apiConfig, cliArg *string) error {
	channelsTitle, err := cfg.DB.ViewChannel(context.Background())
	if err != nil {
		fmt.Println("> viewing your followed channels has failed")
		return err
	}
	for _, channel := range channelsTitle {
		fmt.Println(channel)
	}
	return nil
}
