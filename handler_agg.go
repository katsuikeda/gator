package main

import (
	"context"
	"fmt"

	"github.com/katsuikeda/gator/internal/rss"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: %s", cmd.name)
	}

	feed, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("couldn't fetch RSS feed: %w", err)
	}

	fmt.Printf("Fetched RSS feed successfully:\n %+v\n", feed)
	return nil
}
