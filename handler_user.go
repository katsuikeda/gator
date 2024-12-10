package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}

	userName := cmd.args[0]

	err := s.cfg.SetUser(userName)
	if err != nil {
		return fmt.Errorf("error setting username: %w", err)
	}

	fmt.Printf("The username: %s has been set\n", userName)
	return nil
}
