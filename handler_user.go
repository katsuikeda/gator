package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/katsuikeda/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}

	userName := cmd.args[0]

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      userName,
	})
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	if err := s.cfg.SetUser(userName); err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User created successfully:")
	printUser(user)
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}

	userName := cmd.args[0]

	_, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	if err := s.cfg.SetUser(userName); err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("User switched successfully to %s\n", userName)
	return nil
}

func handlerUsers(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: %s", cmd.name)
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't list users: %w", err)
	}

	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
			continue
		}
		fmt.Printf("* %s\n", user.Name)
	}

	return nil
}
