package main

import ( "fmt"
	"time"
	"github.com/google/uuid"
	"gator/internal/database"
	"context")

func follow(s *state, cmd command, user database.User) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("No args")
	}
	url := cmd.args[0]
	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return err
	}
	
	
  	params :=  database.CreateFeedFollowParams{
		ID:uuid.New(),
		CreatedAt:time.Now(),
		UpdatedAt:time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), params)

	if err != nil {
		return err
	}

	fmt.Printf("%s ... %s\n", feedFollow.FeedName, feedFollow.UserName)

	return nil
	
 }
