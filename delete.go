package main

import ("context"
		"gator/internal/database"
		"fmt")


func unfollow(s *state, cmd command, user database.User) error{
		if len(cmd.args) == 0 {
			return fmt.Errorf("No commands within this args")
		}
		url := cmd.args[0]
		feed, err := s.db.GetFeedByURL(context.Background(),url)
		if err != nil {
			return err
		}

		deleteParams := database.DeleteFeedFollowParams{
			FeedID: feed.ID,
			UserID: user.ID,
		}

		err = s.db.DeleteFeedFollow(context.Background(),deleteParams)
		if err != nil {
			return err
		}

	return nil
		
	
	
}
