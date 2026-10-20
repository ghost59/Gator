 package main

 import ("fmt"
 		"context")

func Feeds(s *state, cmd command) error {

	currentFeed, err := s.db.Getfeed(context.Background())
	if err != nil {
		return err
	}

	for _, name := range currentFeed {

		users, err := s.db.GetUserId(context.Background(), name.UserID)

		if err != nil {
			return err
		}
		
		
		fmt.Println(name.Name, name.Url, users)
	}


	return nil
}
