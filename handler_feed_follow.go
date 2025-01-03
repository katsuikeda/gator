package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/katsuikeda/gator/internal/database"
)

func handlerFollow(s *state, cmd command, currentUser database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.name)
	}

	feedURL := cmd.args[0]

	feed, err := s.db.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("couldn't find feed by URL: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    currentUser.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}

	fmt.Println("Feed followed successfully:")
	fmt.Printf("* Feed:          %s\n", feedFollow.FeedName)
	fmt.Printf("* Followed By:   %s\n", feedFollow.UserName)

	return nil
}

func handlerFollowing(s *state, cmd command, currentUser database.User) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: %s", cmd.name)
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feed follows for the current user: %w", err)
	}

	if len(feedFollows) == 0 {
		fmt.Println("No feeds follows found for the current user")
		return nil
	}

	fmt.Println("The current user follows these feeds:")
	for _, follow := range feedFollows {
		fmt.Printf("* %s\n", follow.FeedName)
	}

	return nil
}

func handlerUnfollow(s *state, cmd command, currentUser database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.name)
	}

	feedURL := cmd.args[0]

	feed, err := s.db.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("couldn't find feed by URL: %w", err)
	}

	if err := s.db.DeleteFeedFollowForUser(context.Background(), database.DeleteFeedFollowForUserParams{
		FeedID: feed.ID,
		UserID: currentUser.ID,
	}); err != nil {
		return fmt.Errorf("couldn't delete feed follow: %w", err)
	}

	fmt.Println("Feed unfollowed successfully:")
	fmt.Printf("* %s (unfollowed)\n", feed.Name)

	return nil
}
