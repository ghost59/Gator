package main

import ("fmt"
		"context")

func handlerListUsers(s *state, cmd command) error {
		users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err


	}
	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Println("*",user.Name,"(current)")
		} else {
			fmt.Println("*",user.Name)
		}
	}

	return nil
}
