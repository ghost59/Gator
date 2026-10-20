package main

import ("gator/internal/database"
	"strconv"
	"fmt"
	"context"
	)

func Browse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.args) == 1 {
		parsedLimit, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("invalided limit: %w", err) 
		}

	limit = parsedLimit
	}
		posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: int32(limit),
	})
	if err != nil {
		return fmt.Errorf("couldn't get posts for user: %w", err)
	}

	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Format("Mon Jan 2"), post.Title)
		fmt.Printf("--- %s ---\n",post.Title)
		fmt.Printf("   %s\n", post.Description)
		fmt.Printf("Link: %s\n",post.Url)
		fmt.Println()
	}
	return nil
}
