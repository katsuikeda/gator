package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/katsuikeda/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, currentUser database.User) error {
	if len(cmd.args) > 1 {
		return fmt.Errorf("usage: %s <limit> (optional)", cmd.name)
	}

	limit := 2
	if len(cmd.args) == 1 {
		if parsedLimit, err := strconv.Atoi(cmd.args[0]); err == nil {
			limit = parsedLimit
		} else {
			return fmt.Errorf("invalid limit: %w", err)
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: currentUser.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("couldn't get posts for user: %w", err)
	}

	fmt.Printf("Found %d posts for user:\n", len(posts))
	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}

	return nil
}
