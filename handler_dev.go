// This file is only for development
package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: %s", cmd.name)
	}

	if err := s.db.DeleteAllUsers(context.Background()); err != nil {
		return fmt.Errorf("couldn't delete users: %w", err)
	}

	fmt.Println("Database reset successfully!")
	return nil
}
