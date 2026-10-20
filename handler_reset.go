package main

import (
		"context"
		)


func handleReset(s *state, cmd command) error {

	err := s.db.DeleteAllUsers(context.Background())

	if err != nil {
		return err
	}

	return nil
	
	
}
