package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't delete users: %w", err)
	}
	err = s.cfg.SetUser("")
	if err != nil {
		return fmt.Errorf("couldn't reset current user: %w", err)
	}
	fmt.Println("Database reset successfully!")
	return nil
}
