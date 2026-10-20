package main

import ("fmt"
		"time"
		"github.com/google/uuid"
		"gator/internal/database"
		"context"
		"log")

func handlerLogin(s *state, cmd command) error {
	fmt.Println("tacos")
	if len(cmd.args) == 0 {
		return fmt.Errorf("not commands") 
	}
	args := cmd.args[0]
	_, err := s.db.GetUser(context.Background(),args)
	if err != nil {
		return err
	}
	err = s.cfg.SetUser(args)
	if err != nil {
		return err
	}
	
	fmt.Println(s.cfg.CurrentUserName)
	return nil
}   

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("no commands")
	}
	args := cmd.args[0]
	user := database.CreateUserParams{
		ID:uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: args,
	}
	newUser, err := s.db.CreateUser(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}

	err = s.cfg.SetUser(args)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("New user has been added", args)
	
	log.Println(newUser)
	
	return nil
	
}


