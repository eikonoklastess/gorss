package main

import (
	"context"
	"fmt"
)

func commandUnfollow(cfg *apiConfig, cliArg *string) error {
	err := cfg.DB.DeleteChannel(context.Background(), *cliArg)
	if err != nil {
		fmt.Println("> a problem occured while trying to delete a channel")
		return err
	}
	fmt.Println(*cliArg)
	return nil
}
