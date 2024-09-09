package main

import (
	"context"
	"fmt"
)

func commandPosts(cfg *apiConfig, cliArg *string) error {
	channelID, err := cfg.DB.TitleGetChannelID(context.Background(), *cliArg)
	if err != nil {
		return err
	}

	posts, err := cfg.DB.ChannelGetPosts(context.Background(), channelID)
	if err != nil {
		return err
	}

	fmt.Println("> Posts:")
	for _, post := range posts {
		fmt.Printf("> %s\n>    %s\n> %s\n>\n", post.Title, post.Description.String, post.Link)
	}
	return nil
}
