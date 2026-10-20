package main

import ("gator/internal/database"
		"github.com/google/uuid"
		"context"
		"time"
		"fmt")

func addFeed(s *state, cmd command, user database.User) error {
    if len(cmd.args) < 2 {
        return fmt.Errorf("...")
    }
    name := cmd.args[0]
    url := cmd.args[1]

    // 1. get the current user - fix this to capture both return values
    
    // 2. build CreateFeedParams - what goes here?
	feedParams := database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: name,
		Url: url,
		UserID: user.ID,
		
	}    
	
    // 3. call s.db.CreateFeed(...)
    newFeed, err := s.db.CreateFeed(context.Background(), feedParams)
    if err != nil {
    	return err
    }

    params := database.CreateFeedFollowParams{
    	ID: uuid.New(),
    	CreatedAt:time.Now(),
    	UpdatedAt:time.Now(),
    	UserID:user.ID,
    	FeedID: newFeed.ID,
    }

    _ , err = s.db.CreateFeedFollow(context.Background(),params)
    if err != nil {
    	return err
    }

    // 4. print the new feed's fields
    fmt.Println(newFeed) 

    return nil
 }
