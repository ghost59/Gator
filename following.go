package main

import (
		"fmt"
		"context"
		"gator/internal/database"
		)


func following(s *state, cmd command, user database.User) error {
	followingFeeds, err := s.db.GetFeedFollowsForUser(context.Background(),user.ID)
	if err != nil {
		return err
	}

	for _, feed := range followingFeeds {
		fmt.Println(feed.FeedName)
	}

	
	return nil
}
