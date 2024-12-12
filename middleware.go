package main

import (
	"context"
	"fmt"

	"github.com/katsuikeda/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, currentUser database.User) error) func(s *state, cmd command) error {
	return func(s *state, cmd command) error {
		currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("couldn't find user: %w", err)
		}

		return handler(s, cmd, currentUser)
	}
}
